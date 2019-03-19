package xpnk_checkTwitterId

/**************************************************************************************
* Takes a user object
* Checks Users table for the Twitter ID
* Checks for Instagram and Disqus credentials if user found
* If user found, gets user's groups
* If Twitter ID is found, returns user ID, bool values for presence of IG and Disqus credentials, and list of all user's groups
* If Twitter ID is found, updates Twitter info in user record
**************************************************************************************/

import (
	"fmt"
	"strconv"
	"xpnk-shared/db_connect"
	"xpnk-user/xpnk_user_structs"
	"xpnk_twitter/get_user_by_twitter"
	"xpnk-user/xpnk_get_groups"
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
	
	fmt.Printf("\nUser Xapnik id: %s", xpnk_id)
	
	if xpnk_id == 0 {
		user_status.TwitterLoginNeeded = true 
	} else {
		user_status.TwitterLoginNeeded = false
		
		profile_image 		:= twitter_user.ProfileImage
		twitter_name		:= twitter_user.TwitterUser
		twitter_token		:= twitter_user.TwitterToken
		twitter_secret		:= twitter_user.TwitterSecret
		
		dbmap := db_connect.InitDb()
		defer dbmap.Db.Close()
		
		res, err := dbmap.Exec("UPDATE USERS SET profile_image=?, twitter_user=?, twitter_ID=?, twitter_authtoken=?, twitter_secret=? where user_ID=?", profile_image, twitter_name, twitter_id, twitter_token,  twitter_secret, xpnk_id)

		if err != nil {
			fmt.Printf("Error updating user's Twitter creds: %e", err.Error())
		} else {
			fmt.Printf("User Twitter creds rows updated: %v", res)
		}

	}
	
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
	
	if user_object.TwitterUser != "" {
		user_groups.ScreenName = user_object.TwitterUser
	} else if user_object.InstaUser != "" {
		user_groups.ScreenName = user_object.InstaUser
	}
	
	user_status.UserGroups = user_groups
	
		fmt.Printf("\nUser_status object: %+o\n", user_status)
	
	return user_status, err
	
}