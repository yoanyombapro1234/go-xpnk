package xpnk_stats

import (
"fmt"
//"strings"
_ "github.com/go-sql-driver/mysql"
"xpnk-shared/db_connect"
)

type GroupStats struct{
	GroupName	string
	GroupURL	string
	Stats		[]MemberStats
}

type MemberStats struct{
	XpnkID		int
	SlackName	string
	Tweets		int
	IGs			int
	Comments	int
}

type GroupMember struct{
	XpnkID		int			`db:"user_ID"`
	SlackName	string		`db:"slack_name"`
	TwitterID	string		`db:"twitter_ID"`
	InstaID		string		`db:"insta_userid"`
	DisqID		string		`db:"disqus_userid"`
}

var groupID int
var groupName string

/*
1) get all the group member ids
2) get all the group member handles and associate with their member id
3) for each member, get count of all tweets
4) for each member, get count of all IG's
5) for each member, get count of all Disqus comments
6) put everything in a GroupStats struct and return
*/

func GetStats(groupID int, groupName string) GroupStats {

	var memberIDs []int
	memberIDs = GetGroupMembers(groupID)
	fmt.Printf("These are your group members: %+v", memberIDs)

	var members []GroupMember
	for i := 0; i < len(memberIDs); i++ { 
		var thisUser GroupMember
		thisID				:= memberIDs[i]
		thisUser			= GetSocialIDs(thisID)
		members = append(members, thisUser)
		fmt.Printf("This is a GroupMember:  %+v", thisUser)
	}
	
	var groupStats GroupStats
	groupStats.GroupName = groupName
	groupStats.GroupURL = "https://xapnik.com/"+groupStats.GroupName
	
	fmt.Printf("This is the group URL: %v", groupStats.GroupURL)
	
	for i := 0; i < len(members); i++ { 
		var memberStats MemberStats
		tweetCount := GetTweetCount(members[i].TwitterID)
		instaCount := GetInstaCount(members[i].InstaID)
		disqCount := GetDisqCount(members[i].DisqID)
		memberStats.SlackName = members[i].SlackName
		memberStats.Tweets = tweetCount
		memberStats.IGs = instaCount
		memberStats.Comments = disqCount
		groupStats.Stats = append(groupStats.Stats, memberStats)
		
		fmt.Printf("These are this user's stats: Slackname:%s, Tweets:%t, IGs:%v, Disqus:%w", memberStats.SlackName, memberStats.Tweets, memberStats.IGs, memberStats.Comments)
	} 
	
	return groupStats
 
}

func GetGroupMembers (groupID int) []int{
	dbmap 					:= db_connect.InitDb()
	defer dbmap.Db.Close()
	
	var groupMembers []int
	
	_, err := dbmap.Select(&groupMembers, "SELECT user_ID FROM user_groups WHERE Group_ID=?", groupID)
	
	if err == nil {
		return groupMembers
	} else {
		fmt.Printf("\n==========\n GetGroupMembers - Problemz with getting user_IDs for group in xpnk_stats line 45: \n%+v\n",err)
		return groupMembers
	}
}

func GetSocialIDs (userID int) GroupMember {
	dbmap 					:= db_connect.InitDb()
	defer dbmap.Db.Close()
	
	var socialHandles GroupMember
	
	err := dbmap.SelectOne(&socialHandles, "SELECT `user_ID`,`slack_name`,`twitter_ID`,`insta_userid`,`disqus_userid` FROM USERS WHERE user_ID=?", userID)
	
	if err == nil {
		return socialHandles
	} else {
		fmt.Printf("\n==========\n GetSocialIDs - Problemz with getting social id's for group in xpnk_stats line 87: \n%+v\n",err)
		return socialHandles
	}
}

func GetTweetCount(twitterID string) int{
	dbmap 					:= db_connect.InitDb()
	defer dbmap.Db.Close()
	
	var tweetCount int
	
	err := dbmap.SelectOne(&tweetCount, "SELECT COUNT(*) FROM TWEETS WHERE twitter_ID=?", twitterID)
	
	if err == nil {
		return tweetCount
	} else {
		fmt.Printf("\n==========\n GetGroupMembers - Problemz with getting tweetCount for user in xpnk_stats line 103: \n%+v\n",err)
		return tweetCount
	}
}

func GetInstaCount(instaID string) int{
	dbmap 					:= db_connect.InitDb()
	defer dbmap.Db.Close()
	
	var instaCount int
	
	err := dbmap.SelectOne(&instaCount, "SELECT COUNT(*) FROM instagram_posts WHERE insta_userid=?", instaID)
	
	if err == nil {
		return instaCount
	} else {
		fmt.Printf("\n==========\n GetInstaCount - Problemz with getting instaCount for user in xpnk_stats line 126: \n%+v\n",err)
		return instaCount
	}
}

func GetDisqCount(disqID string) int{
	dbmap 					:= db_connect.InitDb()
	defer dbmap.Db.Close()
	
	var disqCount int
	
	err := dbmap.SelectOne(&disqCount, "SELECT COUNT(*) FROM disqus_comments WHERE disqus_userid=?", disqID)
	
	if err == nil {
		return disqCount
	} else {
		fmt.Printf("\n==========\n GetDisqCount - Problemz with getting disqCount for user in xpnk_stats line 144: \n%+v\n",err)
		return disqCount
	}
}
