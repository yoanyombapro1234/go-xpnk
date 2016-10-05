package posts_manager

/**************************************************************************************
Queen file of xpnk_disqus.
1) Queries the DB for all the Disqus user ID's  
2) iterates through passing each one to getDisqusUserPosts
3) prepares the posts for insertion into db - createDisqusInsert
4) inserts Disqus posts into the db - insertDisqus
**************************************************************************************/

import (
	"fmt"
	"database/sql"
	"strconv"
   	_ "github.com/go-sql-driver/mysql"
   	"github.com/gopkg.in/gorp.v1"
    "log"
    "xpnk_disqus/golang-disqus/disqus"
    "xpnk_disqus/xpnk_getDisqusUserPosts"
    "xpnk_disqus/xpnk_createDisqusInsert"
    "xpnk_disqus/xpnk_insertDisqus"
    "xpnk_disqus/xpnk_deleteDisqus"
)

type MaxDisqusPID struct {
	MaxPID				sql.NullString		`db:"MAX(CAST(disqus_pid AS Unsigned))"`
}

func Get_posts() {
	
	fmt.Println("\n==========\nGetting all Disqus users.\n")
	all_disqusers := get_disqusUsernames()
	var disqusUserPosts []disqus.Content
	var insert []xpnk_createDisqusInsert.Disqus_Insert
	
	for i := 0; i < len( all_disqusers ); i++ {	
	//TODO eventually wrap this in if statement contingent upon non-empty accesstoken
		var this_disqus_user xpnk_getDisqusUserPosts.DisqusUser
		this_disqus_user.DisqusName = all_disqusers[i].DisqusName
		this_disqus_user.Disqus_accesstoken = all_disqusers[i].Disqus_accesstoken
					
		this_maxID := get_maxID(this_disqus_user.DisqusName)
		this_disqus_user.Disqus_maxID = this_maxID
		//collecting this but not using it since Disqus API is broken for 'since' param
		
		disqusUserPosts = xpnk_getDisqusUserPosts.GetDisqusUserPosts(this_disqus_user)
		insert = xpnk_createDisqusInsert.CreateDisqusInsert(disqusUserPosts)
		
		//before we update the Disqus posts table, we have to empty it for this user
		//why? because Disqus API 'since' param doesn't work 
		this_disqus_username := this_disqus_user.DisqusName
		xpnk_deleteDisqus.DeleteDisqus(this_disqus_username)
		
		xpnk_insertDisqus.InsertDisqus(insert)	
	}
}//end main

func get_disqusUsernames() []xpnk_getDisqusUserPosts.DisqusUser {
	
	dbmap := initDb()
	defer dbmap.Db.Close()
	dbmap.AddTableWithName(xpnk_getDisqusUserPosts.DisqusUser{}, "USERS")
	
	var disqus_usernames []xpnk_getDisqusUserPosts.DisqusUser
	
	//get all Disqus usernames and access tokens from the USERS table but no empty ones
	
	_,err := dbmap.Select(&disqus_usernames, "select disqus_accesstoken, disqus_username from USERS WHERE disqus_username != '' ") 
	//TODO right now we aren't requiring the disqus_accesstoken but we should in future
		
	checkErr(err, "Select failed")

	fmt.Printf("\n==========\nDISQUS USERS:%+v\n",disqus_usernames)
	
	return disqus_usernames
	
}//end get_disqusUsernames  

func get_maxID(disqus_username string) string {
	fmt.Println("\n==========\nStarting get_maxID...\n")
	var disqus_max_id MaxDisqusPID
	
	dbmap := initDb()
	defer dbmap.Db.Close()
	//dbmap.AddTableWithName(disqus_max_id, "disqus_comments")
	
	err := dbmap.SelectOne(&disqus_max_id, "select MAX(CAST(disqus_pid AS Unsigned)) from disqus_comments where disqus_user = ?", disqus_username)
	
	checkErr(err, "Select failed:  ")
	
	fmt.Printf("\ndisqus_max_id is %+v\n", disqus_max_id)

	
	var max_id string
	
	if disqus_max_id.MaxPID.String != "" {
		maxpid, err := strconv.ParseInt(disqus_max_id.MaxPID.String, 10, 64)
		int_id := maxpid
		fmt.Printf("\nHere's our max_id + 1 %v\n", int_id)
		max_id = strconv.FormatInt(int_id, 10)
		fmt.Printf("\nmax_id = %v\n", max_id)
		fmt.Print("Errors: %v", err)
	} else {
		max_id = ""
		fmt.Printf("\nmax_id is an empty string\n")
	}	

	
	return max_id
}

/***************************
* db connection config
***************************/	
func initDb() *gorp.DbMap {
db, err := sql.Open("mysql",
	"")
checkErr(err, "sql.Open failed")

dbmap := &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{}}

return dbmap
}

func checkErr(err error, msg string) {
  if err != nil {
  log.Fatalln(msg, err)
	}	
} 