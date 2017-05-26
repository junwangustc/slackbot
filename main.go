package main

import (
	"fmt"

	"github.com/nlopes/slack"
)

func main() {
	api := slack.New("")
	//	api.SetDebug(true)
	ch, err := api.GetChannels(false)
	if err != nil {
		fmt.Println(err)
	} else {
		for k, v := range ch {
			fmt.Println("<----", k, v, "----->")
		}
	}
	rtm := api.NewRTM()
	go rtm.ManageConnection()
Loop:
	for {
		select {
		case msg := <-rtm.IncomingEvents:
			switch ev := msg.Data.(type) {
			case *slack.HelloEvent:
				// Ignore hello
			case *slack.ConnectedEvent:
			case *slack.MessageEvent:
				if ev.Type == "message" && ev.Username != "bot" {
					params := slack.PostMessageParameters{}
					attachment := slack.Attachment{
						Pretext: "",
						Text:    "hello world",
					}
					params.Attachments = []slack.Attachment{attachment}
					fmt.Println("======", ev.Channel, "----", ev.Text, "=======", ev.User)
					api.PostMessage(ev.Channel, "", params)

				}
				//				if ev.Channel == channelID && ev.Team == teamID {

				//				}
			case *slack.PresenceChangeEvent:
			case *slack.LatencyReport:
			case *slack.RTMError:
				fmt.Printf("Error: %s\n", ev.Error())
			case *slack.InvalidAuthEvent:
				fmt.Printf("Invalid credentials")
				break Loop
			default:
			}
		}
	}

}
