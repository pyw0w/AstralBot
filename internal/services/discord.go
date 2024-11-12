package services

import (
	"AstralBot/internal/httpclient"
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
	"time"

	"github.com/bwmarrin/discordgo"
	"go.uber.org/zap"
)

type NewsService struct {
	discordSession        *discordgo.Session
	logger                *zap.Logger
	discordChannelID      string
	httpClient            *httpclient.Client
	accessToken           string
	groupID               int
	alexandriaID          int
	checkInterval         time.Duration
	delayBetweenNewsFetch time.Duration
	isRunning             bool
	cancelFunc            context.CancelFunc
	lastPostID            int
}

func NewNewsService(discordSession *discordgo.Session, logger *logger.Logger, httpClient *http.Client, accessToken string) *NewsService {
	return &NewsService{
		discordSession:        discordSession,
		logger:                logger,
		discordChannelID:      "1219685603378069646",
		httpClient:            httpclient.NewClient(),
		accessToken:           accessToken,
		groupID:               37468416,
		alexandriaID:          34788285,
		checkInterval:         20 * time.Minute,
		delayBetweenNewsFetch: 10 * time.Second,
	}
}

func (s *NewsService) Start() {
	if s.isRunning {
		s.logger.Info("NewsService is already running.")
		return
	}

	s.logger.Info("Started NewsService")
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
		s.logger.Info("NewsService is not running.")
		return
	}

	s.logger.Info("Stopped NewsService")
	s.cancelFunc()
	s.isRunning = false
}

func (s *NewsService) fetchNews(group int, token string) {
	requestURI := fmt.Sprintf("https://api.vk.com/method/wall.get?owner_id=-%d&count=1&offset=1&access_token=%s&v=5.131", group, token)
	resp, err := s.httpClient.Get(requestURI)
	if err != nil {
		s.logger.Error("Failed to fetch news", zap.Error(err))
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		s.logger.Error("Failed to fetch news", zap.Int("status_code", resp.StatusCode))
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		s.logger.Error("Failed to read response body", zap.Error(err))
		return
	}

	var news VKRoot
	if err := json.Unmarshal(body, &news); err != nil {
		s.logger.Error("Failed to unmarshal response body", zap.Error(err))
		return
	}

	if len(news.Response.Items) == 0 {
		s.logger.Info("No news to process.")
		return
	}

	oldIDs := s.loadLastPostIDs()

	for _, post := range news.Response.Items {
		if oldIDs[post.ID] || post.ID == s.lastPostID {
			s.logger.Info("No new news to process.")
			continue
		}

		text := regexp.MustCompile(`^.#.$`).ReplaceAllString(post.Text, "")
		timestamp := time.Unix(int64(post.Date), 0)

		embed := &discordgo.MessageEmbed{
			Description: text,
			Color:       s.getEmbedColor(group),
			Timestamp:   timestamp.Format(time.RFC3339),
		}

		message := &discordgo.MessageSend{
			Embeds: []*discordgo.MessageEmbed{embed},
		}

		for _, attachment := range post.Attachments {
			switch attachment.Type {
			case "photo":
				s.addPhotoAttachment(message, post, attachment)
			case "video":
				s.addVideoAttachment(message, post, attachment)
			case "doc":
				s.addDocAttachment(message, post, attachment)
			}
		}

		channel, err := s.discordSession.Channel(s.discordChannelID)
		if err != nil {
			s.logger.Error("Failed to get channel", zap.Error(err))
			continue
		}

		if _, err := s.discordSession.ChannelMessageSendComplex(channel.ID, message); err != nil {
			s.logger.Error("Failed to send message", zap.Error(err))
			continue
		}

		s.lastPostID = post.ID
		s.saveLastPostID(post.ID)
	}
}

func (s *NewsService) getEmbedColor(group int) int {
	if group == s.alexandriaID {
		return 3447003 // Blue
	}
	return 15158332 // Red
}

