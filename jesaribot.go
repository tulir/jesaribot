package main

import (
	"strings"

	"maunium.net/go/maubot"
	"maunium.net/go/maubot/slack"
	"maunium.net/go/maubot/telegram"
	"maunium.net/go/mauflag"
)

var tgToken = mauflag.MakeFull("t", "telegram", "The Telegram bot token to use", "").String()
var slackToken = mauflag.MakeFull("s", "slack", "The Slack bot token to use", "").String()

func main() {
	mauflag.Parse()
	bot := maubot.Create()

	if tgToken != nil && len(*tgToken) > 0 {
		tg, err := telegram.New(*tgToken)
		if err != nil {
			panic(err)
		}

		bot.Add(tg)
		err = tg.Connect()
		if err != nil {
			panic(err)
		}
	}

	if slackToken != nil && len(*slackToken) > 0 {
		slck, err := slack.New(*slackToken)
		if err != nil {
			panic(err)
		}

		bot.Add(slck)
		err = slck.Connect()
		if err != nil {
			panic(err)
		}
	}

	for message := range bot.Messages {
		if strings.Contains(strings.ToLower(message.Text()), "jesari") {
			message.ReplyWithRef("https://i.imgur.com/BftYhcU.gif")
		}
	}
}
