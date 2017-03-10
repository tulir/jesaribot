package main

import (
	"strings"

	"maunium.net/go/maubot"
	"maunium.net/go/maubot/matrix"
	"maunium.net/go/maubot/slack"
	"maunium.net/go/maubot/telegram"
	"maunium.net/go/mauflag"
)

var tgToken = mauflag.MakeFull("t", "telegram", "The Telegram bot token to use", "").String()
var slackToken = mauflag.MakeFull("s", "slack", "The Slack bot token to use", "").String()
var matrixHomeserver = mauflag.MakeFull("m", "matrix-server", "The Matrix homeserver to use", "").String()
var matrixUser = mauflag.MakeFull("u", "matrix-user", "The Matrix user localpart to use", "").String()
var matrixPassword = mauflag.MakeFull("p", "matrix-password", "The Matrix password to use", "").String()
var wantHelp, _ = mauflag.MakeHelpFlag()

func main() {
	mauflag.SetHelpTitles(
		"A simple maubot example that replies with a nice GIF when it receives the word \"jesari\" (duct tape in Finnish slang)",
		"jesaribot [-t telegramToken] [-s slackToken] [-s matrixHomeserver] [-m matrixUser] [-p matrixPassword]")
	mauflag.Parse()
	if *wantHelp {
		mauflag.PrintHelp()
		return
	}

	bot := maubot.New()

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

	if matrixHomeserver != nil && len(*matrixHomeserver) > 0 {
		matrixx, err := matrix.New(*matrixHomeserver, *matrixUser, *matrixPassword)
		if err != nil {
			panic(err)
		}

		bot.Add(matrixx)
		err = matrixx.Connect()
		if err != nil {
			panic(err)
		}
	}

	for message := range bot.Messages() {
		if strings.Contains(strings.ToLower(message.Text()), "jesari") {
			message.Reply("https://i.imgur.com/BftYhcU.gif")
		}
	}
}
