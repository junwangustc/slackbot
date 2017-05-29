package main

import (
	"fmt"

	slackPkg "github.com/nlopes/slack"
)

func main() {
	fmt.Println("hello world")
	api := slackPkg.New("Your-Token")
	Users, err := api.GetUsers()
	if err != nil {
		fmt.Printf("Get Slack users info failed: %v", err)
	}
	for _, user := range Users {
		fmt.Println(user.Profile.Email, user.ID)
	}
	Channels, err := api.GetChannels(true)
	if err != nil {
		fmt.Printf("Get Slack channels info failed: %v", err)
	}
	for _, channel := range Channels {
		fmt.Println(channel.Name, channel.ID)
	}

}
