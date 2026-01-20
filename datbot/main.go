package main

import (
	"context"
	"fmt"
	"go-discord-wrapper/connection"
	"go-discord-wrapper/functions"
	"go-discord-wrapper/types"
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()

	bot := connection.NewDiscordClient(
		os.Getenv("TOKEN"),
		types.AllIntentsExceptDirectMessage,
		&connection.DiscordClientSharding{
			TotalShards: 1,
			ShardID:     0,
		},
	)

	bot.OnGuildCreate(func(session *connection.DiscordClient, event *types.DiscordGuildCreateEvent) {
		fmt.Println("New guild")
	})

	bot.OnMessageCreate(func(session *connection.DiscordClient, event *types.DiscordMessageCreateEvent) {
		session.Logger.Info().Msgf("Received message: %s", event.Content)
	})

	bot.OnInteractionCreate(func(session *connection.DiscordClient, event *types.DiscordInteractionCreateEvent) {
		if event.IsCommand() {
			bot.Logger.Debug().Msgf("Received interaction command %s from %s", event.GetFullCommand(), event.Member.User.DisplayName())

			_, err := event.Reply(&types.DiscordInteractionResponse{
				Data: &types.DiscordInteractionResponseData{
					Content: fmt.Sprintf("You invoked the command: %s", event.GetFullCommand()),
					Flags:   types.DiscordMessageFlagEphemeral,
					Components: &[]types.AnyComponent{
						&types.ActionRow{
							Components: []types.AnyComponent{
								types.ButtonComponent{
									Style:    types.ButtonStyleSecondary,
									Label:    "Click Me!",
									CustomID: "button_click_me",
								},
							},
						},

						&types.ActionRow{
							Components: []types.AnyComponent{
								types.StringSelectMenuComponent{
									CustomID: "select_menu_1",
									Options: &[]types.StringSelectMenuComponentOption{
										{
											Label:       "Option 1",
											Value:       "option_1",
											Description: "This is the first option",
											Default:     true,
											Emoji: &types.DiscordEmoji{
												Name: "🔥",
											},
										},
										{
											Label:       "Option 2",
											Value:       "option_2",
											Description: "This is the second option",
										},
										{
											Label:       "Option 3",
											Value:       "option_3",
											Description: "This is the third option",
										},
									},
									Placeholder: "Choose an option",
									MinValues:   functions.PointerTo(1),
									MaxValues:   functions.PointerTo(1),
								},
							},
						},

						&types.ActionRow{
							Components: []types.AnyComponent{
								types.UserSelectMenuComponent{
									CustomID:    "select_menu_2",
									Placeholder: "Choose an option",
									MinValues:   functions.PointerTo(1),
									MaxValues:   functions.PointerTo(1),
									DefaultValues: &[]types.SelectDefaultValue{
										{
											Type: types.SelectDefaultValueTypeUser,
											ID:   event.Member.User.ID,
										},
									},
								},
							},
						},
					},
				},
				Type: types.DiscordInteractionCallbackTypeChannelMessageWithSource,
			})

			if err != nil {
				bot.Logger.Error().Msgf("Failed to create interaction response: %v", err)
			}
		}

		if event.IsButton() {
			bot.Logger.Debug().Msgf("Received button interaction with custom ID %s from %s", event.GetCustomID(), event.Member.User.DisplayName())

			if event.GetCustomID() == "button_click_me" {
				_, err := event.Reply(&types.DiscordInteractionResponse{
					Data: &types.DiscordInteractionResponseData{
						Content: "You clicked the button!",
						Flags:   types.DiscordMessageFlagEphemeral,
					},
					Type: types.DiscordInteractionCallbackTypeChannelMessageWithSource,
				})

				if err != nil {
					bot.Logger.Error().Msgf("Failed to create button interaction response: %v", err)
				}
			}
		}

		if event.IsAnySelectMenu() {
			bot.Logger.Debug().Msgf("Received select menu interaction with custom ID %s from %s", event.GetCustomID(), event.Member.User.DisplayName())
		}

		if event.IsAutocomplete() {
			bot.Logger.Debug().Msgf("Received autocomplete interaction for command %s from %s", event.GetFullCommand(), event.Member.User.DisplayName())
		}

		if event.IsModalSubmit() {
			bot.Logger.Debug().Msgf("Received modal submit interaction with custom ID %s from %s", event.GetCustomID(), event.Member.User.DisplayName())
		}
	})

	if err := bot.Login(); err != nil {
		panic(err)
	}

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	<-ctx.Done()
	bot.Logger.Info().Msg("Shutting down bot")
	bot.Shutdown()
}
