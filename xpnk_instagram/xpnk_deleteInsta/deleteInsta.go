package xpnk_deleteInsta

/**************************************************************************************

Takes an insta_userid and deletes all its rows from instagram_posts table 
This is necessary to avoid duplicate posts being stored b/c IG API min_id param
currently doesn't work

**************************************************************************************/

import (
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
    "github.com/gopkg.in/gorp.v1"
    "log"
)

func DeleteInsta(instaID string) string{

	id := instaID

	dbmap := initDb()
	defer dbmap.Db.Close()
	
	_, err := dbmap.Exec("delete from instagram_posts where insta_userid = ?", id)
	checkErr(err,"Exec failed")
	
	count, err := dbmap.SelectInt("select count(*) from instagram_posts where insta_userid = ?", id)
	checkErr(err, "select count(*) failed")
	
	if count == 0 {	return "deleted" } else { return "delete failed" }
		
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
