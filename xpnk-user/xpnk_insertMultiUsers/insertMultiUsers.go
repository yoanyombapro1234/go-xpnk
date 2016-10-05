package xpnk_insertMultiUsers

/**************************************************************************************
Takes a slice of User_Insert objects and inserts them into USERS table
**************************************************************************************/

import (
	"xpnk-user/xpnk_createUserInsert"
   	"xpnk-shared/db_connect"
	"fmt"
   	"log"
)

func InsertMultiUsers(users []xpnk_createUserInsert.User_Insert) string{

	//Initialize a map variable to hold all our User_Insert structs
	var set map[int]xpnk_createUserInsert.User_Insert

	dbmap := db_connect.InitDb()
	defer dbmap.Db.Close()
	
//map the []xpnk_createUserInsert.User_Insert struct to the 'USERS' db table
	dbmap.AddTableWithName(xpnk_createUserInsert.User_Insert{}, "USERS")
	
//Create the map that will contain all our structs
	set = make(map[int]xpnk_createUserInsert.User_Insert)
		
	for i := 0; i < len(users); i++ {
		set[i] = users[i]
	}
		
	fmt.Printf("\n==========\nset is now:%+v\n",set)
	
	
	//Insert the users!	
	for _, v := range set {	
			
		//db insert function 
		err := dbmap.Insert(&v)
		if err != nil {fmt.Printf("There was an error ", err)
			
		}
	}	
		return "inserted"	
}
//end db insert routine	
//Maps are not inherently safe for concurrency - will have to use sync.RWMutex	

func checkErr(err error, msg string) {
  if err != nil {
  log.Fatalln(msg, err)
	}	
} 