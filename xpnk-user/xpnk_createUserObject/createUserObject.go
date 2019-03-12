package xpnk_createUserObject

/**************************************************************************************
Takes a generic User struct and prepares it for insertion into the USERS table
**************************************************************************************/

import (
	"fmt"
	"net/url"
)

//the one User struct to rule them all
type User_Object struct {
	Id					int		`db:"user_ID"`				
	SlackName			string		`db:"slack_name"			json:"SlackName"`
	SlackID				string		`db:"slack_userid"			json:"SlackID"`
	SlackAvatar			string		`db:"slack_avatar"			json:"SlackAvatar"`
	TwitterUser			string		`db:"twitter_user"			json:"TwitterUser"`
	TwitterID			string		`db:"twitter_ID"			json:"TwitterID"`
	TwitterToken		string		`db:"twitter_authtoken"		json:"TwitterToken"`
	TwitterSecret		string		`db:"twitter_secret"		json:"TwitterSecret"`
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

func CreateUserObject(newUser User_Object) User_Object {
		var this_user_object User_Object
		
		strippedProfileImage 				:= stripUrlScheme(newUser.ProfileImage)
		
		this_user_object.SlackName			= newUser.SlackName
		this_user_object.SlackID			= newUser.SlackID
		this_user_object.SlackAvatar	 	= newUser.SlackAvatar
		this_user_object.TwitterUser	 	= newUser.TwitterUser
		this_user_object.TwitterID		 	= newUser.TwitterID
		this_user_object.TwitterToken		= newUser.TwitterToken
		this_user_object.TwitterSecret		= newUser.TwitterSecret
		this_user_object.LastTweet		 	= newUser.LastTweet
		this_user_object.InstaUser		 	= newUser.InstaUser
		this_user_object.InstaUserID	 	= newUser.InstaUserID
		this_user_object.InstaAccessToken	= newUser.InstaAccessToken
		this_user_object.DisqusUserName	 	= newUser.DisqusUserName
		this_user_object.DisqusUserID	 	= newUser.DisqusUserID
		this_user_object.DisqusAccessToken	= newUser.DisqusAccessToken
		this_user_object.DisqusRefreshToken = newUser.DisqusRefreshToken
		this_user_object.FirstName		 	= newUser.FirstName
		this_user_object.LastName		 	= newUser.LastName
		this_user_object.ProfileImage	 	= strippedProfileImage		
						
		fmt.Printf("\n==========\nTHIS_USER_OBJECT: \n%+v\n",this_user_object)
		
		return this_user_object
}

func stripUrlScheme (img_url string) string {
	u, err := url.Parse(img_url)
	if err != nil {
		fmt.Println(err)
	}
	strippedUrl := "//" + u.Host + u.Path
	return (strippedUrl)
}