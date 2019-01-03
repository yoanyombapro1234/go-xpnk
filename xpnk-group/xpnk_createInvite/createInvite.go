package xpnk_createInvite

/**************************************************************************************
Queen file for setting up a new Xapnik group from a Slack group.
1) Takes a group id 
2) Gets the group name associated with the group id 
3) Gets an xpnk_token
4) Updates group_tokens table with xpnk_token and group id
5) Returns the invitation URL
**************************************************************************************/

import (
	"fmt"
	"strings"
	"xpnk-group/xpnk_groupTokens"
	//"xpnk_auth"
	"xpnk-shared/db_connect"
)

func CreateInvite(group_id int, source string, identifier string) (string, error) {
	
	//retrieve the group_name for the new group
	dbmap 					:= db_connect.InitDb()
	defer dbmap.Db.Close()
	
	var xpnk_groupName 		string
	err 					:= dbmap.SelectOne(&xpnk_groupName, "SELECT group_name FROM GROUPS WHERE Group_ID=?", group_id)
	if err == nil {
	    fmt.Printf("\n==========\nCreateInvite xpnk_groupName: %+v", xpnk_groupName)
	} else {
		fmt.Printf("\n==========\nCreateInvite Problemz with select: \n%+v\n",err)
	}
	
	var token_request xpnk_groupTokens.GroupCount
	token_request.Source	 = source
	token_request.Identifier = identifier
	token_request.XpnkGroup  = group_id
	
	token	 				:= xpnk_groupTokens.SaveInviteToken(token_request)
		
//create invitation url
	group_name				:= strings.ToLower(xpnk_groupName)
	group_path	 			:= strings.Replace(group_name, " ", "-", -1)
	invite_domain			:= "https://xapnik.com/"
	invite_url				:= invite_domain+group_path+"/?xpnk_tkn="+token

return invite_url, err
		
}