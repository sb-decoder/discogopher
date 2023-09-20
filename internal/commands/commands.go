package commands

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"

	"github.com/bwmarrin/discordgo"
)

var (
	// Array containing all the bots command details
	Commands = []*discordgo.ApplicationCommand{
		{
			Name:        "roll",
			Description: "returns a random number from the chosen dice",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionInteger,
					Name:        "d4",
					Description: "4 sided dice",
					MaxValue:    10,
				},
				{
					Type:        discordgo.ApplicationCommandOptionInteger,
					Name:        "d6",
					Description: "6 sided dice",
					MaxValue:    10,
				},
				{
					Type:        discordgo.ApplicationCommandOptionInteger,
					Name:        "d8",
					Description: "8 sided dice",
					MaxValue:    10,
				},
				{
					Type:        discordgo.ApplicationCommandOptionInteger,
					Name:        "d10",
					Description: "10 sided dice",
					MaxValue:    10,
				},
				{
					Type:        discordgo.ApplicationCommandOptionInteger,
					Name:        "d12",
					Description: "12 sided dice",
					MaxValue:    10,
				},
				{
					Type:        discordgo.ApplicationCommandOptionInteger,
					Name:        "d20",
					Description: "20 sided dice",
					MaxValue:    10,
				},
			},
		},
	}

	// Command handlers executing the commands logic
	CommandHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
		// Dice roll command - Allows the user to chose D4 - D20 and chose how many of each should be rolled
		"roll": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			options := i.ApplicationCommandData().Options
			optionMap := make(map[string]*discordgo.ApplicationCommandInteractionDataOption, len(options))
			delimiter := "---------------"
			total := 0
			result := []string{}

			for _, opt := range options {
				optionMap[opt.Name] = opt
			}

			for k, v := range optionMap {
				if key, err := strconv.Atoi(k[1:]); err == nil {
					for i := 0; i < int(v.IntValue()); i++ {
						val := rand.Intn(key-1) + 1
						total += val
						result = append(result, fmt.Sprintf("D%v  Result: %v", key, val))
					}
					result = append(result, delimiter)
				}
			}

			if len(result) != 0 {
				result = append(result, fmt.Sprintf("Total: %v", total))
			}

			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: strings.Join(result, "\n"),
				},
			})
		},
	}
)
