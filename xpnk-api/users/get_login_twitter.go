package users

import (
	 "github.com/gin-gonic/gin"
	 "xpnk-user/xpnk_createUserObject"
	 "xpnk-user/xpnk_checkTwitterId"
)

func LoginTwitter (c *gin.Context) {
	var twitter_creds xpnk_createUserObject.User_Object
	var err error
	c.Bind(&twitter_creds)
	token 				:= twitter_creds.TwitterToken
	secret 				:= twitter_creds.TwitterSecret
	id 					:= twitter_creds.TwitterID
	twitter_user 		:= twitter_creds.TwitterUser
	twitter_avatar 		:= twitter_creds.ProfileImage
	if token == "" || secret == "" || id == "" || twitter_user == "" || twitter_avatar == ""{
		c.JSON(400, "User token, secret, user name, user id and profile image are required. One or all are missing.")
		return
	}
	
	user_groups, err := xpnk_checkTwitterId.CheckTwitterId(twitter_creds)
	if err !=  nil {
		c.JSON(400, err.Error())	
	} else {
		c.JSON(200, user_groups)
	}
}