package users

import (
	"fmt"
	 _ "github.com/go-sql-driver/mysql"
	 "github.com/gin-gonic/gin"
	 "xpnk-user/xpnk_user_structs"
	 "xpnk-shared/db_connect"
)

func SaveSub (c *gin.Context) {
	var newSub					xpnk_user_structs.UserSub
	var err_msg					error
	c.Bind(&newSub)
	fmt.Printf("newSub to add:  %+v \n", newSub)
	if newSub.Id < 1 || newSub.Endpoint == "" || newSub.Type < 0 || newSub.P256dh == "" || newSub.Auth == "" {
		c.JSON(400, "All fields must have a value - Id, Endpoint, Type, P256dh, Auth. One or more is empty.")
		return
	}
	
	subAdded, err_msg 			:=  insertSub(newSub)
	if err_msg != nil {
		c.JSON(400, err_msg.Error())	
	} else {
		c.JSON(200, subAdded)
	}
}

func insertSub (sub xpnk_user_structs.UserSub) (int, error) {
	dbmap 					:= db_connect.InitDb()
	defer dbmap.Db.Close()
	
	dbmap.AddTableWithName(xpnk_user_structs.UserSub{}, "user_subs")
		
	err := dbmap.Insert(&sub)
		fmt.Printf("\nNew sub added for user id: %v\n", sub.Id)
		newSubUserID := sub.Id
		if err != nil {
			fmt.Printf("There was an error ", err)			
		}
	return newSubUserID, err	

}