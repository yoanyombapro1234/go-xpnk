package db_connect

import (
    "database/sql"
   	_ "github.com/go-sql-driver/mysql"
   	"github.com/gopkg.in/gorp.v1"
   	"log"
)

/***************************
* db connection config
***************************/	
func InitDb() *gorp.DbMap {
db, err := sql.Open("mysql",
	"dbconnect credentials here")
checkErr(err, "sql.Open failed")

dbmap := &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{}}

return dbmap
}

func checkErr(err error, msg string) {
  if err != nil {
  log.Fatalln(msg, err)
	}	
} 