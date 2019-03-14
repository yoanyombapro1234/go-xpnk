package get_user_object

import (
	"fmt"
   	"xpnk-shared/db_connect"
   	"xpnk-user/xpnk_createUserObject"
)

func GetUserObject(xpnk_id string) (xpnk_createUserObject.User_Object, error) {
	dbmap := db_connect.InitDb()
	defer dbmap.Db.Close()
	
	var xpnkUser			xpnk_createUserObject.User_Object
	xpnkId					:= xpnk_id
	
	err	:= dbmap.SelectOne(&xpnkUser, "SELECT * FROM USERS WHERE User_Id=?", xpnkId)
	if err != nil {
		fmt.Printf("\n==========\nget_user_object - Problemz with selecting user: \n%+v\n",err.Error())
	} else {
		fmt.Printf("\n==========\nfound user: \n%+v\n",xpnkUser)
	} 
	
	return xpnkUser, err
}