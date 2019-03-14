package xpnk_checkTwitterId

/**************************************************************************************
* Takes a user object
* Checks Users table for the Twitter ID
* Checks for Instagram and Disqus credentials if user found
* Checks User_Groups table for user ID associated with group ID
* TODO change this to just a not found - If Twitter ID not found, creates new user & returns new user ID and empty groups list
* If Twitter ID is found, returns user ID, bool values for presence of IG and Disqus credentials, and list of all user's groups
* If Twitter ID is found, updates Twitter token & secret in user record
**************************************************************************************/

import (
	"fmt"
	"strconv"
	"xpnk-user/xpnk_user_structs"
	"xpnk_twitter/get_user_by_twitter"
	"xpnk-user/xpnk_get_groups"
	//"xpnk-user/xpnk_insertMultiUsers"
	"xpnk-user/xpnk_createUserObject"
	"xpnk-user/xpnk_getUserObject"
)

func CheckTwitterId (twitter_user xpnk_createUserObject.User_Object) (xpnk_user_structs.UserStatus, error) {

	var user_status	xpnk_user_structs.UserStatus
	
	twitter_id := twitter_user.TwitterID
		
	fmt.Printf("\nTwitter userid: %+v\n", twitter_id)

	xpnk_id, err := get_user_by_twitter.GetUserByTwitter(twitter_id)
	if err != nil {
		fmt.Printf("checkerTwitterId threw an error: %e", err)
	} 
	
/*	
	if xpnk_id == 0 {
		fmt.Printf("Creating a new user from : %e", twitter_user.TwitterID)
		
		var userInsert				[]xpnk_createUserObject.User_Object
		userInsert 				 =  append(userInsert, twitter_user)
	
		newID, err_msg 			:=  xpnk_insertMultiUsers.InsertMultiUsers_2(userInsert)
		
		var user_groups xpnk_user_structs.UserGroups
		user_groups.Xpnk_id = strconv.Itoa(newID)
		return user_groups, err_msg
	}
*/	
	
	fmt.Printf("\nUser Xapnik id: %s", xpnk_id)
	
	user_status.TwitterLoginNeeded = false
	
	var user_object xpnk_createUserObject.User_Object
	
	xpnk_user := strconv.Itoa(xpnk_id)
	
	user_object, err = get_user_object.GetUserObject(xpnk_user)
	if err != nil {
		fmt.Printf("Get_user_object returned an error: %e", err.Error())
	}
	
	fmt.Printf("user_object: %+v", user_object)
	
	if user_object.InstaUserID == "" {
		user_status.InstagramLoginNeeded = true
	} else {
		user_status.InstagramLoginNeeded = false
	}
	
	if user_object.DisqusUserID == "" {
		user_status.DisqusLoginNeeded = true
	} else {
		user_status.DisqusLoginNeeded = false
	}
	
	user_groups, err := xpnk_get_groups.GetGroups(xpnk_user)
	
	user_status.UserGroups = user_groups
	
	fmt.Printf("\nUser_status object: %+o\n", user_status)
	
	return user_status, err
	
}