package services

import (
	"AstralBot/config"
	vk "AstralBot/handlers/discord/services/structrs"
	"AstralBot/internal/logger"
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"sync"
	"time"

	"github.com/bwmarrin/discordgo"
)

type NewsService struct {
	discordClient         *discordgo.Session
	logger                *logger.Logger
	discordChannelID      string
	httpClient            *http.Client
	accessToken           string
	groupID               int
	alexandriaID          int
	checkInterval         time.Duration
	delayBetweenNewsFetch time.Duration
	isRunning             bool
	cancelFunc            context.CancelFunc
	lastPostID            int
	mu                    sync.Mutex
}

func NewNewsService(discordClient *discordgo.Session, logger *logger.Logger, cfg *config.Config) *NewsService {
	return &NewsService{
		discordClient:         discordClient,
		logger:                logger,
		discordChannelID:      cfg.DiscordChannelID,
		httpClient:            &http.Client{},
		accessToken:           cfg.AccessToken,
		groupID:               37468416,
		alexandriaID:          34788285,
		checkInterval:         20 * time.Minute,
		delayBetweenNewsFetch: 10 * time.Second,
	}
}

func (s *NewsService) Start() {
	if s.isRunning {
		s.logger.Info("NewsService", "NewsService is already running.")
		return
	}

	s.logger.Info("NewsService", "Started NewsService")
	s.isRunning = true
	ctx, cancel := context.WithCancel(context.Background())
	s.cancelFunc = cancel

	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			default:
				s.fetchNews(s.groupID, s.accessToken)
				time.Sleep(s.delayBetweenNewsFetch)
				s.fetchNews(s.alexandriaID, s.accessToken)
				time.Sleep(s.checkInterval)
			}
		}
	}()
}

func (s *NewsService) Stop() {
	if !s.isRunning {
		s.logger.Info("NewsService", "NewsService is not running.")
		return
	}

	s.logger.Info("NewsService", "Stopped NewsService")
	s.cancelFunc()
	s.isRunning = false
}

func (s *NewsService) fetchNews(group int, token string) {
	requestURI := fmt.Sprintf("https://api.vk.com/method/wall.get?owner_id=-%d&count=1&offset=1&access_token=%s&v=5.131", group, token)
	resp, err := s.httpClient.Get(requestURI)
	if err != nil || resp.StatusCode != http.StatusOK {
		s.logger.Error("NewsService", fmt.Sprintf("Failed to fetch news. Status code: %d", resp.StatusCode))
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		s.logger.Error("NewsService", fmt.Sprintf("Failed to read response body: %v", err))
		return
	}

	var news vk.VKRoot
	if err := json.Unmarshal(body, &news); err != nil {
		s.logger.Error("NewsService", fmt.Sprintf("Failed to unmarshal response body: %v", err))
		return
	}

	if news.Response.Count == 0 {
		s.logger.Info("NewsService", "No news to process.")
		return
	}

	oldIds := s.loadLastPostIds()

	for _, post := range news.Response.Items {
		if oldIds[post.Id] || post.Id == s.lastPostID {
			s.logger.Info("NewsService", "No new news to process.")
			continue
		}

		text := regexp.MustCompile(`^.#.$`).ReplaceAllString(post.Text, "")
		timestamp := time.Unix(int64(post.Date), 0)

		mainEmbed := &discordgo.MessageEmbed{
			Description: text,
			Color:       s.getEmbedColor(group),
			Timestamp:   timestamp.Format(time.RFC3339),
		}

		message := &discordgo.MessageSend{
			Embeds: []*discordgo.MessageEmbed{mainEmbed},
		}

		for _, attachment := range post.Attachments {
			attachment := attachment // convert to the correct type
			switch attachment.Type {
			case "photo":
				s.addPhotoAttachment(message, post, attachment)
			case "video":
				s.addVideoAttachment(message, post, attachment)
			case "doc":
				s.addDocAttachment(message, post, attachment)
			}
		}

		channel, err := s.discordClient.Channel(s.discordChannelID)
		if err != nil {
			s.logger.Error("NewsService", fmt.Sprintf("Failed to get channel: %v", err))
			continue
		}

		if _, err := s.discordClient.ChannelMessageSendComplex(channel.ID, message); err != nil {
			s.logger.Error("NewsService", fmt.Sprintf("Failed to send message: %v", err))
			continue
		}

		s.lastPostID = post.Id
		s.saveLastPostId(post.Id)
	}
}

