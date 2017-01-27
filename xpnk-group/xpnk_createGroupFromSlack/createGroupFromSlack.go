package xpnk_createGroupFromSlack

/**************************************************************************************
Queen file for setting up a new Xapnik group from a Slack group.
1) Gets Group info from Slack  
2) Creates a new Xapnik Group from the Slack Group data
3) Gets the Slack Team members' data from Slack
4) Creates new Xapnik Users from the Slack Team members' data
5) Gets all the new Xapnik Users xpnkID's
6) Adds the new User xpnkID's to the User_Groups table associated the new group_ID
7) Sends a Slack DM to each group member with the Xapnik group URL
**************************************************************************************/

import (
	"fmt"
	"strings"
	"xpnk-group/xpnk_getSlackGroup"
	"xpnk-group/xpnk_createGroupInsert"
	"xpnk-group/xpnk_insertGroup"
	"xpnk-group/xpnk_getSlackTeam"
	"xpnk-group/xpnk_getSlackTeamXPNKIDs"
	"xpnk-group/xpnk_addUsersToGroup"
	"xpnk-user/xpnk_createUserInsert"
	"xpnk-user/xpnk_createMultiUserInsert"
	"xpnk-user/xpnk_insertMultiUsers"
	"xpnk-shared/db_connect"
)

type SlackTeamTokens struct {
	TeamToken			string					`form:"team_token" binding:"required"`
	BotToken			string					`form:"bot_token"  binding:"required"`
} 

func CreateGroup(tokens SlackTeamTokens) string {
	
	token := tokens.TeamToken
	bot_token := tokens.BotToken

	groupInfo := xpnk_getSlackGroup.GetSlackGroup(token)
	
	fmt.Printf("\n==========\nGROUP INFO: \n%+v\n",groupInfo)
	
	newGroup := xpnk_createGroupInsert.CreateGroupInsert(groupInfo)
	
	fmt.Printf("\n==========\nNEW GROUP: \n%+v\n",newGroup)
	
	xpnk_insertGroup.InsertGroup(newGroup)
	
	//retrieve the group_id for the new group
	
	dbmap := db_connect.InitDb()
	defer dbmap.Db.Close()
	
	var xpnk_groupID int
	err := dbmap.SelectOne(&xpnk_groupID, "SELECT Group_ID FROM GROUPS WHERE source='Slack' and source_id='" + newGroup.SourceID + " ' ")
	if err == nil {
	    fmt.Printf("\n==========\nXPNK_GROUPID: %+v", xpnk_groupID)
	} else {
		fmt.Printf("\n==========\nProblemz with select: \n%+v\n",err)
	}
	
	//retrieve the team members info from Slack
	
	var teammembers []xpnk_createUserInsert.User_Insert
	teammembers = xpnk_getSlackTeam.GetSlackTeam(token)
	
	//add the team members to xapnik as users
	teaminsert := xpnk_createMultiUserInsert.CreateMultiUserInsert(teammembers)
	xpnk_insertMultiUsers.InsertMultiUsers(teaminsert)
	
	//add the team members xpnk_id's to the group
	var memberXPNKIDs []int
	memberIDs := xpnk_getSlackTeamXPNKIDs.GetSlackUserXPNKID(teammembers)

	for i :=0; i <len(memberIDs); i++ {
		thisXPNKID := memberIDs[i].XPNK_ID
		
		memberXPNKIDs = append(memberXPNKIDs, thisXPNKID)
	}
		
	
	xpnk_addUsersToGroup.AddUsersToGroup(xpnk_groupID, memberXPNKIDs)
	
	//send Slack invites
	xpnkGroupPath := strings.Replace(groupInfo.GroupName, " ", "-", -1)
	var thisSlacker Slacker
	for i := 0; i <len(teammembers); i++ {
		thisSlacker.Slacker 	= teammembers[i].SlackName
		thisSlacker.Token		= bot_token
		thisSlacker.SlackGroup	= groupInfo.GroupName
		thisSlacker.XpnkGroup	= xpnkGroupPath
		
		Invite(thisSlacker)
		
		fmt.Printf("Invited %s to %s with this token %s\n", thisSlacker.Slacker, thisSlacker.XpnkGroup, thisSlacker.Token)
	}

	return "Check the database!"
}