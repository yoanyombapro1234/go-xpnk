package get_user_by_insta

import (
	"fmt"
   	"xpnk-shared/db_connect"
)

func GetUserByInsta(insta_id string) (int, error) {
	dbmap := db_connect.InitDb()
	defer dbmap.Db.Close()
	
	var xpnkUser			int
	instaId					:= insta_id
	
	err	:= dbmap.SelectOne(&xpnkUser, "SELECT `user_ID` FROM USERS WHERE insta_userid=?", instaId)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			xpnkUser = 0
			fmt.Printf("\n==========\nNo user found: %+v\n",xpnkUser)
			err = nil
			return xpnkUser, err
		} else {
			fmt.Printf("\n==========\nSql error in finding user by insta id: %e", err)
		return xpnkUser, err
		} 
	}	
	
	fmt.Printf("\n==========\nUser ID: %+v\n",xpnkUser)
	return xpnkUser, err
}