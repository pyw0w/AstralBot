package discord

import (
	"fmt"
	"time"

	"github.com/bwmarrin/discordgo"
)

// Status updates the bot's status every 15 seconds
func Status(s *discordgo.Session) {
	go func() {
		for {
			uniqueUsers := make(map[string]bool)
			for _, guild := range s.State.Guilds {
				for _, member := range guild.Members {
					uniqueUsers[member.User.ID] = true
				}
			}

			status := fmt.Sprintf("🤖 Commands: %d | 🤖 Uptime: %s", commandCount, time.Since(startTime).String())
			err := s.UpdateCustomStatus(status)
			if err != nil {
				return
			}

			time.Sleep(15 * time.Second)
		}
	}()
}
