package xpnk_twitter

//gets all the group ID's from the Groups table and puts into a struct for use by other routines

import (
	"fmt"
	"log"
   	_ "github.com/go-sql-driver/mysql"
)

//stores only the group_id of a group
type GroupID struct {
    GroupID	int		`db:"Group_ID"`
}

func Get_Groups () []GroupID {
	
	dbmap := initDb()
	defer dbmap.Db.Close()
	
	var group_ids []GroupID
	
	//get all group ids from the GROUPS table	

	_,err := dbmap.Select(&group_ids, "SELECT `Group_ID` FROM groups")
	
	if err != nil {fmt.Printf("There was an error ", err)}
	
	checkErr(err, "Select failed")

	fmt.Printf("\n==========\nGROUP IDS:%+v\n",group_ids)
	
	return group_ids
	
}//end Get_Groups  

func checkErr(err error, msg string) {
  if err != nil {
  log.Fatalln(msg, err)
	}	
} 
