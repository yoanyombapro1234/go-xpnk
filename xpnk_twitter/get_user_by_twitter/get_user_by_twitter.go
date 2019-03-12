package get_user_by_twitter

import (
	"fmt"
   	"xpnk-shared/db_connect"
)

func GetUserByTwitter(twitter_id string) (int, error) {
	dbmap := db_connect.InitDb()
	defer dbmap.Db.Close()
	
	var xpnkUser			int
	var err_msg				error
	twitterId				:= twitter_id
	
	err	:= dbmap.SelectOne(&xpnkUser, "SELECT `user_ID` FROM USERS WHERE twitter_ID=?", twitterId)
	if err != nil {
		fmt.Printf("\n==========\nget_user_by_twitter - Problemz with selecting user by twitterID: \n%+v\n",err)
		err_msg = err
		fmt.Printf("\n==========\nget_user_by_twitter - Problemz with selecting user by twitterID: \n%+v\n",err_msg)
		if err_msg.Error() == "sql: no rows in result set" {
			xpnkUser = 0
			fmt.Printf("\n==========\nNo user found: %+v\n",xpnkUser)
		}
		
	} else {
		fmt.Printf("\n==========\nfound user: \n%+v\n",xpnkUser)
	} 
	
	fmt.Printf("\n==========\nUser ID: %+v\n",xpnkUser)
	return xpnkUser, err_msg
}