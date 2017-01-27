package xpnk_createGroupFromSlack

import (
	"fmt"
	"github.com/nlopes/slack"
)

/**************************************************************************************
*
*Takes a Slack user name, Slack group ID, and Xapnik group name. 
*Sends a Slack DM inviting the user to join the Xapnik group, along with the Xapnik URL.
*
**************************************************************************************/

type Slacker struct {
	Token				string
	Slacker				string
	SlackGroup			string
	XpnkGroup			string
}

func Invite (slacker Slacker) string {

	slacker_token			:= slacker.Token
	slacker_name 			:= "@"+slacker.Slacker
	slacker_group 			:= slacker.SlackGroup
	xpnk_group 				:= slacker.XpnkGroup
	
	api := slack.New(slacker_token)
	api.SetDebug(true)
	
	invite_text				:= "Hi there! Your Slack team, "+slacker_group+", has created a Xapnik group so you can easily boost each other on social media and never miss a thing. Just click this link to get started in about 15 seconds. It's private to team members only. http://localhost:8000/XAPNIK/#/group/"+xpnk_group+"/slack-invite"
	
	params					:= slack.PostMessageParameters{}
	channelID, timestamp, 
	err 					:= api.PostMessage(slacker_name, invite_text, params)
	
	if err != nil {
		fmt.Printf("SLACK MESSAGE FAILED  %s/n", err)
		return "FAILED"
	}
	fmt.Printf("SLACK MESSAGE SENT TO CHANNEL %s at %s/n", channelID, timestamp)
	
	return "Success!"
	
}