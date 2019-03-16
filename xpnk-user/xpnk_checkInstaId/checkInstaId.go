package xpnk_checkInstaId

/**************************************************************************************
* Takes a user object
* Checks Users table for the Instagram ID
* Checks for Twitter and Disqus credentials if user found
* Checks User_Groups table for user ID associated with group ID
* TODO change this to just a not found - If Insta ID not found, creates new user & returns new user ID and empty groups list
* If Insta ID is found, returns user ID, bool values for presence of IG and Disqus credentials, and list of all user's groups
* TODO If Insta ID is found, updates Insta token & secret in user record
**************************************************************************************/

import (
	"fmt"
	"strconv"
	"xpnk-user/xpnk_user_structs"
	"xpnk_instagram/get_user_by_insta"
	"xpnk-user/xpnk_get_groups"
	//"xpnk-user/xpnk_insertMultiUsers"
	"xpnk-user/xpnk_createUserObject"
	"xpnk-user/xpnk_getUserObject"
)

func CheckInstaId (insta_user xpnk_createUserObject.User_Object) (xpnk_user_structs.UserStatus, error) {

	var user_status	xpnk_user_structs.UserStatus
	
	insta_id := insta_user.InstaUserID
		
	fmt.Printf("\nInsta userid: %+v \n", insta_id)

	xpnk_id, err := get_user_by_insta.GetUserByInsta(insta_id)
	if err != nil {
		fmt.Printf("checkerInstaId threw an error: %e", err)
	} 
		
	fmt.Printf("\nUser Xapnik id: %s", xpnk_id)
	
	if xpnk_id == 0 {
		user_status.InstagramLoginNeeded = true
	} else { 
		user_status.InstagramLoginNeeded = false
	}
	
	var user_object xpnk_createUserObject.User_Object
	
	xpnk_user := strconv.Itoa(xpnk_id)
	
	user_object, err = get_user_object.GetUserObject(xpnk_user)
	if err != nil {
		fmt.Printf("Get_user_object returned an error: %e", err.Error())
	}
	
	fmt.Printf("user_object: %+v", user_object)
	
	if user_object.TwitterID == "" {
		user_status.TwitterLoginNeeded = true
	} else {
		user_status.TwitterLoginNeeded = false
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