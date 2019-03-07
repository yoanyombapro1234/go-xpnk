package xpnk_twitter

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
	} else {
		fmt.Printf("\n==========\nfound user: \n%+v\n",xpnkUser)
	}
	return xpnkUser, err_msg
}