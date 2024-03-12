// main.go

package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"fcoder_assitant/handlers"
	"fcoder_assitant/website"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	dg, err := discordgo.New("Bot " + os.Getenv("DISCORD_BOT_TOKEN"))
	if err != nil {
		fmt.Println("Error creating Discord session:", err)
		return
	}

	handlers.StartHandle(dg)

	// Start the web server
	go website.StartWebServer(dg)

	err = dg.Open()
	if err != nil {
		fmt.Println("Error opening connection:", err)
		return
	}

	fmt.Println("Bot is now running. Press CTRL+C to exit.")

	// Wait for a CTRL+C signal to gracefully close the bot
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	// Close the Discord session cleanly before exiting
	dg.Close()
}
