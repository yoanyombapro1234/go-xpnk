package users

import (
	"fmt"
	 "github.com/gin-gonic/gin"
	 "xpnk-shared/db_connect"
)

type NewTwitterAuth struct {
	 Twttr_userid		string					`json:"twitter_userid"`
}

func PostTwttrAuth(c *gin.Context) {
	var twitter_user			NewTwitterAuth
	c.Bind(&twitter_user)
	twitterId					:= twitter_user.Twttr_userid
	fmt.Printf("twitterId:  %v \n", twitterId)
	if twitterId == ""{
		c.JSON(422, gin.H{"error": "Invalid or missing Twitter user ID."})
	} else {
	
		twitter_user_check		:= check_twitter_id(twitterId)
		if twitter_user_check != 0 {
			fmt.Printf("\nXPNKUSERID RETURNED BY check_twitter_id: %+v\n", twitter_user_check)
			c.JSON(201, twitter_user_check)
		} else {
			c.JSON(422, gin.H{"error": twitter_user_check })
		}		 
	}	
}

func check_twitter_id(twitter_id string) int {
	dbmap := db_connect.InitDb()
	defer dbmap.Db.Close()
	twitterId				:= twitter_id
	var user_xpnkid 		int
	var xpnkid				int
	err	:= dbmap.SelectOne(&user_xpnkid, "SELECT user_ID FROM USERS WHERE twitter_ID=?", twitterId)
	if err == nil {
		xpnkid 				= user_xpnkid
	} else {
		fmt.Printf("\n==========\ncheck_twitter_id - Problemz with selecting user_xpnkid: \n%+v\n",err)
		xpnkid = 0
	}
	return xpnkid
}