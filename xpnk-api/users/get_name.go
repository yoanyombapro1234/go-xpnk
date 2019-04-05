package users

import (
	 _ "github.com/go-sql-driver/mysql"
	 "xpnk-shared/db_connect"
)

func GetUserName(user_id string) (string, error){
	dbmap := db_connect.InitDb()
	defer dbmap.Db.Close()
	
	type UserNames struct {
		Twitter_user		string		   `db:"twitter_user"`		
		Insta_user			string		   `db:"insta_user"`			
	} 
	
	var userNames			UserNames
	var err_msg				error
	var user_name			string
	id						:= user_id
	
	err	:= dbmap.SelectOne(&userNames, "SELECT `twitter_user`, `insta_user` FROM USERS WHERE `user_ID`=?", id)
			
	if err != nil {
		err_msg				= err
	} 
	
	if userNames.Twitter_user != "" {
		user_name = userNames.Twitter_user
	} else { user_name = userNames.Insta_user } 
	
	return user_name, err_msg
}