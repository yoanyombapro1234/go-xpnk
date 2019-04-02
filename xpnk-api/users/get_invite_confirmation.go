package users

import (
	"fmt"
	 _ "github.com/go-sql-driver/mysql"
	 "github.com/gin-gonic/gin"
	 "xpnk-user/xpnk_checkUserInvite"
)

type NewUserInvite 		struct {
	Xpnk_token			string					`form:"xpnk_token" binding:"required"`
	Group_name			string					`form:"xpnk_group_name" binding:"required"`
}

func CheckUserInvite (c *gin.Context) {
	var user_invite			NewUserInvite
	var user_invite_check	xpnk_checkUserInvite.GroupObj
	c.Bind(&user_invite)
	fmt.Printf("CheckUserInvite Xpnk_token:  %v \n", user_invite.Xpnk_token)
	fmt.Printf("CheckUserInvite Group_name:  %v \n", user_invite.Group_name)
	
	user_invite_check		= xpnk_checkUserInvite.CheckUserInvite(user_invite.Xpnk_token, user_invite.Group_name)
	if user_invite_check.GroupName == user_invite.Group_name {
		c.JSON(201, user_invite_check)
	} else {
		c.JSON(400, user_invite_check)
	}
}