package commands

import (
	"github.com/bwmarrin/discordgo"
)

var (
	// Array containing all the bots command details
	Commands = []*discordgo.ApplicationCommand{
		{
			Name:        "basic-command",
			Description: "Basic command",
		},
	}

	// Command handlers executing the commands logic
	CommandHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
		"basic-command": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: "Hey there! Congratulations, you just executed your first slash command",
				},
			})
		},
	}
)
