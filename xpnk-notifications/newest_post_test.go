package xpnk_notifications

import (
	"testing"
	"fmt"
	"xpnk-group/xpnk_group_structs"
)

func TestCheckNewestPost(t *testing.T) {

	groupID := 1
	var 	newest_posts	xpnk_group_structs.NewestPosts
		
	newest_posts, err := CheckNewestPost(groupID)
	
	if err != nil {
		t.Errorf("Test failed with error  ", err)
		fmt.Printf("\nError:  %+v \n", err)
	} else {
		fmt.Printf("\nnewest_posts: %+v\n", newest_posts)
	}
				
}

func TestGetGroupTweets (t *testing.T){
	groupID := 60
	var twitter_ids []string
	
	twitter_ids, err := GetGroupTweets(groupID)
	
	if err != nil {
		t.Errorf("Test failed with error  ", err)
		fmt.Printf("\nError:  %+v \n", err)
	} else {
		fmt.Printf("\ntwitter_ids: %+v\n", twitter_ids)
	}
}

func TestGetTweetDates (t *testing.T){
	twitter_ids := []string{"131547767", "9999999999999999999"}
	var tweet_dates []string
	tweet_dates, err := GetTweetDates(twitter_ids)
	if err.Error() != "sql: no rows in result set" {
		t.Errorf("Test failed with error  ", err)
		fmt.Printf("\nError:  %+v \n", err)
	} else {
		fmt.Printf("\ntweet_dates: %+v\n", tweet_dates)
	}
}

func TestCountNewTweets (t *testing.T){
	base_tweet := "2019-04-21 03:23:37"
	tweet_dates := []string{"2019-04-22 03:23:37","2019-04-23 03:23:37"}
	count, err := CountNewTweets(base_tweet, tweet_dates)
	if err != nil {
		t.Errorf("Test failed with error  ", err)
		fmt.Printf("\nError:  %+v \n", err)
	} else {
		fmt.Printf("\ncount: %+v\n", count)
	}
}