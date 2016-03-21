package xpnk_storeInstagramID

//stores the Instagram ID of a user in the user's db record

import (
	"fmt"
   	_ "github.com/go-sql-driver/mysql"
   	"database/sql"
   	"github.com/gopkg.in/gorp.v1"
)

//stores Instagram user ID for insertion into db
type InstagramUserID struct {
	Instagram_username	string	`db:"insta_user"	json:"insta_user"`
	Instagram_ID		string	`db:"insta_userid" 	json:"insta_userid"`
}

func StoreInstaUserId(instaUserId string, instaUserName string) int64 {

	dbmap := initDb()
	defer dbmap.Db.Close()
	
	this_Insta_UserID := instaUserId
	this_Insta_UserName := instaUserName
	var this_Insta_UserInsert InstagramUserID
	
	this_Insta_UserInsert.Instagram_username = this_Insta_UserName
	this_Insta_UserInsert.Instagram_ID = this_Insta_UserID
	
	fmt.Printf("\n==========\nUSERNAME: \n%v\n",this_Insta_UserInsert.Instagram_username)
	
	stmt, err := dbmap.Exec("UPDATE USERS SET insta_userid = ? WHERE insta_user = ?", this_Insta_UserID, this_Insta_UserName)
	count, err := stmt.RowsAffected()
	if err != nil {
		fmt.Printf("stmt.RowsAffected() returned error: %s", err.Error())
	}
	if count != 1 {
		fmt.Printf("expected 1 affected row, got %d", count)
	}

	return count
}

/***************************
* db connection config
***************************/	
func checkErr(err error) {
	if err != nil {
		panic("psql err: " + err.Error())
	}
}

func initDb() *gorp.DbMap {
db, err := sql.Open("mysql",
	"root:root@tcp(localhost:8889)/xapnik")
checkErr(err)

dbmap := &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{}}

return dbmap
}