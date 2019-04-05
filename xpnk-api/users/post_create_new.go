package users

import (
	"fmt"
	 _ "github.com/go-sql-driver/mysql"
	 "github.com/gin-gonic/gin"
	 "xpnk-user/xpnk_createUserObject"
	 "xpnk-user/xpnk_createUserInsert"
	 "xpnk-user/xpnk_insertMultiUsers"
)

func UsersNew_2 (c *gin.Context) {
	var newUser					xpnk_createUserObject.User_Object
	var err_msg					error
	c.Bind(&newUser)
	fmt.Printf("newUser to add:  %+v \n", newUser)
	if newUser.TwitterID == "" && newUser.InstaUserID == "" {
		c.JSON(400, "Need either a Twitter user ID or a Instagram user ID to create a new user.")
		return
	}
	var userInsert				[]xpnk_createUserObject.User_Object
	userInsert 				 =  append(userInsert, newUser)
	
	newID, err_msg 			:=  xpnk_insertMultiUsers.InsertMultiUsers_2(userInsert)
	if err_msg != nil {
		c.JSON(400, err_msg.Error())	
	} else {
		c.JSON(200, newID)
	}
}

func UsersNew (c *gin.Context) {
	var newUser					xpnk_createUserInsert.User_Insert
	c.Bind(&newUser)
	fmt.Printf("newUser to add:  %+v \n", newUser)
	var userInsert				[]xpnk_createUserInsert.User_Insert
	userInsert 				 =  append(userInsert, newUser)
	
	insertUser 				:=  xpnk_insertMultiUsers.InsertMultiUsers(userInsert)
	if insertUser == "inserted" {
		c.JSON(200, "New user inserted.")
	}	else {
		c.JSON(202, "New user was not inserted. Check the API logs.")
	}		
}