func (s *NewsService) addPhotoAttachment(message *discordgo.MessageSend, post VKPost, attachment VKAttachment) {
	photoSize := getMaxSize(attachment.Photo.Sizes)
	if photoSize != nil {
		photoStream, err := s.httpClient.Get(photoSize.URL)
		if err != nil {
			s.logger.Error("Failed to get photo", zap.Error(err))
			return
		}
		defer photoStream.Body.Close()

		photoFileName := fmt.Sprintf("photo_%d_%d.jpg", post.ID, attachment.Photo.ID)
		message.Files = append(message.Files, &discordgo.File{
			Name:   photoFileName,
			Reader: photoStream.Body,
		})
	}
}

func (s *NewsService) addVideoAttachment(message *discordgo.MessageSend, post VKPost, attachment VKAttachment) {
	videoSize := getMaxSize(attachment.Video.Image)
	if videoSize != nil {
		videoStream, err := s.httpClient.Get(videoSize.URL)
		if err != nil {
			s.logger.Error("Failed to get video", zap.Error(err))
			return
		}
		defer videoStream.Body.Close()

		videoFileName := fmt.Sprintf("video_%d_%d.jpg", post.ID, attachment.Video.ID)
		message.Files = append(message.Files, &discordgo.File{
			Name:   videoFileName,
			Reader: videoStream.Body,
		})
	}
}

func (s *NewsService) addDocAttachment(message *discordgo.MessageSend, post VKPost, attachment VKAttachment) {
	docSize := getMaxSize(attachment.Doc.Preview.Photo.Sizes)
	if docSize != nil {
		docStream, err := s.httpClient.Get(docSize.Src)
		if err != nil {
			s.logger.Error("Failed to get doc", zap.Error(err))
			return
		}
		defer docStream.Body.Close()

		docFileName := fmt.Sprintf("doc_%d_%d.jpg", post.ID, attachment.Doc.ID)
		message.Files = append(message.Files, &discordgo.File{
			Name:   docFileName,
			Reader: docStream.Body,
		})
	}
}

func (s *NewsService) saveLastPostID(lastPostID int) {
	file, err := os.OpenFile("LastPostId.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		s.logger.Error("Failed to open file", zap.Error(err))
		return
	}
	defer file.Close()

	if _, err := file.WriteString(fmt.Sprintf("%d\n", lastPostID)); err != nil {
		s.logger.Error("Failed to write to file", zap.Error(err))
	}
}

func (s *NewsService) loadLastPostIDs() map[int]bool {
	lastPostIDs := make(map[int]bool)

	file, err := os.Open("LastPostId.txt")
	if err != nil {
		if os.IsNotExist(err) {
			return lastPostIDs
		}
		s.logger.Error("Failed to open file", zap.Error(err))
		return lastPostIDs
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		id, err := strconv.Atoi(scanner.Text())
		if err != nil {
			s.logger.Error("Failed to parse post ID", zap.Error(err))
			continue
		}
		lastPostIDs[id] = true
	}

	if err := scanner.Err(); err != nil {
		s.logger.Error("Failed to read file", zap.Error(err))
	}

	return lastPostIDs
}

func getMaxSize(sizes []VKSize) *VKSize {
	var maxSize *VKSize
	for _, size := range sizes {
		if maxSize == nil || size.Width*size.Height > maxSize.Width*maxSize.Height {
			maxSize = &size
		}
	}
	return maxSize
}

type VKRoot struct {
	Response struct {
		Items []VKPost `json:"items"`
	} `json:"response"`
}

type VKPost struct {
	ID          int            `json:"id"`
	Date        int            `json:"date"`
	Text        string         `json:"text"`
	Attachments []VKAttachment `json:"attachments"`
}

type VKAttachment struct {
	Type  string  `json:"type"`
	Photo VKPhoto `json:"photo,omitempty"`
	Video VKVideo `json:"video,omitempty"`
	Doc   VKDoc   `json:"doc,omitempty"`
}

type VKPhoto struct {
	ID    int      `json:"id"`
	Sizes []VKSize `json:"sizes"`
}

type VKVideo struct {
	ID    int      `json:"id"`
	Image []VKSize `json:"image"`
}

type VKDoc struct {
	ID      int       `json:"id"`
	Preview VKPreview `json:"preview"`
}

type VKPreview struct {
	Photo struct {
		Sizes []VKSize `json:"sizes"`
	} `json:"photo"`
}

type VKSize struct {
	URL    string `json:"url"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
}
