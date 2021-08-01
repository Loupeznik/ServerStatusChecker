package common

import (
	"log"

	client "github.com/bozd4g/go-http-client"
	"github.com/joho/godotenv"
	"github.com/slack-go/slack"
)

type Bot struct {
	Token   string
	Channel string
}

func CheckHosts(hosts []string, bot Bot) {
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
			SendMessage("Server at "+host+" seems to be down", bot)
		} else {
			log.Printf("Host " + host + " is up")
		}
	}
}

func SendMessage(message string, bot Bot) {
	api := slack.New(bot.Token)

	_, _, err := api.PostMessage(
		bot.Channel,
		slack.MsgOptionText(message, false),
		slack.MsgOptionAsUser(true),
	)

	if err != nil {
		log.Fatalf("%s\n", err)
	}

	log.Printf("Message sent " + message)
}

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
