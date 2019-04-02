package users

import (
	"fmt"
	 _ "github.com/go-sql-driver/mysql"
	 "database/sql"
	 "github.com/gin-gonic/gin"
	 "xpnk-shared/db_connect"
)

type XPNKUser 			struct {
	User_ID				int			   `db:"user_ID"			json:"user_ID"`
	Slack_userid		string		   `db:"slack_userid"		json:"slack_userid"`
	Slack_name			string		   `db:"slack_name"			json:"slack_name"`
	Twitter_user		string		   `db:"twitter_user"		json:"twitter_user"`
	Twitter_ID			string		   `db:"twitter_ID"			json:"twitter_ID"`
	Twitter_token		string		   `db:"twitter_authtoken"	json:"twitter_authtoken"`
	Twitter_secret		string		   `db:"twitter_secret"		json:"twitter_secret"`
	Insta_user			string		   `db:"insta_user"			json:"insta_user"`
	Insta_userid		string		   `db:"insta_userid"		json:"insta_userid"`
	Insta_token			string		   `db:"insta_accesstoken"	json:"insta_accesstoken"`
	Disqus_username		sql.NullString `db:"disqus_username"	json:"disqus_username"`
	Disqus_userid		sql.NullString `db:"disqus_userid"		json:"disqus_userid"`
	Disqus_token		string		   `db:"disqus_accesstoken"	json:"disqus_accesstoken"`
	Profile_image		string		   `db:"profile_image"		json:"profile_image"`
}

type TwitterID struct {
	 Twttr_userid		string					`form:"id" binding:"required"`
}

func UsersByTwitterID_2 (c *gin.Context) {
	twitter_id					:= c.Param("id")
	var user					XPNKUser 	
	var err_msg					error		
	fmt.Printf("twitter_id:  %v \n", twitter_id)
	if twitter_id == "" {
		c.JSON(422, gin.H{"error": "Invalid or missing Twitter user ID."})
	} else {
		user, err_msg 			= get_user_by_twitter(twitter_id)
		if user.Twitter_ID != twitter_id {
			c.JSON(400, err_msg.Error())
		} else {
			c.JSON(200, user)
		}
	}
}

func UsersByTwitterID (c *gin.Context) {
	var twitterId				TwitterID
	var user					XPNKUser 
	var err_msg					error			
	c.Bind(&twitterId)
	twitter_id					:= twitterId.Twttr_userid
	fmt.Printf("twitter_id:  %v \n", twitter_id)
	if twitter_id == "" {
		c.JSON(422, gin.H{"error": "Invalid or missing Twitter user ID."})
	} else {
		user, err_msg 			= get_user_by_twitter(twitter_id)
		if user.Twitter_ID != twitter_id {
			c.JSON(202, user)
			fmt.Printf("error:  %v \n", err_msg)
		} else {
			c.JSON(200, user)
		}
	}
}

func get_user_by_twitter(twitter_id string) (XPNKUser, error) {
	dbmap := db_connect.InitDb()
	defer dbmap.Db.Close()
	
	var xpnkUser			XPNKUser
	var err_msg				error
	twitterId				:= twitter_id
	
	err	:= dbmap.SelectOne(&xpnkUser, "SELECT `user_ID`, `twitter_user`, `twitter_ID`, `twitter_authtoken`, `twitter_secret`, `insta_user`, `insta_userid`, `insta_accesstoken`, `disqus_username`, `disqus_userid`, `disqus_accesstoken`, `profile_image` FROM USERS WHERE twitter_ID=?", twitterId)
	if err != nil {
		fmt.Printf("\n==========\nget_user_by_twitter - Problemz with selecting user by twitterID: \n%+v\n",err)
		err_msg = err
		fmt.Printf("\n==========\nget_user_by_twitter - Problemz with selecting user by twitterID: \n%+v\n",err_msg)
	} else {
		fmt.Printf("\n==========\nfound user: \n%+v\n",xpnkUser.User_ID)
	}
	return xpnkUser, err_msg
}