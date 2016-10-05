package xpnk_addUsersToGroup

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
	User				int				`db:"user_ID"`
	Group				int				`db:"Group_ID"`
}

func AddUsersToGroup(group int, users []int) string{

	fmt.Printf("\n==========\nUsers is now:%+v\n",users)
	
	var thisGroup []Group_User
	
	for i :=0; i <len(users); i++ {
		var thisUser Group_User
		thisUser.User = users[i]
		thisUser.Group = group
		thisGroup = append(thisGroup, thisUser)
	}
	
	fmt.Printf("\n==========\nThisGroup is now:%+v\n",thisGroup)


	//Initialize a map variable to hold all our Group_User structs
	var set map[int]Group_User

	dbmap := db_connect.InitDb()
	defer dbmap.Db.Close()
	
	//map the Group_User struct to the 'USER_GROUPS' table
	dbmap.AddTableWithName(Group_User{}, "USER_GROUPS")
	
	//Create the map that will contain all our structs
	set = make(map[int]Group_User)
		
	for i := 0; i < len(thisGroup); i++ {
		set[i] = thisGroup[i]
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
