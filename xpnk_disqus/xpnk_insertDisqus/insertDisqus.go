package xpnk_insertDisqus

/**************************************************************************************
Takes a slice of Disqus_Insert objects and inserts them into disqus_comments table
**************************************************************************************/

import (
	"fmt"
    "database/sql"
   	_ "github.com/go-sql-driver/mysql"
   	"github.com/gopkg.in/gorp.v1"
   	"xpnk_disqus/xpnk_createDisqusInsert"
   	"log"
)

func InsertDisqus(disqusposts []xpnk_createDisqusInsert.Disqus_Insert) string{

	//delete posts older than 24 hours, func from delete-old-instagrams.go
	//defer Dodelete() 

	//Initialize a map variable to hold all our Instagram_Insert structs (posts)
	var set map[int]xpnk_createDisqusInsert.Disqus_Insert

	dbmap := initDb()
	defer dbmap.Db.Close()
	
//map the []Disqus_Insert struct to the 'disqus_comments' db table
	dbmap.AddTableWithName(xpnk_createDisqusInsert.Disqus_Insert{}, "disqus_comments")
	
//Create the map that will contain all our structs from Posts
	set = make(map[int]xpnk_createDisqusInsert.Disqus_Insert)
		
	for i := 0; i < len(disqusposts); i++ {
		set[i] = disqusposts[i]
	}
		
	fmt.Printf("\n==========\nset is now:%+v\n",set)
	
	
	//Insert the the Disqus comments!	
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