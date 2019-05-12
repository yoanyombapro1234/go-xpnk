package users

import (
	"fmt"
	"strconv"
	 _ "github.com/go-sql-driver/mysql"
	 "github.com/gin-gonic/gin"
	 "xpnk-user/xpnk_user_structs"
	 "xpnk-shared/db_connect"
)

func SaveFireBaseKey (c *gin.Context) {
	var newSub					xpnk_user_structs.UserSub
	var err_msg					error
	user_id						:= c.Params.ByName("id")
	user_id_string, err			:= strconv.Atoi(user_id)
	newSub.Id					= user_id_string	
	newSub.FirebaseKey			=	c.Params.ByName("key")
	newSub.Type					= 2
	if newSub.Id < 1 || err != nil || newSub.FirebaseKey == "" {
		c.JSON(400, "Missing either the user id or the Firebase key, or both.")
		return
	}
	
	keyAdded, err_msg 			:=  insertFireBaseKey(newSub)
	if err_msg != nil {
		c.JSON(400, err_msg.Error())	
	} else {
		c.JSON(200, keyAdded)
	}
}

func insertFireBaseKey (sub xpnk_user_structs.UserSub) (int, error) {
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