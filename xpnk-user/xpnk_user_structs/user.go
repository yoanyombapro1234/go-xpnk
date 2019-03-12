package xpnk_user_structs

type User_Object struct {
	Id					int			`db:"user_ID"`				
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