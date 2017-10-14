package xpnk_createGroupFromSlack

import (
	"fmt"
	"strings"
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
	TestToken			string
	Slacker				string
	SlackGroup			string
	XpnkGroup			string
	XpnkToken			string
}

func Invite (slacker Slacker) string {

	slacker_token			:= slacker.Token
	test_token				:= slacker.TestToken
	slacker_name 			:= "@"+slacker.Slacker
	slacker_group 			:= slacker.SlackGroup
	xpnk_group_name         := slacker.XpnkGroup
	xpnk_group 				:= strings.ToLower(xpnk_group_name)
	xpnk_token				:= slacker.XpnkToken
	
	api := slack.New(slacker_token)
	api.SetDebug(true)
	
	var invite_domain		string
	
	if test_token != "true" {
		invite_domain		= "https://xapnik.com/"
	} else {
		invite_domain		= "http://localhost:8100/"
	}		
	
	invite_url				:= invite_domain+xpnk_group+"/slack-invite/?xpnk_tkn="+xpnk_token
	
	fmt.Printf("Invite url:   %s/n", invite_url)
	
	invite_text				:= "Hi there! Your Slack team, "+slacker_group+", has created a Xapnik group so you can easily boost each other on social media and never miss a thing. Just click this link to get started in about 15 seconds. It's private to team members only:  "+invite_url
	
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