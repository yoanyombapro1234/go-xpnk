package xpnk_slack

import (
"net/http"
"fmt"
"strings"
)

type SlackCommand struct {
	Token				string			`form:"token" binding:"required"`			
	TeamID				string			`form:"team_id" binding:"required"`
	TeamDomain			string			`form:"team_domain" binding:"required"`
	ChannelID			string			`form:"channel_id" binding:"required"`
	ChannelName			string			`form:"channel_name" binding:"required"`
	UserID				string			`form:"user_id" binding:"required"`
	UserName			string			`form:"user_name" binding:"required"`
	Command				string			`form:"command" binding:"required"`
	Text				string			`form:"text" binding:"required"`
	ResponseURL			string			`form:"response_url" binding:"required"`
	TriggerID			string			`form:"trigger_id" binding:"required"`
} 
	
func SlackGroupStatus(slack_command SlackCommand) string{

	webhook				:=	slack_command.ResponseURL
	
	body				:= strings.NewReader(`"Hi there, "+slack_command.UserName`)
	
	req, err := http.NewRequest("POST", webhook, body)
	if err != nil {
		fmt.Printf("There was an error creating the request:  %s\n", err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("There was an error sending message to webhook:  %s\n", err)
	}
	defer resp.Body.Close()
	
	teamID				:= slack_command.TeamID
	
	return teamID
	
}