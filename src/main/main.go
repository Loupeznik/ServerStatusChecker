package main

import (
	"os"

	"loupeznik/serverstatuschecker/src/common"
)

func main() {
	common.LoadEnv()

	bot := common.Bot{
		Token:   os.Getenv("OAUTH_TOKEN"),
		Channel: os.Getenv("CHANNEL_ID"),
	}

	hosts := []string{}

	common.CheckHosts(hosts, bot)
}
