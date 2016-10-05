package xpnk_getSlackGroup

/**************************************************************************************
Takes a Slack team token and makes request to team.info Slack api endpoint
**************************************************************************************/

import (
  "fmt"
  "github.com/nlopes/slack"
  "xpnk-group/xpnk_createGroupInsert"
)
	
func GetSlackGroup(token string) xpnk_createGroupInsert.New_Group{

	source := "Slack"
	groupToken := token
	
	api := slack.New(groupToken)
	api.SetDebug(true)
	
	groupInfo, err := api.GetTeamInfo()
	fmt.Printf("response object:  %+v", groupInfo)
	
	if err != nil {
		fmt.Printf("At line 37 of getSlackGroup:  %+v",err)
	}	
	
	fmt.Printf("Response Object:  %+v\n", groupInfo)

	//pull team Slack ID and name out of the json response and put into New_Group struct

	var newGroup xpnk_createGroupInsert.New_Group
	newGroup.SourceID = groupInfo.ID
	fmt.Printf("\n====================\nGroup ID: %v\n", newGroup.SourceID)
	newGroup.GroupName = groupInfo.Name
	fmt.Printf("\n====================\nGroupName: %v\n", newGroup.GroupName)
	newGroup.Source = source	
		
	return newGroup
}	