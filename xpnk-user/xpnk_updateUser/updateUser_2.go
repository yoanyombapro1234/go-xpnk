package xpnk_updateUser

/**************************************************************************************
Takes a slice of User_Update objects and inserts them into USERS table
**************************************************************************************/

import (
   	"xpnk-shared/db_connect"
   	"xpnk-user/xpnk_createUserObject"
	"fmt"
   	//"log" TODO: add this back after retiring v.1 of this function
)

func UpdateUser_2(userupdate xpnk_createUserObject.User_Object) (int64, error){

	var err_msg				error

	fmt.Printf("\n==========\nuserupdate: \n%+v\n", userupdate)

	dbmap := db_connect.InitDb()
	defer dbmap.Db.Close()
	
//map the User_Update struct to the 'USERS' db table
	dbmap.AddTableWithName(xpnk_createUserObject.User_Object{}, "USERS").SetKeys(true, "user_ID")
	
	count, err := dbmap.Update(&userupdate)
	if err != nil {
		err_msg 	= err
		fmt.Printf("\n==========\nUPDATE DIDN'T WORK: \n%+v\n", err)
		checkErr(err, "\nupdateUser couldn't update the db:")
		return 0, err_msg
	} else {
		fmt.Printf("\n==========\nuserupdate returned: \n%+v\n", count)
		return 1, err_msg
	}
} 

/*
func checkErr(err error, msg string) {
  if err != nil {
  log.Fatalln(msg, err)
	}	
} 
TODO: add this back after retiring version 1 of this function
*/
