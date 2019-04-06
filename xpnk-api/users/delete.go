package users

import (
	"fmt"
	"strconv"
	 "github.com/gin-gonic/gin"
	 "xpnk-shared/db_connect"
)

func UsersDelete (c *gin.Context) {
	userid, err 			:= 	strconv.Atoi(c.Params.ByName("id"))
	if err != nil {
		c.JSON( 400, err.Error())
		return
	}
	
	if userid <= 0 {
	  	c.JSON(422, gin.H{"error": "No User_id was sent."})
	  	return
	} else {
		userdel, err 	:= delUser(userid)
		if err != nil {
			 fmt.Printf("\nERROR DELETING USER: %+v\n", err)
			c.JSON(400, err.Error())
			return
		} else {
			fmt.Printf("\nUSER DELETED: %+v\n", userdel)	
			returnstring := "User deleted: " + c.Params.ByName("id")
			c.JSON(201, returnstring)
		}	
	}		 
}

func delUser (userID int) (int64, error) {
	type User struct {
		User_ID 			int 			`db:"user_ID"`
	}
	
	var user_id 			User 
	user_id.User_ID = userID 
	fmt.Printf("\n==============\n User_ID to be deleted: %+v", user_id.User_ID)
	
	dbmap 					:= db_connect.InitDb()
	defer dbmap.Db.Close()
	
	dbmap.AddTableWithName(User{}, "USERS").SetKeys(true, "user_ID")
	
	_, err := dbmap.Delete(&user_id)
	fmt.Printf("\n==============\n deleted: %+v", user_id)
	
	count, err := dbmap.SelectInt("select count(*) from USERS where user_ID=?", user_id.User_ID)
	fmt.Printf("\n==============\n COUNT: %+v", count)
	
	res, err2 := DelUserGroups(user_id.User_ID)
	if err2 != nil {
		fmt.Printf("\n===========\n delUserGroups error: %+v", err)
	} else {
		fmt.Printf("\n===========\n delUserGroups response: %+v", res)
	}
	
	return count, err
}