package instagram_json_workmanager

/**************************************************************************************
Queries the DB for all the group ID's and iterates through passing each one to Insta_json
**************************************************************************************/

import (
	"fmt"
	"database/sql"
   	_ "github.com/go-sql-driver/mysql"
   	"github.com/gopkg.in/gorp.v1"
    "log"
    "xpnk_instagram/xpnk_createInstaJSON"
)

//stores only the group_id of a group
type GroupID struct {
    GroupID	int		`db:"Group_ID"`
}

func Create_group_ig_json() {

	all_groups := get_groups()

	for i := 0; i < len(all_groups); i++ {	
		xpnk_createInstaJSON.CreateInstaJSON( all_groups[i].GroupID )
	}
}//end main

func get_groups() []GroupID {
	
	dbmap := initDb()
	defer dbmap.Db.Close()
	
	var group_ids []GroupID
	
	//get all group ids from the GROUPS table	

	_,err := dbmap.Select(&group_ids, "SELECT `Group_ID` FROM groups")
	
	if err != nil {fmt.Printf("There was an error ", err)}
	
	checkErr(err, "Select failed")
	
	return group_ids
	
}//end Get_Groups  

/***************************
* db connection config
***************************/	
func initDb() *gorp.DbMap {
db, err := sql.Open("mysql",
	"root:root@tcp(localhost:8889)/password")
checkErr(err, "sql.Open failed")

dbmap := &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{}}

return dbmap
}

func checkErr(err error, msg string) {
  if err != nil {
  log.Fatalln(msg, err)
	}	
} 