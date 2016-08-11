package posts_manager

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
	"strconv"
   	_ "github.com/go-sql-driver/mysql"
   	"github.com/gopkg.in/gorp.v1"
    "log"
    "github.com/yanatan16/golang-instagram/instagram"
    "xpnk_instagram/xpnk_getInstaUserPosts"
    "xpnk_instagram/xpnk_createInstaInsert"
    "xpnk_instagram/xpnk_insertInsta"
    "xpnk_instagram/xpnk_deleteInsta"
)

type MaxIGPID struct {
	MaxPID				sql.NullString		`db:"MAX(CAST(instagram_pid AS Unsigned))"`
}

func Get_posts() {
	
	fmt.Println("\n==========\nGetting all instagrammers.\n")
	all_instagrammers := get_instaIDs()
	var instaUserPosts *instagram.PaginatedMediasResponse
	var insert []xpnk_createInstaInsert.Instagram_Insert
	
	for i := 0; i < len( all_instagrammers ); i++ {	
		if all_instagrammers[i].Insta_accesstoken != "" {
			var this_insta_user xpnk_getInstaUserPosts.InstaUser
			this_insta_user.InstaID = all_instagrammers[i].InstaID
			this_insta_user.Insta_accesstoken = all_instagrammers[i].Insta_accesstoken
						
			this_maxID := get_maxID(this_insta_user.InstaID)
			this_insta_user.Insta_maxID = this_maxID
			//collecting this but not using it since IG API is broken for this param
			
			instaUserPosts = xpnk_getInstaUserPosts.GetInstaUserPosts( this_insta_user )
			insert = xpnk_createInstaInsert.CreateInstaInsert(instaUserPosts)
			
			//before we update the IG posts table, we have to empty it for this user
			//why? because IG API doesn't work properly and sends duplicates
			this_igid := all_instagrammers[i].InstaID
			xpnk_deleteInsta.DeleteInsta(this_igid)
			
			xpnk_insertInsta.InsertInsta(insert)
		} else {
			fmt.Println("\n==========\nNo Instagram token yet.\n")
		}	
	}
}//end main

func get_instaIDs() []xpnk_getInstaUserPosts.InstaUser {
	
	dbmap := initDb()
	defer dbmap.Db.Close()
	dbmap.AddTableWithName(xpnk_getInstaUserPosts.InstaUser{}, "USERS")
	
	var insta_ids []xpnk_getInstaUserPosts.InstaUser
	
	//get all Instagram user ids and access tokens from the USERS table	but no empty ones
	
	_,err := dbmap.Select(&insta_ids, "select insta_accesstoken, insta_userid from USERS WHERE insta_userid != ''  && insta_accesstoken != '' ")
		
	checkErr(err, "Select failed")
	
	return insta_ids
	
}//end get_instaIDs  

func get_maxID(iguser_id string) string {
	var insta_max_id MaxIGPID
	
	dbmap := initDb()
	defer dbmap.Db.Close()
	
	err := dbmap.SelectOne(&insta_max_id, "select MAX(CAST(instagram_pid AS Unsigned)) from instagram_posts where insta_userid = ?", iguser_id)
	
	checkErr(err, "Select failed:  ")
	
	var max_id string
	
	if insta_max_id.MaxPID.String != "" {
		maxpid, err := strconv.ParseInt(insta_max_id.MaxPID.String, 10, 64)
		int_id := maxpid + 1
		max_id = strconv.FormatInt(int_id, 10)
		//max_id += "_" + iguser_id
	} else {
		max_id = ""
	}	
	
	return max_id
}

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