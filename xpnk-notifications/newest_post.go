package xpnk_notifications

import (
	"fmt"
	"time"
	 _ "github.com/go-sql-driver/mysql"
	 "xpnk-shared/db_connect"
	 "xpnk-group/xpnk_group_structs"
)

func CheckNewestPost (groupID int) (xpnk_group_structs.NewestPosts, error) {

	fmt.Printf("\n groupID: \n%+v", groupID)

	var newest_posts xpnk_group_structs.NewestPosts

	dbmap := db_connect.InitDb()
	defer dbmap.Db.Close()
	
	err := dbmap.SelectOne(&newest_posts, "SELECT * FROM group_newest_posts WHERE Group_ID=?", groupID)
	
	if err != nil {
		fmt.Printf("\n==========\n Check Newest Post - Problemz with getting newest posts for group: \n%+v\n",err)
		return newest_posts, err
	}
	
	fmt.Printf("\n Newest posts: \n%+v", newest_posts)
	return newest_posts, err
}	

func GetGroupTweets (groupID int) ([]string, error) {
	dbmap := db_connect.InitDb()
	defer dbmap.Db.Close()
	
	var group_members []int
	var twitter_ids []string

	_, err := dbmap.Select(&group_members, "SELECT `user_ID` FROM `USER_GROUPS` WHERE `Group_ID`=?", groupID)
	if err != nil {
		fmt.Printf("\n==========\n Get Group Tweets - Problemz with getting user ids for group: \n%+v\n",err)
		return twitter_ids, err
	}
	twitter_ids, err = GetTwitterIDs(group_members)
	return twitter_ids, err
}

func GetTwitterIDs (group_members []int) ([]string, error) {
	dbmap := db_connect.InitDb()
	defer dbmap.Db.Close()
	var err error
	var twitter_ids []string
	for i := 0; i < len(group_members); i++ {
		tweeter := group_members[i]
		var thisTweeter string
		fmt.Printf("this tweeter: %v", tweeter)

		err := dbmap.SelectOne(&thisTweeter,"SELECT `twitter_ID` FROM `USERS` WHERE `user_ID`=?", tweeter)
					
		twitter_ids = append(twitter_ids, thisTweeter)
		if err != nil {
			fmt.Printf("\n==========\n Get Group Tweets - Problemz with getting Twitter ids for tweeter: \n%+v\n",err)
			fmt.Printf("Tweeter: %v", tweeter)
			return twitter_ids, err
		}
	}
	return twitter_ids, err
}	

func GetTweetDates (twitter_ids []string)([]string, error){
	dbmap := db_connect.InitDb()
	defer dbmap.Db.Close()
	var err error
	var tweet_dates []string
	var tweet_date string
	for i := 0; i < len(twitter_ids); i++ {
		twitter_id := twitter_ids[i]
		err := dbmap.SelectOne(&tweet_date,"SELECT `tweet_date` FROM `tweets` WHERE `twitter_ID`=?", twitter_id)
		tweet_dates = append(tweet_dates, tweet_date)
		if err != nil {
			fmt.Printf("\n==========\n Get Group Tweets - Problemz with getting Twitter ids for tweeter: \n%+v\n",err)
			fmt.Printf("Tweeter: %v", twitter_id)
			return tweet_dates, err
		}
	}
	return tweet_dates, err
}

func CountNewTweets (base_tweet string, tweet_dates []string) (int, error) {
	var count int
	layout := "2006-01-02 15:04:05"
	base_tweet_time, err := time.Parse(layout, base_tweet)
	if err != nil {
		fmt.Printf("\n==========\n Error converting base_tweet_date to time: \n%+v\n",err)
	}
	count = 0
	for i := 0; i < len(tweet_dates); i++ {
		tweet_date := tweet_dates[i]
		t, err := time.Parse(layout, tweet_date)
		if err != nil {
			fmt.Printf("\n==========\n Error converting tweet_date to time: \n%+v\n",err)
		}
		if base_tweet_time.Before(t) {
			count++
		}
	}
	return count, err	
}

