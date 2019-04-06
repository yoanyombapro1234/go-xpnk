package users

import (
	"fmt"
	 _ "github.com/go-sql-driver/mysql"
   	 "database/sql"
	 "xpnk-shared/db_connect"
)

func DelUserGroups (userID int) (sql.Result, error) {
	dbmap 					:= db_connect.InitDb()
	defer dbmap.Db.Close()
		
	res, err := dbmap.Exec("delete from USER_GROUPS where User_ID=?", userID)
	
	if err != nil {
		fmt.Printf("\n===========\n delUserGroups error: %+v", err)
	} else {
		fmt.Printf("\n===========\n delUserGroups response: %+v", res)
	}
	
	return res, err
}
