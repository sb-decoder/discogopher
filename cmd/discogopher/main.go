package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	c "github.com/Zwnow/discogopher/internal/commands"
	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error while loading .env file")
	}
}

func main() {
	// Start a new discord session by authenticating with the bot-token
	d, err := discordgo.New("Bot " + os.Getenv("TOKEN"))
	if err != nil {
		log.Fatal("Error while starting discord session.")
	}

	// Add handlers for commands so the bot knows what to do once a command is initiated by the user.
	d.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if h, ok := c.CommandHandlers[i.ApplicationCommandData().Name]; ok {
			h(s, i)
		}
	})

	// Opens the connection -> bot goes online
	err = d.Open()
	if err != nil {
		fmt.Println("error opening connection, ", err)
		return
	}

	// Register commands from commands.go file
	registeredCommands := make([]*discordgo.ApplicationCommand, len(c.Commands))
	for i, v := range c.Commands {
		cmd, err := d.ApplicationCommandCreate(d.State.User.ID, "818211523109191700", v)
		if err != nil {
			log.Panicf("Cannot create '%v' command: %v", v.Name, err)
		}
		registeredCommands[i] = cmd
	}

	// Properly shut down the bot once exit signal is received
	fmt.Println("Bot is now running. Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	d.Close()
}
