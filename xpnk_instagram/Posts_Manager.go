package main

/**************************************************************************************
Queen file of xpnk_instagram.
1) Queries the DB for all the instagram user ID's  
2) iterates through passing each one to getInstaUserPosts
3) prepares the posts for insertion into db (process includes fetching oembeds)
4) inserts Instagram posts into the db
**************************************************************************************/

import (
	"fmt"
	"database/sql"
   	_ "github.com/go-sql-driver/mysql"
   	"github.com/gopkg.in/gorp.v1"
    "log"
    "github.com/yanatan16/golang-instagram/instagram"
    "xpnk_instagram/xpnk_getInstaUserPosts"
    "xpnk_instagram/xpnk_createInstaInsert"
    "xpnk_instagram/xpnk_insertInsta"
)

//stores only the group_id of a group
type InstaID struct {
    InstaID	string		`db:"insta_userid"`
}

func main() {

	all_instagrammers := get_instaIDs()
	var instaUserPosts *instagram.PaginatedMediasResponse
	var insert []xpnk_createInstaInsert.Instagram_Insert
	
	for i := 0; i < len( all_instagrammers ); i++ {	
		if all_instagrammers[i].InstaID != "" {
			instaUserPosts = xpnk_getInstaUserPosts.GetInstaUserPosts( all_instagrammers[i].InstaID )
			insert = xpnk_createInstaInsert.CreateInstaInsert(instaUserPosts)
			xpnk_insertInsta.InsertInsta(insert)
		}	
	}
}//end main

func get_instaIDs() []InstaID {
	
	dbmap := initDb()
	defer dbmap.Db.Close()
	
	var insta_ids []InstaID
	
	//get all Instagram user ids from the USERS table	

	_,err := dbmap.Select(&insta_ids, "SELECT `insta_userid` FROM USERS")
	
	if err != nil {fmt.Printf("There was an error ", err)}
	
	checkErr(err, "Select failed")

	fmt.Printf("\n==========\nUSER INSTAGRAM IDS:%+v\n",insta_ids)
	
	return insta_ids
	
}//end Get_Groups  

/***************************
* db connection config
***************************/	
func initDb() *gorp.DbMap {
db, err := sql.Open("mysql",
	"root:root@tcp(localhost:8889)/xapnik")
checkErr(err, "sql.Open failed")

dbmap := &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{}}

return dbmap
}

func checkErr(err error, msg string) {
  if err != nil {
  log.Fatalln(msg, err)
	}	
} 