# Server Status Checker
A simple script to check if your webserver/website is up or not. If the site is down, it sends a notification to a Slack channel of your choosing.

![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)
[![License](https://img.shields.io/github/license/Loupeznik/ServerStatusChecker?style=for-the-badge)](./LICENSE)

## Prerequisites
- A Slack workspace with an authorized Slackbot
- A room to post alerts to within the Slack workspace
- One or more publicly accessible webservers to be checked (for now, these are filled in the *hosts* slice inside the *main* function - this will be changed in the future)

I found [this article](https://golangdocs.com/golang-create-your-own-slack-bot) very helpful when I was setting up the Slackbot for the first time.

## Installation
```bash
git clone https://github.com/Loupeznik/ServerStatusChecker.git
cd ServerStatusChecker
go get .
cp .env.example .env # fill the .env values with your Slack creds
go run .
```

You can also fill the *hosts* slice and build an executable with `go build .`, in that case, the .env file needs to exist in the same directory as the executable (or the .env variables need to be exported as your environment variables).

## License
This project is MIT licensed.
