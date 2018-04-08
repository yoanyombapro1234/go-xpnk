package xpnk_slack

import (
"net/http"
"fmt"
"strings"
"strconv"
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
	
func SlackGroupStatus(slack_command SlackCommand) string{

	var stats xpnk_stats.GroupStats
	var statsText string

	teamSlackID			:= slack_command.TeamID
	var xpnkTeam XpnkTeam
	xpnkTeam			= GetGroupPosts(teamSlackID, "slack")
	stats				= xpnk_stats.GetStats(xpnkTeam.GroupID, xpnkTeam.GroupName)
	
	groupName			:=	stats.GroupName
	groupURL			:=	stats.GroupURL
	groupStats			:=	stats.Stats
	statsText 			= StringifyStats(groupStats)
	fmt.Printf("groupStats: %+s", groupStats)
		
	responseObject	 	:= strings.NewReader(`{
		"response_type": "ephemeral",      
		"text": "Here are `+groupName+` team's social media stats for the last 24 hours:",
		"attachments": [
			{  
				"text": "`+statsText+`\n You can see your team's activity and boost them by following this link:\n `+groupURL+`"
			}
		]
	}`)
	
	webhook				:= slack_command.ResponseURL	
	body				:= responseObject
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
	
	fmt.Printf("\nSLACK WEBHOOK RESPONSE:  %+s\n", resp)
	
	response := string(`{
		"response_type": "ephemeral",
		"text": "Great job! Get out there and beat the bots."
	}`)		
	
	return response
	
}

func GetGroupPosts (teamSlackID string, source string) XpnkTeam{
	var groupInfo XpnkTeam
	dbmap 					:= db_connect.InitDb()
	defer dbmap.Db.Close()
	
	err := dbmap.SelectOne(&groupInfo, "SELECT `group_name`, `Group_ID` FROM GROUPS WHERE source=? AND source_id=?", source, teamSlackID)
	
	if err == nil {
		//convert the group name into a hyphenated string for use in json filename
		groupInfo.GroupName = strings.Replace(groupInfo.GroupName, " ", "-", -1)	
	
		//convert all characters to lowercase
		groupInfo.GroupName = strings.ToLower(groupInfo.GroupName)
		
		return groupInfo
	} else {
		fmt.Printf("\n==========\n GetGroupPosts - Problemz with getting group_name for group in slackCommands line 85: \n%+v\n",err)
		return groupInfo
	}	
}

func StringifyStats(group_stats []xpnk_stats.MemberStats) string {
    statsString := ""
    for _, elem := range group_stats {
        statsString += elem.SlackName+": "+strconv.Itoa(elem.Tweets)+" tweets, "+strconv.Itoa(elem.IGs)+" Instagrams, "+strconv.Itoa(elem.Comments)+" comments.\n"
    }
    return statsString
}