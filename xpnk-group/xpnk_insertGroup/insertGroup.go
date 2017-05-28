package xpnk_insertGroup

/**************************************************************************************
Takes a Group_Insert object and inserts it into groups table
**************************************************************************************/

import (
	"fmt"
   	_ "github.com/go-sql-driver/mysql"
   	"xpnk-group/xpnk_createGroupInsert"
   	"xpnk-shared/db_connect"
)

var	groupID	int
var return_val string

func InsertGroup(group xpnk_createGroupInsert.Group_Insert) string{

	//Initialize a map variable to hold all our Group_Insert structs
	//var set map[int]xpnk_createGroupInsert.Group_Insert

	dbmap := db_connect.InitDb()
	defer dbmap.Db.Close()
	
//map Group_Insert struct to the 'groups' db table
	dbmap.AddTableWithName(xpnk_createGroupInsert.Group_Insert{}, "GROUPS")
	
	err := dbmap.SelectOne(&groupID, "SELECT Group_ID FROM GROUPS WHERE group_name=?", group.GroupName)
	if err != nil {fmt.Printf("There was an error ", err)}
	if groupID > 0 {
		fmt.Printf("Group already exists: ", groupID)
		return_val = "group already exists"
	} else {
		
	//Insert the the Group		
		err = dbmap.Insert(&group)
		if err != nil {fmt.Printf("There was an error ", err)
	
		}
		return_val = "inserted"	
	}		
	return return_val
}
//end db insert routine	
//Maps are not inherently safe for concurrency - will have to use sync.RWMutex	
