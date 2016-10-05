package xpnk_deleteDisqus

/**************************************************************************************

Takes a disqus_user and deletes all its rows from disqus_comments table 
This is necessary to avoid duplicate posts being stored b/c Disqus API 'since' param
currently doesn't work

**************************************************************************************/

import (
    "database/sql"
   	_ "github.com/go-sql-driver/mysql"
   	"github.com/gopkg.in/gorp.v1"
   	"log"
)

func DeleteDisqus(disqusUser string) string{

	username := disqusUser

	dbmap := initDb()
	defer dbmap.Db.Close()
	
	_, err := dbmap.Exec("delete from disqus_comments where disqus_user = ?", username)
	checkErr(err,"Exec failed")
	
	count, err := dbmap.SelectInt("select count(*) from disqus_comments where disqus_user = ?", username)
	checkErr(err, "select count(*) failed")
	
	if count == 0 {	return "deleted" } else { return "delete failed" }
		
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