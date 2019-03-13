package xpnk_checkTwitterId

/**************************************************************************************
* Takes a Twitter user auth token and secret, group ID
* Retrieves the User Object from Twitter 
* Checks Users table for the Twitter ID
* Checks User_Groups table for user ID associated with group ID
* If Twitter ID not found, creates new user & returns new user ID and empty groups list
* TODO if Twitter ID found, update token and secret in USERS table
* If Twitter ID is found, returns user ID and list of all user's groups
**************************************************************************************/

import (
	"fmt"
	"strconv"
	"xpnk-user/xpnk_user_structs"
	"xpnk_twitter/get_user_by_twitter"
	"xpnk-user/xpnk_get_groups"
	"xpnk-user/xpnk_insertMultiUsers"
	"xpnk-user/xpnk_createUserObject"
)

func CheckTwitterId (twitter_user xpnk_createUserObject.User_Object) (xpnk_user_structs.UserGroups, error) {

	twitter_id := twitter_user.TwitterID
		
	fmt.Printf("\nTwitter userid: %+v\n", twitter_id)

	xpnk_id, err := get_user_by_twitter.GetUserByTwitter(twitter_id)
	if err != nil {
		fmt.Printf("checkerTwitterId threw an error: %e", err)
	} 
	
	if xpnk_id == 0 {
		fmt.Printf("Creating a new user from : %e", twitter_user.TwitterID)
		
		var userInsert				[]xpnk_createUserObject.User_Object
		userInsert 				 =  append(userInsert, twitter_user)
	
		newID, err_msg 			:=  xpnk_insertMultiUsers.InsertMultiUsers_2(userInsert)
		
		var user_groups xpnk_user_structs.UserGroups
		user_groups.Xpnk_id = strconv.Itoa(newID)
		return user_groups, err_msg
	}
	
	fmt.Printf("\nUser Xapnik id: %s", xpnk_id)
	
	xpnk_user := strconv.Itoa(xpnk_id)
	
	user_groups, err := xpnk_get_groups.GetGroups(xpnk_user)
	
	return user_groups, err
	
}