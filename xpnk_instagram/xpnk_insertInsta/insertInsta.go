package xpnk_insertInsta

/**************************************************************************************
Takes a slice of Instagram_Insert objects and inserts them into the database
**************************************************************************************/

import (
    "fmt"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
    "github.com/gopkg.in/gorp.v1"
    "xpnk_instagram/xpnk_createInstaInsert"
    "log"
)

func InsertInsta(instagramposts []xpnk_createInstaInsert.Instagram_Insert) string{

	//delete posts older than 24 hours, func from delete-old-instagrams.go
	//defer Dodelete() 

	//Initialize a map variable to hold all our Instagram_Insert structs (posts)
	var set map[int]xpnk_createInstaInsert.Instagram_Insert

	dbmap := initDb()
	defer dbmap.Db.Close()
	
//map the []Instagram_Insert struct to the 'instagram_posts' db table
	dbmap.AddTableWithName(xpnk_createInstaInsert.Instagram_Insert{}, "instagram_posts")
	
//Create the map that will contain all our structs from Posts
	set = make(map[int]xpnk_createInstaInsert.Instagram_Insert)
		
	for i := 0; i < len(instagramposts); i++ {
		set[i] = instagramposts[i]
	}
		
	//Insert the the Instagrams!	
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
