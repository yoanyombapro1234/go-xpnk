package xpnk_slack

import (
"net/http"
"fmt"
"strings"
_ "github.com/go-sql-driver/mysql"
"xpnk_stats"
"xpnk-shared/db_connect"
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

type XpnkTeam struct {
	GroupID				int				`db:"Group_ID"`
	GroupName			string			`db:"group_name"`
}
	
func SlackGroupStatus(slack_command SlackCommand) xpnk_stats.GroupStats{

	webhook				:= slack_command.ResponseURL	
	body				:= strings.NewReader(`{"text": "Hi there, "+slack_command.UserName}`)
	teamSlackID			:= slack_command.TeamID
	
	var xpnkTeam XpnkTeam
	xpnkTeam			= GetGroupPosts(teamSlackID, "slack")
	
	response			:= xpnk_stats.GetStats(xpnkTeam.GroupID, xpnkTeam.GroupName)
	
	req, err := http.NewRequest("POST", webhook, body)
	if err != nil {
		fmt.Printf("There was an error creating the request:  %s\n", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("There was an error sending message to webhook:  %s\n", err)
	}
	defer resp.Body.Close()
	
	fmt.Printf("/nSLACK WEBHOOK RESPONSE:  %+s/n", resp)
	
	return response
	
}

func GetGroupPosts (teamSlackID string, source string) XpnkTeam{
	var groupInfo XpnkTeam
	dbmap 					:= db_connect.InitDb()
	defer dbmap.Db.Close()
	
	err := dbmap.SelectOne(&groupInfo, "SELECT `group_name`, `Group_ID` FROM groups WHERE source=? AND source_id=?", source, teamSlackID)
	
	if err == nil {
		//convert the group name into a hyphenated string for use in json filename
		groupInfo.GroupName = strings.Replace(groupInfo.GroupName, " ", "-", -1)	
	
		//convert all characters to lowercase
		groupInfo.GroupName = strings.ToLower(groupInfo.GroupName)
		
		return groupInfo
	} else {
		fmt.Printf("\n==========\n GetGroupPosts - Problemz with getting group_name for group in slackCommands line 54: \n%+v\n",err)
		return groupInfo
	}
	
	
}