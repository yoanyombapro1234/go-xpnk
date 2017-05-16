package xpnk_createUserInsert

/**************************************************************************************
Takes a generic User struct and prepares it for insertion into the USERS table
**************************************************************************************/

import (
	"fmt"
)

//the one User struct to rule them all
type User_Insert struct {
	SlackName			string		`db:"slack_name"			json:"SlackName"`
	SlackID				string		`db:"slack_userid"			json:"SlackID"`
	SlackAvatar			string		`db:"slack_avatar"			json:"SlackAvatar"`
	TwitterUser			string		`db:"twitter_user"			json:"TwitterUser"`
	TwitterID			string		`db:"twitter_ID"			json:"TwitterID"`
	TwitterToken		string		`db:"twitter_authtoken"		json:"TwitterToken"`
	LastTweet			string		`db:"last_tweet"			json:"LastTweet"`
	InstaUser			string		`db:"insta_user"			json:"InstaUser"`
	InstaUserID			string		`db:"insta_userid"			json:"InstaUserID"`
	InstaAccessToken	string		`db:"insta_accesstoken"		json:"InstaAccessToken"`
	DisqusUserName		string		`db:"disqus_username"		json:"DisqusUserName"`
	DisqusUserID		string		`db:"disqus_userid"			json:"DisqusUserID"`
	DisqusAccessToken	string		`db:"disqus_accesstoken"	json:"DisqusAccessToken"`
	DisqusRefreshToken	string	    `db:"disqus_refreshtoken"	json:"DisqusRefreshToken"`
	FirstName			string		`db:"first_name"			json:"FirstName"`
	LastName			string		`db:"last_name"				json:"LastName"`
	ProfileImage		string		`db:"profile_image"			json:"ProfileImage"`
}

func CreateUserInsert(newUser User_Insert) User_Insert {
//func CreateUserInsert(newUser User_Insert) string {

		var this_user_insert User_Insert
		
		this_user_insert.SlackName			= newUser.SlackName
		this_user_insert.SlackID			= newUser.SlackID
		this_user_insert.SlackAvatar	 	= newUser.SlackAvatar
		this_user_insert.TwitterUser	 	= newUser.TwitterUser
		this_user_insert.TwitterID		 	= newUser.TwitterID
		this_user_insert.TwitterToken		= newUser.TwitterToken
		this_user_insert.LastTweet		 	= newUser.LastTweet
		this_user_insert.InstaUser		 	= newUser.InstaUser
		this_user_insert.InstaUserID	 	= newUser.InstaUserID
		this_user_insert.InstaAccessToken	= newUser.InstaAccessToken
		this_user_insert.DisqusUserName	 	= newUser.DisqusUserName
		this_user_insert.DisqusUserID	 	= newUser.DisqusUserID
		this_user_insert.DisqusAccessToken	= newUser.DisqusAccessToken
		this_user_insert.DisqusRefreshToken = newUser.DisqusRefreshToken
		this_user_insert.FirstName		 	= newUser.FirstName
		this_user_insert.LastName		 	= newUser.LastName
		this_user_insert.ProfileImage	 	= newUser.ProfileImage		
						
		fmt.Printf("\n==========\nTHIS_USEr_INSERT: \n%+v\n",this_user_insert)
		
		return this_user_insert
}