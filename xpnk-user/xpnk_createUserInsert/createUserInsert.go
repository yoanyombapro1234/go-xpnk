package xpnk_createUserInsert

/**************************************************************************************
Takes a generic User struct and prepares it for insertion into the USERS table
**************************************************************************************/

import (
	"fmt"
	"database/sql"
)

//the one User struct to rule them all
type User_Insert struct {
	SlackName			string					`db:"slack_name"`
	SlackID				string					`db:"slack_userid"`
	SlackAvatar			string					`db:"slack_avatar"`
	TwitterUser			string					`db:"twitter_user"`
	InstaUser			string					`db:"insta_user"`
	FirstName			string					`db:"first_name"`
	LastName			string					`db:"last_name"`
	TwitterID			sql.NullString			`db:"twitter_ID"`
	LastTweet			string					`db:"last_tweet"`
	InstaUserID			sql.NullString			`db:"insta_userid"`
	InstaAccessToken	string					`db:"insta_accesstoken"`
	DisqusUserName		string					`db:"disqus_username"`
	DisqusUserID		sql.NullString			`db:"disqus_userid"`
	DisqusAccessToken	string					`db:"disqus_accesstoken"`
	DisqusRefreshToken	string					`db:"disqus_refreshtoken"`
	ProfileImage		string					`db:"profile_image"`			
}

func CreateUserInsert(newUser User_Insert) User_Insert {
//func CreateUserInsert(newUser User_Insert) string {

		var this_user_insert User_Insert
		
		this_user_insert.SlackName			= newUser.SlackName
		this_user_insert.SlackID			= newUser.SlackID
		this_user_insert.SlackAvatar	 	= newUser.SlackAvatar
		this_user_insert.TwitterUser	 	= newUser.TwitterUser
		this_user_insert.InstaUser		 	= newUser.InstaUser
		this_user_insert.FirstName		 	= newUser.FirstName
		this_user_insert.LastName		 	= newUser.LastName
		this_user_insert.TwitterID		 	= newUser.TwitterID
		this_user_insert.LastTweet		 	= newUser.LastTweet
		this_user_insert.InstaUserID	 	= newUser.InstaUserID
		this_user_insert.InstaAccessToken	= newUser.InstaAccessToken
		this_user_insert.DisqusUserName	 	= newUser.DisqusUserName
		this_user_insert.DisqusUserID	 	= newUser.DisqusUserID
		this_user_insert.DisqusAccessToken	= newUser.DisqusAccessToken
		this_user_insert.DisqusRefreshToken = newUser.DisqusRefreshToken
		this_user_insert.ProfileImage	 	= newUser.ProfileImage		
						
		fmt.Printf("\n==========\nTHIS_USEr_INSERT: \n%+v\n",this_user_insert)
		
		return this_user_insert
}