package discord

import (
	"fmt"
	"time"

	"github.com/bwmarrin/discordgo"
)

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
			s.UpdateCustomStatus(status)

			time.Sleep(15 * time.Second)
		}
	}()
}
