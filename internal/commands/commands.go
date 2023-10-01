package commands

import (
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
)

const (
	THUMBS_UP   string = "👍"
	THUMBS_DOWN string = "👎"
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
		{
			Name:        "event",
			Description: "creates a event for the channel",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "name",
					Description: "the events name",
					Required:    true,
				},
				{
					Type:         discordgo.ApplicationCommandOptionChannel,
					Name:         "channel",
					Description:  "the events voice channel",
					ChannelTypes: []discordgo.ChannelType{2},
					Required:     true,
				},
				{
					Type:        discordgo.ApplicationCommandOptionInteger,
					Name:        "time",
					Description: "starting time of the event from now in minutes",
					Required:    true,
				},
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "description",
					Description: "the events description",
				},
			},
		},
		{
			Name:        "poll",
			Description: "creates a yes/no poll",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "topic",
					Description: "the polls topic",
					Required:    true,
				},
			},
		},
		{
			Name:        "wiki",
			Description: "looks up a given topic on wikipedia",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "topic",
					Description: "the topic to search for",
					Required:    true,
				},
			},
		},
	}

	// Command handlers executing the commands logic
	CommandHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
		// Dice roll command - Allows the user to chose D4 - D20 and chose how many of each should be rolled
		"roll": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			// Get options
			options := i.ApplicationCommandData().Options
			optionMap := make(map[string]*discordgo.ApplicationCommandInteractionDataOption, len(options))
			delimiter := "---------------"
			total := 0
			result := []string{}

			for _, opt := range options {
				optionMap[opt.Name] = opt
			}

			// Generate results
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

			// Respond
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: 4,
				Data: &discordgo.InteractionResponseData{
					Content: strings.Join(result, "\n"),
				},
			})
		},
		"event": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			// Get command option values
			options := i.ApplicationCommandData().Options
			optionMap := make(map[string]*discordgo.ApplicationCommandInteractionDataOption, len(options))
			for _, opt := range options {
				optionMap[opt.Name] = opt
			}

			// Create server event
			t := optionMap["time"].IntValue()
			startingTime := time.Now().Add(time.Duration(t) * time.Minute)
			endingTime := startingTime.Add(720 * time.Minute)

			_, err := s.GuildScheduledEventCreate(i.GuildID, &discordgo.GuildScheduledEventParams{
				Name:               optionMap["name"].StringValue(),
				Description:        optionMap["description"].StringValue(),
				ScheduledStartTime: &startingTime,
				ScheduledEndTime:   &endingTime,
				EntityType:         discordgo.GuildScheduledEventEntityTypeVoice,
				ChannelID:          optionMap["channel"].ChannelValue(s).ID,
				PrivacyLevel:       discordgo.GuildScheduledEventPrivacyLevelGuildOnly,
			})
			if err != nil {
				log.Printf("Error creating scheduled event: %v", err)

			}

			// Respond
			err = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: 4,
				Data: &discordgo.InteractionResponseData{
					Content: "Event created!",
					Flags:   discordgo.MessageFlagsEphemeral,
				},
			})
			if err != nil {
				fmt.Printf("Failed to respond to event creation: %v", err)
			}
		},
		"poll": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			// Get options
			options := i.ApplicationCommandData().Options
			topic := options[0].StringValue()

			// Respond
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: 4,
				Data: &discordgo.InteractionResponseData{
					Content: fmt.Sprintf("%v", topic),
				},
			})

			// React
			message, err := s.InteractionResponse(i.Interaction)
			if err != nil {
				fmt.Println(err)
			}
			s.MessageReactionAdd(message.ChannelID, message.ID, THUMBS_UP)
			s.MessageReactionAdd(message.ChannelID, message.ID, THUMBS_DOWN)
		},
		"wiki": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			options := i.ApplicationCommandData().Options
			topic := options[0].StringValue()
			fmt.Println(topic)

		},
	}
)
