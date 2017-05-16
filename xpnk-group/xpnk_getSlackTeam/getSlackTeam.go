package xpnk_getSlackTeam

/**************************************************************************************
Takes a Slack team token and makes request to team.info Slack api endpoint
**************************************************************************************/

import (
  "fmt"
  "github.com/nlopes/slack"
  "xpnk-user/xpnk_createUserInsert"
)
	
func GetSlackTeam(token string) []xpnk_createUserInsert.User_Insert{

	groupToken := token
	
	api := slack.New(groupToken)
	api.SetDebug(true)
			
	teamList, err := api.GetUsers()
	fmt.Printf("response object:  %+v", teamList)
	
	if err != nil {
		fmt.Printf("At line 19 of getSlackTeam:  %+v",err)
	}	
	
	fmt.Printf("\nResponse Object:  %+v\n", teamList)
		
	var teamMembers []xpnk_createUserInsert.User_Insert
	
	for i := 0; i < len(teamList); i++ {
		var this_team_member xpnk_createUserInsert.User_Insert
		
		if teamList[i].ID != "U3T10SMN1" && 
		teamList[i].ID != "USLACKBOT" &&
		teamList[i].Deleted != true  {
			this_team_member.SlackID 		= teamList[i].ID
			this_team_member.SlackName 		= teamList[i].Name
			this_team_member.SlackAvatar 	= teamList[i].Profile.Image192
			this_team_member.ProfileImage 	= teamList[i].Profile.Image192
			
			teamMembers = append(teamMembers, this_team_member)
		}
	}
	
	fmt.Printf("\nTeam Slice:  %+v\n", teamMembers)
		
	return teamMembers
}	