func (s *NewsService) getEmbedColor(group int) int {
	if group == s.alexandriaID {
		return 0x0000FF // Blue
	}
	return 0xFF0000 // Red
}

func (s *NewsService) addPhotoAttachment(message *discordgo.MessageSend, post vk.VKItem, attachment vk.VKAttachment) {
	photoSize := s.getLargestPhotoSize(attachment.Photo.Sizes)
	if photoSize != nil {
		photoStream, err := s.httpClient.Get(photoSize.Url)
		if err != nil {
			s.logger.Error("NewsService", fmt.Sprintf("Failed to get photo: %v", err))
			return
		}
		defer photoStream.Body.Close()

		photoFileName := fmt.Sprintf("photo_%d_%d.jpg", post.Id, attachment.Photo.Id)
		message.Files = append(message.Files, &discordgo.File{
			Name:   photoFileName,
			Reader: photoStream.Body,
		})
	}
}

func (s *NewsService) addVideoAttachment(message *discordgo.MessageSend, post vk.VKPost, attachment vk.VKAttachment) {
	videoSize := s.getLargestVideoSize(attachment.Video.Image)
	if videoSize != nil {
		videoStream, err := s.httpClient.Get(videoSize.Url)
		if err != nil {
			s.logger.Error("NewsService", fmt.Sprintf("Failed to get video: %v", err))
			return
		}
		defer videoStream.Body.Close()

		videoFileName := fmt.Sprintf("video_%d_%d.jpg", post.Id, attachment.Video.Id)
		message.Files = append(message.Files, &discordgo.File{
			Name:   videoFileName,
			Reader: videoStream.Body,
		})
	}
}

func (s *NewsService) addDocAttachment(message *discordgo.MessageSend, post vk.VKPost, attachment vk.VKAttachment) {
	docSize := s.getLargestPhotoSize(attachment.Doc.Preview.Photo.Sizes)
	if docSize != nil {
		docStream, err := s.httpClient.Get(docSize.Src)
		if err != nil {
			s.logger.Error("NewsService", fmt.Sprintf("Failed to get doc: %v", err))
			return
		}
		defer docStream.Body.Close()

		docFileName := fmt.Sprintf("doc_%d_%d.jpg", post.Id, attachment.Doc.Id)
		message.Files = append(message.Files, &discordgo.File{
			Name:   docFileName,
			Reader: docStream.Body,
		})
	}
}

func (s *NewsService) getLargestPhotoSize(sizes []vk.VKSizes) *vk.VKSizes {
	var largest *vk.VKSizes
	for _, size := range sizes {
		if largest == nil || size.Width*size.Height > largest.Width*largest.Height {
			largest = &size
		}
	}
	return largest
}

func (s *NewsService) getLargestVideoSize(images []vk.VKImage) *vk.VKImage {
	var largest *vk.VKImage
	for _, image := range images {
		if largest == nil || image.Width*image.Height > largest.Width*largest.Height {
			largest = &image
		}
	}
	return largest
}

func (s *NewsService) saveLastPostId(lastPostId int) {
	file, err := os.OpenFile("LastPostId.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		s.logger.Error("NewsService", fmt.Sprintf("Failed to open file: %v", err))
		return
	}
	defer file.Close()

	if _, err := file.WriteString(fmt.Sprintf("%d\n", lastPostId)); err != nil {
		s.logger.Error("NewsService", fmt.Sprintf("Failed to write to file: %v", err))
	}
}

func (s *NewsService) loadLastPostIds() map[int]bool {
	lastPostIds := make(map[int]bool)

	file, err := os.Open("LastPostId.txt")
	if err != nil {
		if !os.IsNotExist(err) {
			s.logger.Error("NewsService", fmt.Sprintf("Failed to open file: %v", err))
		}
		return lastPostIds
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		id, err := strconv.Atoi(scanner.Text())
		if err == nil {
			lastPostIds[id] = true
		}
	}

	if err := scanner.Err(); err != nil {
		s.logger.Error("NewsService", fmt.Sprintf("Failed to read file: %v", err))
	}

	return lastPostIds
}
