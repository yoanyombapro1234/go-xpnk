package users

import (
	 "github.com/gin-gonic/gin"
	 "xpnk-user/xpnk_createUserObject"
	 "xpnk-user/xpnk_checkInstaId"
)

func LoginInsta (c *gin.Context) {
	var insta_creds xpnk_createUserObject.User_Object
	var err error
	c.Bind(&insta_creds)
	token 				:= insta_creds.InstaAccessToken
	//secret 			:= insta_creds.TwitterSecret
	id 					:= insta_creds.InstaUserID
	insta_user 			:= insta_creds.InstaUser
	insta_avatar 		:= insta_creds.ProfileImage
	if token == "" /*|| secret == ""*/ || id == "" || insta_user == "" || insta_avatar  == ""{
		c.JSON(400, "User token, secret, user name, user id and profile image are required. One or all are missing.")
		return
	}
	
	user_groups, err := xpnk_checkInstaId.CheckInstaId(insta_creds)
	if err !=  nil {
		c.JSON(400, err.Error())	
	} else {
		c.JSON(200, user_groups)
	}
}