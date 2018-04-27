package main

import (
	"strings"

	"maunium.net/go/mauflag"
	"maunium.net/go/gomatrix"
	"encoding/json"
)

var matrixHomeserver = mauflag.MakeFull("m", "matrix-server", "The Matrix homeserver to use", "").String()
var matrixUser = mauflag.MakeFull("u", "matrix-user", "The Matrix user localpart to use", "").String()
var matrixPassword = mauflag.MakeFull("p", "matrix-password", "The Matrix password to use", "").String()
var wantHelp, _ = mauflag.MakeHelpFlag()

func main() {
	mauflag.SetHelpTitles(
		"A simple gomatrix example that replies with a nice GIF when it receives the word \"jesari\" (duct tape in Finnish slang)",
		"jesaribot [-s matrixHomeserver] [-m matrixUser] [-p matrixPassword]")
	mauflag.Parse()
	if *wantHelp {
		mauflag.PrintHelp()
		return
	}

	client, _ := gomatrix.NewClient(*matrixHomeserver, "", "")

	resp, err := client.Login(&gomatrix.ReqLogin{
		Type:                     "m.login.password",
		User:                     *matrixUser,
		Password:                 *matrixPassword,
		InitialDeviceDisplayName: "Jesaribot",
	})
	if err != nil {
		panic(err)
	}

	client.AccessToken = resp.AccessToken
	client.UserID = resp.UserID

	syncer := client.Syncer.(*gomatrix.DefaultSyncer)
	syncer.OnEventType("m.room.message", func(evt *gomatrix.Event) {
		text, _ := evt.Content["body"].(string)
		if strings.Contains(strings.ToLower(text), "jesari") {
			client.SendMessageEvent(evt.RoomID, "m.room.message", json.RawMessage(`{
  "body": "putkiteippi.gif",
  "info": {
    "mimetype": "image/gif",
    "thumbnail_info": {
      "mimetype": "image/png",
      "h": 153,
      "w": 364,
      "size": 51302
    },
    "h": 153,
    "thumbnail_url": "mxc://maunium.net/iivOnCDjcGqGvnwnNWxSbAvb",
    "w": 364,
    "size": 2079294
  },
  "msgtype": "m.image",
  "url": "mxc://maunium.net/IkSoSYYrtaYJQeCaABSLqKiD"
}`))
		}
	})

	client.Sync()
}
