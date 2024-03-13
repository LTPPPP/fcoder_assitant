package handlers

import "os"

// DiscordMessageHandler là một cấu trúc để xử lý tin nhắn trong Discord.
type Config struct {
	Prefix string
}

func BotConfig() *Config {
	return &Config{
		Prefix: os.Getenv("DISCORD_BOT_PREFIX"),
	}
}
