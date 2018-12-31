package xpnk_addUsersToGroup2

/**************************************************************************************
*
*Takes a slice of xpnk_userID's and a Group ID. For each userID inserts userID 
*and Group ID into USER_GROUPS table.
*
**************************************************************************************/

import (
   	"xpnk-shared/db_connect"
	"fmt"
   	"log"
)

type Group_User struct {
	User				string			`db:"user_ID"`
	Group				int				`db:"Group_ID"`
	Owner				bool			`db:"group_owner"`
	Admin				bool			`db:"group_admin"`
}

func AddUsersToGroup(user_group Group_User) (Group_User, error) {

	fmt.Printf("\n==========\nUser is now:%+v\n", user_group)
		
	dbmap := db_connect.InitDb()
	defer dbmap.Db.Close()
	
	//map the xpnk_createGroupInsert.Group_Insert struct to the 'USER_GROUPS' db table
	dbmap.AddTableWithName(Group_User{}, "USER_GROUPS")
		
	//db insert function 
	err := dbmap.Insert(&user_group)
	if err != nil {
		fmt.Printf("There was an error ", err)
		return user_group, err
	}
	
	fmt.Printf("\n==========\nNEW GROUP MEMBER INSERTED: \n%+v\n", user_group)
				
	return user_group, err
}
//end db insert routine	
//Maps are not inherently safe for concurrency - will have to use sync.RWMutex	

func checkErr(err error, msg string) {
  if err != nil {
  log.Fatalln(msg, err)
	}	
} 
