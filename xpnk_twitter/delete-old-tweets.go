package xpnk_twitter

//delete all tweets that are older than 24 hours

import (
   	_ "github.com/go-sql-driver/mysql"
   	"time"

)

func Dodelete() string{

	now := time.Now()

    //fmt.Println("now:", now)

    yesterday := now.AddDate(0, 0, -1)

    //fmt.Println("Yesterday:", yesterday) 

// delete rows manually via Exec
	dbmap := initDb()
	defer dbmap.Db.Close()

    _, err := dbmap.Exec("DELETE FROM TWEETS WHERE tweet_date<?", yesterday)
    checkErr(err, "Exec failed")
    
    return "deleted"
}
//end Dodelete

//db connection details	are in full_monty and called from there

//checkErr is defined in full_monty and called from there