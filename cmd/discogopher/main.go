package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/Zwnow/discogopher/pkg/commands"
	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error while loading .env file")
	}
	flag.Parse()
}

func main() {

	d, err := discordgo.New("Bot " + os.Getenv("TOKEN"))
	if err != nil {
		log.Fatal("Error while starting discord session.")
	}

	d.AddHandler(commands.PingPong)

	d.Identify.Intents = discordgo.IntentsGuildMessages

	err = d.Open()

	if err != nil {
		fmt.Println("error opening connection, ", err)
		return
	}

	fmt.Println("Bot is now running. Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	d.Close()
}
