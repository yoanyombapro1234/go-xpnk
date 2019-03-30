package xpnk_updateUser

/**************************************************************************************
Takes a slice of User_Update objects and inserts them into USERS table
**************************************************************************************/

import (
	"strconv"
	"fmt"
	"bytes"
	"strings"
	 "xpnk-shared/db_connect"
   	"xpnk-user/xpnk_createUserObject"
   	//"log" TODO: add this back after retiring v.1 of this function
)

func UpdateUser_2(userupdate xpnk_createUserObject.User_Object) (int64, error) {

	var err_msg				error

	fmt.Printf("\n==========\nuserupdate: \n%+v\n", userupdate)

	dbmap := db_connect.InitDb()
	defer dbmap.Db.Close()
	
update_sql_command_string := bytes.Buffer{}
//sql_vars := ""

if userupdate.TwitterUser != "" { 
	update_sql_command_string.WriteString(fmt.Sprintf("`twitter_user`='%s', ", userupdate.TwitterUser)) 
}		
if userupdate.TwitterID != "" { 
	update_sql_command_string.WriteString(fmt.Sprintf("`twitter_ID`='%s', ", userupdate.TwitterID)) 
}				
if userupdate.TwitterToken != "" { 
	update_sql_command_string.WriteString(fmt.Sprintf("`twitter_authtoken`='%s', ", 
	userupdate.TwitterToken)) 
}
if userupdate.TwitterSecret != "" { 
	update_sql_command_string.WriteString(fmt.Sprintf("`twitter_secret`='%s', " ,
	userupdate.TwitterSecret)) 
}
if userupdate.InstaUser != "" { 
	update_sql_command_string.WriteString(fmt.Sprintf("`insta_user`='%s', ", 
	 userupdate.InstaUser)) 
}
if userupdate.InstaUserID != "" { 
	update_sql_command_string.WriteString(fmt.Sprintf("`insta_userid`='%s', ", 
	userupdate.InstaUserID)) 
}
if userupdate.InstaAccessToken != "" { 
	update_sql_command_string.WriteString(fmt.Sprintf("`insta_accesstoken`='%s', ", 
	userupdate.InstaAccessToken)) 
}
if userupdate.DisqusUserName != "" { 
	update_sql_command_string.WriteString(fmt.Sprintf("`disqus_username`='%s', ", 
	userupdate.DisqusUserName)) 
}
if userupdate.DisqusUserID != "" { 
	update_sql_command_string.WriteString(fmt.Sprintf("`disqus_userid`='%s', ", 
	userupdate.DisqusUserID)) 
}
if userupdate.DisqusAccessToken != "" { 
	update_sql_command_string.WriteString(fmt.Sprintf("`disqus_accesstoken`='%s', ", 
	userupdate.DisqusAccessToken)) 
}
if userupdate.DisqusRefreshToken != "" { 
	update_sql_command_string.WriteString(fmt.Sprintf("`disqus_refreshtoken`='%s', ", 
	userupdate.DisqusRefreshToken)) 
}
if userupdate.ProfileImage != "" { 
	update_sql_command_string.WriteString(fmt.Sprintf("`profile_image`='%s', ", 
	userupdate.ProfileImage)) 
}
	

if userupdate.Id > 0 {
	id := strconv.Itoa(userupdate.Id)
	update_sql_command_string.WriteString(fmt.Sprintf(" WHERE `user_ID`=%s", id ))
}		

update_user := "UPDATE USERS SET " + update_sql_command_string.String() + ";"

update_user = strings.Replace(update_user, ",  WHERE", " WHERE", -1)

fmt.Printf("update_user: %v\n", update_user)
	
//map the User_Update struct to the 'USERS' db table

	res, err := dbmap.Exec(update_user)
	
/*	
	count, err := dbmap.Update(&userupdate)
*/
	
	if err != nil {
		err_msg 	= err
		fmt.Printf("\n==========\nUPDATE DIDN'T WORK: \n%+v\n", err)
		checkErr(err, "\nupdateUser couldn't update the db:")
		return 0, err_msg
	} else {
		fmt.Printf("\n==========\nuserupdate returned: \n%+v\n", res)
		return 1, err_msg
	}


return 1, err_msg
} 

/*
func checkErr(err error, msg string) {
  if err != nil {
  log.Fatalln(msg, err)
	}	
} 
TODO: add this back after retiring version 1 of this function
*/
