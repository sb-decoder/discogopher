package commands

import (
	"fmt"
	"math/rand"
	"strings"

	"github.com/bwmarrin/discordgo"
)

var (
	// Array containing all the bots command details
	Commands = []*discordgo.ApplicationCommand{
		{
			Name:        "basic-command",
			Description: "Basic command",
		},
		{
			Name:        "roll",
			Description: "returns a random number from the chosen dice",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionBoolean,
					Name:        "d4",
					Description: "4 sided dice",
				},
				{
					Type:        discordgo.ApplicationCommandOptionBoolean,
					Name:        "d6",
					Description: "6 sided dice",
				},
				{
					Type:        discordgo.ApplicationCommandOptionBoolean,
					Name:        "d8",
					Description: "8 sided dice",
				},
				{
					Type:        discordgo.ApplicationCommandOptionBoolean,
					Name:        "d10",
					Description: "10 sided dice",
				},
				{
					Type:        discordgo.ApplicationCommandOptionBoolean,
					Name:        "d12",
					Description: "12 sided dice",
				},
				{
					Type:        discordgo.ApplicationCommandOptionBoolean,
					Name:        "d20",
					Description: "20 sided dice",
				},
			},
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
		"roll": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			options := i.ApplicationCommandData().Options
			optionMap := make(map[string]*discordgo.ApplicationCommandInteractionDataOption, len(options))
			total := 0
			result := []string{}

			for _, opt := range options {
				optionMap[opt.Name] = opt
			}

			for k := range optionMap {
				switch k {
				case "d4":
					val := rand.Intn(3) + 1
					total += val
					result = append(result, fmt.Sprintf("D4  Result: %v", val))
				case "d6":
					val := rand.Intn(5) + 1
					total += val
					result = append(result, fmt.Sprintf("D6  Result: %v", val))
				case "d8":
					val := rand.Intn(7) + 1
					total += val
					result = append(result, fmt.Sprintf("D8  Result: %v", val))
				case "d10":
					val := rand.Intn(9) + 1
					total += val
					result = append(result, fmt.Sprintf("D10 Result: %v", val))
				case "d12":
					val := rand.Intn(11) + 1
					total += val
					result = append(result, fmt.Sprintf("D12 Result: %v", val))
				case "d20":
					val := rand.Intn(19) + 1
					total += val
					result = append(result, fmt.Sprintf("D20 Result: %v", val))
				default:
					result = append(result, "No dice was rolled.")
				}
			}

			if result[0] != "No dice was rolled." {
				result = append(result, fmt.Sprintf("---------------\n**Total: %v**", total))
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
