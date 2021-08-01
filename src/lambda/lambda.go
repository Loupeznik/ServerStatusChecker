package main

import (
	"os"

	"github.com/aws/aws-lambda-go/lambda"

	"loupeznik/serverstatuschecker/src/common"
)

func main() {
	lambda.Start(checker)
}

func checker() {
	common.LoadEnv()

	bot := common.Bot{
		Token:   os.Getenv("OAUTH_TOKEN"),
		Channel: os.Getenv("CHANNEL_ID"),
	}

	hosts := []string{}

	common.CheckHosts(hosts, bot)
}
