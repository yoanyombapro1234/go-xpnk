package users

//TODO RETURN ERROR WHEN USER ID DOESN'T EXIST

import (
	"fmt"
	"reflect"
	 "github.com/gin-gonic/gin"
	 "xpnk-user/xpnk_createUserObject"
	 "xpnk-user/xpnk_updateUser"
)

func UsersUpdate_2(c *gin.Context) {

	var thisUser xpnk_createUserObject.User_Object
	var err_msg error
	c.Bind(&thisUser)
	if thisUser.InstaUserID == "" && thisUser.TwitterID == "" {
		fmt.Printf("You must send either an InstaUserID or TwitterID param, or both. If you passed an int value for either, please change it to a string.")
		c.JSON(400, "You must send either an InstaUserID or TwitterID param, or both. If you passed an int value for either, please change it to a string.")	
		return
	}
		
	if (reflect.TypeOf(thisUser.InstaUserID).String()) != "string" || (reflect.TypeOf(thisUser.TwitterID).String()) != "string" {
		c.JSON(400, "TwitterID and InstaUserID values must be strings. Please check.")
		return
	}
	
	update_thisUser, err_msg 	:=  xpnk_updateUser.UpdateUser_2(thisUser)
	if err_msg != nil {
		fmt.Printf(err_msg.Error())
		c.JSON(400, err_msg.Error())
		return
	}
	
	if update_thisUser == 1 {
		fmt.Printf("\nthisUser updated:  %+v \n ID: %n\n", thisUser.Id)
		c.JSON(200, thisUser)
	}
}