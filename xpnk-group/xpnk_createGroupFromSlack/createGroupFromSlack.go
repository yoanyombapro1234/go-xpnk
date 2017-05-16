package xpnk_createGroupFromSlack

/**************************************************************************************
Queen file for setting up a new Xapnik group from a Slack group.
1) Gets Group info from Slack  
2) Creates a new Xapnik Group from the Slack Group data
3) Gets the Slack Team members' data from Slack
4) Generates an xpnk_token for each Slack group member
5) Updates group_tokens table with user xpnk_tokens and XPNK Group ID
6) Sends a Slack DM to each group member with the Xapnik group URL+user_xpnk_token
**************************************************************************************/

import (
	"fmt"
	"strings"
	"xpnk-group/xpnk_getSlackGroup"
	"xpnk-group/xpnk_createGroupInsert"
	"xpnk-group/xpnk_insertGroup"
	"xpnk-group/xpnk_getSlackTeam"
	//"xpnk-group/xpnk_getSlackTeamXPNKIDs"
	//"xpnk-group/xpnk_addUsersToGroup"
	"xpnk-user/xpnk_createUserInsert"
	"xpnk-group/xpnk_groupTokens"
	//"xpnk-user/xpnk_insertMultiUsers"
	"xpnk_auth"
	"xpnk-shared/db_connect"
)

type SlackTeamTokens struct {
	TeamToken				string				`form:"team_token" binding:"required"`
	BotToken				string				`form:"bot_token"  binding:"required"`
} 

func CreateGroup(tokens SlackTeamTokens) string {
	
	token 					:= tokens.TeamToken
	bot_token 				:= tokens.BotToken

	groupInfo 				:= xpnk_getSlackGroup.GetSlackGroup(token)
	
	fmt.Printf("\n==========\nGROUP INFO: \n%+v\n",groupInfo)
	
	newGroup 				:= xpnk_createGroupInsert.CreateGroupInsert(groupInfo)
	
	fmt.Printf("\n==========\nNEW GROUP: \n%+v\n",newGroup)
	
	xpnk_insertGroup.InsertGroup(newGroup)
	
	//retrieve the group_id for the new group
	dbmap 					:= db_connect.InitDb()
	defer dbmap.Db.Close()
	
	var xpnk_groupID 		int
	err 					:= dbmap.SelectOne(&xpnk_groupID, "SELECT Group_ID FROM GROUPS WHERE source='Slack' and source_id='" + newGroup.SourceID + " ' ")
	if err == nil {
	    fmt.Printf("\n==========\nXPNK_GROUPID: %+v", xpnk_groupID)
	} else {
		fmt.Printf("\n==========\nProblemz with select: \n%+v\n",err)
	}
	
	//retrieve the team members info from Slack
	
	var teammembers 		[]xpnk_createUserInsert.User_Insert
	teammembers 			= xpnk_getSlackTeam.GetSlackTeam(token)
	
	//TODO change to (i) generate an xpnk token for each team member, (ii) add all tokens to GROUP -> xpnk_tokens field
	
	var group_count					xpnk_groupTokens.GroupCount
	
	for i :=0; i <len(teammembers); i++ {
		group_count.Identifier	= teammembers[i].SlackID
		group_count.Source		= "Slack"
		group_count.XpnkGroup	= xpnk_groupID
		
		v						:= xpnk_groupTokens.SaveGroupToken(group_count)
		if v != "Success!" {
		fmt.Printf("Expected Success!, got %v at line 79 of createGroupFromSlack", v)
		}	
		fmt.Printf("Added:  %+v/n", v)
	}
		
	//send Slack invites
	xpnkGroupPath 			:= strings.Replace(groupInfo.GroupName, " ", "-", -1)
	var thisSlacker 		Slacker
	for i := 0; i <len(teammembers); i++ {
		thisSlacker.Slacker 	= teammembers[i].SlackName
		thisSlacker.Token		= bot_token
		thisSlacker.SlackGroup	= groupInfo.GroupName
		thisSlacker.XpnkGroup	= xpnkGroupPath
		thisSlacker.XpnkToken	= xpnk_auth.GetNewGroupToken("Slack", teammembers[i].SlackID)
		
		//commented out for local testing - be sure to uncomment!
		Invite(thisSlacker)
		
		fmt.Printf("Invited %s to %s with this token %s\n", thisSlacker.Slacker, thisSlacker.XpnkGroup, thisSlacker.Token)
	}


	return "Check the database!"
}