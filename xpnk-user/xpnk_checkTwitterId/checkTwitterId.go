package xpnk_checkTwitterId

/**************************************************************************************
* Takes a Twitter user auth token and secret, group ID
* Retrieves the User Object from Twitter 
* Checks Users table for the Twitter ID
* Checks User_Groups table for user ID associated with group ID
* TODO If Twitter ID not found, creates new user & returns new user ID and empty groups list
* If Twitter ID is found, returns user ID and list of all user's groups
**************************************************************************************/

import (
	"fmt"
	"strconv"
	"xpnk-user/xpnk_user_structs"
	"xpnk_twitter/twitter_verify"
	"xpnk_twitter/get_user_by_twitter"
	"xpnk-user/xpnk_get_groups"
)

func CheckTwitterId (usertoken string, usersecret string) (xpnk_user_structs.UserGroups, error) {

	twitter_id, err := twitter_verify.AccountVerify(usertoken, usersecret)	
		
	fmt.Printf("\nTwitter userid: %+v\n", twitter_id)
	
	if twitter_id == "" {
		var user_groups xpnk_user_structs.UserGroups
		return user_groups, err
	} 

	xpnk_id, err := get_user_by_twitter.GetUserByTwitter(twitter_id)
	if err != nil {
		fmt.Printf("checkerTwitterId threw an error: %e", err)
	} 
	
	fmt.Printf("\nUsers Xapnik id: %s", xpnk_id)
	
	xpnk_user := strconv.Itoa(xpnk_id)
	
	user_groups, err := xpnk_get_groups.GetGroups(xpnk_user)
	
	return user_groups, err
	
}