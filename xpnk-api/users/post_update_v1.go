package users

import (
	"fmt"
	 "github.com/gin-gonic/gin"
	 "xpnk-user/xpnk_updateUser"
)

func UsersUpdate (c *gin.Context) {
	var thisUser				xpnk_updateUser.User_Update
	c.BindJSON(&thisUser)
	fmt.Printf("thisUser to add:  %+v \n", thisUser)
	update_thisUser 			:=  xpnk_updateUser.UpdateUser(thisUser)
	if update_thisUser == 1 {
		c.JSON(200, "User updated.")
	}	else {
		c.JSON(202, "User not updated. Check the API logs.")
	}	
}