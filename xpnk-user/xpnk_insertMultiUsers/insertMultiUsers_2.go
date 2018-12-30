package xpnk_insertMultiUsers

/**************************************************************************************
Takes a slice of User_Insert objects and inserts them into USERS table
Returns the new user ID
**************************************************************************************/

import (
	"xpnk-user/xpnk_createUserObject"
   	"xpnk-shared/db_connect"
	"fmt"
)

func InsertMultiUsers_2(users []xpnk_createUserObject.User_Object) (int, error) {

	//Initialize a map variable to hold all our User_Object
	var err_msg				error
	var set map[int]xpnk_createUserObject.User_Object

	dbmap := db_connect.InitDb()
	defer dbmap.Db.Close()
	
//map the []xpnk_createUserObject.User_Object struct to the 'USERS' db table
	dbmap.AddTableWithName(xpnk_createUserObject.User_Object{}, "USERS").SetKeys(true, "Id")
	
//Create the map that will contain all our structs
	set = make(map[int]xpnk_createUserObject.User_Object)
		
	for i := 0; i < len(users); i++ {
		set[i] = users[i]
	}
		
	fmt.Printf("\n==========\nset is now:%+v\n",set)
	
	
	//Insert the users! 
	var newUserID int
	for _, v := range set {	
			
		//db insert function 
		err := dbmap.Insert(&v)
		fmt.Printf("\nThe new user id is: %v\n", v.Id)
		newUserID = v.Id
		if err != nil {
			err_msg 	= err
			fmt.Printf("There was an error ", err)			
		}
	}	
		return newUserID, err_msg	
}
//end db insert routine	
//Maps are not inherently safe for concurrency - will have to use sync.RWMutex	