// handlers/handlers.go

package handlers

import (
	"github.com/bwmarrin/discordgo"
)

func StartHandle(s *discordgo.Session) {
	// List of handlers
	Help().Setup(s)

	// ...
}
