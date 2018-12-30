package xpnk_insertGroup2

/**************************************************************************************
Takes a Group_Insert object and inserts it into groups table
Returns the new group's Group_ID
**************************************************************************************/

import (
	"fmt"
   	_ "github.com/go-sql-driver/mysql"
   	"xpnk-group/xpnk_createGroupInsert"
   	"xpnk-shared/db_connect"
)

var	groupID	int
var return_val int

func InsertGroup2(group xpnk_createGroupInsert.Group_Insert) (int, error){

	var err_msg				error
	
	dbmap := db_connect.InitDb()
	defer dbmap.Db.Close()
	
//map the xpnk_createGroupInsert.Group_Insert struct to the 'GROUPS' db table
	dbmap.AddTableWithName(xpnk_createGroupInsert.Group_Insert{}, "GROUPS").SetKeys(true, "Group_ID")
	
	err := dbmap.SelectOne(&groupID, "SELECT Group_ID FROM GROUPS WHERE group_name=?", group.GroupName)
	if err != nil {fmt.Printf("There was an error ", err)}
	if groupID > 0 {
		fmt.Printf("Group already exists: ", groupID)
		return_val = 0
	} else {
		err 			:= 	dbmap.Insert(&group)
		fmt.Printf("\nThe new group id is: %v\n", group.GroupID)
		return_val 		= group.GroupID
		if err != nil {
			err_msg 	= err
			fmt.Printf("There was an error ", err)			
		} 
	}	
	return return_val, err_msg		 
}

//end db insert routine	
//Maps are not inherently safe for concurrency - will have to use sync.RWMutex	
