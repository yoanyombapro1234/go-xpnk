package xpnk_checkUserInvite

/**************************************************************************************
* Takes an xpnk_token and group_id
* Checks for xpnk_token in the group_tokens table 
* If xpnk_token found in group_tokens table, lookup group_name by associated group_id
* If group_name from url query matches group_name associated with group_id, return TRUE
**************************************************************************************/

import (
  "fmt"
  "strings"
  _ "github.com/go-sql-driver/mysql"
  "xpnk-shared/db_connect"
  "xpnk_auth"
)

type GroupObj struct {
	GroupID				int
	GroupName			string
}

func CheckUserInvite (token string, group_param string) GroupObj {

	user_token					:= validate_token(token)
	
    group_id					:= fetch_groupId(user_token)

	group_name					:= fetch_groupName(group_id)

	var return_obj				GroupObj
	
	if group_name == group_param {
		fmt.Printf("group_name:  %v \n", group_name)
		fmt.Printf("group_param:  %v \n", group_param)
		return_obj.GroupName	 = group_name
		return_obj.GroupID		 = group_id
	} else {
		fmt.Printf("group_name:  %v \n", group_name)
		fmt.Printf("group_param:  %v \n", group_param)
		return_obj.GroupName = "Your token isn't valid or doesn't match this Xapnik group."
	}
	fmt.Printf("return_obj:  %+v \n", return_obj)
	return return_obj
}

func validate_token (token string) string {
	token_auth				:= xpnk_auth.ParseToken(token, xpnk_auth.MySigningKey)
	var return_string		string
	if token_auth != 1  {
		return_string = "We're sorry, your token can't be validated."
		fmt.Printf("Sorry, this token isn't valid. Check the api output.")
	} else {
	  return_string = token	  
	}
	return return_string
}

func fetch_groupId (user_token string) int {
	var user_groupId		int
	var return_int			int
	
	dbmap := db_connect.InitDb()
	defer dbmap.Db.Close()

	err := dbmap.SelectOne(&user_groupId, "SELECT group_id FROM group_tokens WHERE xpnk_token=?",user_token)
	if err == nil {
		fmt.Printf("\n==========\nuser_groupId: %+v", user_groupId)
		return_int = user_groupId
	} else {
		return_int = 0
		fmt.Printf("\n==========\nProblemz with select in checkUserInvite line 45: \n%+v\n",err)
	}
	return return_int
}

func fetch_groupName (group_id int) string {
	user_groupId			:= group_id
	var user_groupName		string
	var return_string		string
	
	dbmap := db_connect.InitDb()
	defer dbmap.Db.Close()

	err := dbmap.SelectOne(&user_groupName, "SELECT group_name FROM GROUPS WHERE Group_ID=?", user_groupId)
	if err == nil {
		fmt.Printf("\n==========\nuser_groupName: %v \n", user_groupName)
		return_string = strings.ToLower(strings.Replace(user_groupName, " ", "-", -1))
	} else {
		fmt.Printf("\n==========\nProblemz with select in checkUserInvite line 54: \n%+v\n",err)
		return_string = "There was a problem retrieving the group_name from the db in checkUserInvite line 72"
	}
	return return_string
}