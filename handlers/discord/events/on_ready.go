package events

import (
	"fmt"

	"time"

	"github.com/bwmarrin/discordgo"
)

var (
	startTime    = time.Now()
	commandCount int
)

func OnReady(s *discordgo.Session, r *discordgo.Ready) {
	// Calculate the number of unique users
	uniqueUsers := make(map[string]bool)
	for _, guild := range s.State.Guilds {
		for _, member := range guild.Members {
			uniqueUsers[member.User.ID] = true
		}
	}

	// Update the status with the number of unique users
	status := fmt.Sprintf("🤖 Commands: %d | 🤖 Uptime: %s", commandCount, time.Since(startTime).String())
	s.UpdateCustomStatus(status)
}
