package main

import (
	"log"
	"os"

	client "github.com/bozd4g/go-http-client"
	"github.com/joho/godotenv"
	"github.com/slack-go/slack"
)

type Bot struct {
	Token   string
	Channel string
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	bot := Bot{
		Token:   os.Getenv("OAUTH_TOKEN"),
		Channel: os.Getenv("CHANNEL_ID"),
	}

	hosts := []string{}

	for _, host := range hosts {
		httpClient := client.New(host)
		request, err := httpClient.Get("/")

		if err != nil {
			panic(err)
		}

		response, err := httpClient.Do(request)
		if err != nil {
			panic(err)
		}

		if response.Get().StatusCode != 200 {
			sendMessage("Server at "+host+" seems to be down", bot)
		}
	}
}

func sendMessage(message string, bot Bot) {
	api := slack.New(bot.Token)

	_, _, err := api.PostMessage(
		bot.Channel,
		slack.MsgOptionText(message, false),
		slack.MsgOptionAsUser(true),
	)

	if err != nil {
		log.Fatalf("%s\n", err)
	}
}
