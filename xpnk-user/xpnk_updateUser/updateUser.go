package xpnk_updateUser

/**************************************************************************************
Takes a slice of User_Update objects and inserts them into USERS table
**************************************************************************************/

import (
   	"xpnk-shared/db_connect"
	"fmt"
   	"log"
)

type User_Update struct {
	User_ID				int			   `db:"user_ID"			json:"user_ID"`
	Slack_userid		string		   `db:"slack_userid"		json:"slack_userid"`
	Slack_name			string		   `db:"slack_name"			json:"slack_name"`
	Twitter_user		string		   `db:"twitter_user"		json:"twitter_user"`
	Twitter_ID			string		   `db:"twitter_ID"			json:"twitter_ID"`
	Twitter_token		string		   `db:"twitter_authtoken"	json:"twitter_token"`
	Twitter_secret		string		   `db:"twitter_secret"		json:"twitter_secret"`
	Insta_user			string		   `db:"insta_user"			json:"insta_user"`
	Insta_userid		string		   `db:"insta_userid"		json:"insta_userid"`
	Insta_token			string		   `db:"insta_accesstoken"	json:"insta_token"`
	Disqus_username		string		   `db:"disqus_username"	json:"disqus_username"`
	Disqus_userid		string		   `db:"disqus_userid"		json:"disqus_userid"`
	Disqus_token		string		   `db:"disqus_accesstoken"	json:"disqus_token"`
	Profile_image		string		   `db:"profile_image"		json:"profile_image"`
}

func UpdateUser(userupdate User_Update) int64{

	fmt.Printf("\n==========\nuserupdate: \n%+v\n", userupdate)

	dbmap := db_connect.InitDb()
	defer dbmap.Db.Close()
	
//map the User_Update struct to the 'USERS' db table
	dbmap.AddTableWithName(User_Update{}, "USERS").SetKeys(true, "user_ID")
	
	count, err := dbmap.Update(&userupdate)
	if err != nil {
		fmt.Printf("\n==========\nUPDATE DIDN'T WORK: \n%+v\n", err)
		checkErr(err, "\nupdateUser couldn't update the db:")
		return 0
	} else {
		fmt.Printf("\n==========\nuserupdate returned: \n%+v\n", count)
		return 1
	}
} 

func checkErr(err error, msg string) {
  if err != nil {
  log.Fatalln(msg, err)
	}	
} 
