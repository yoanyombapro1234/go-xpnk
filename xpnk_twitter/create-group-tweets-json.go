package xpnk_twitter

/**************************************************************************************
Takes a group ID and writes all Tweets  for that group to a json file using a file-naming convention.

(1) get all the user id's for a single group
(2) get all the twitter user id's for each of those user ids
(3) get all the tweets associated with each of those twitter_id's
(4) put all the user's tweets into an object keyed by user's Xpnk_ID
(5) json-encode all the tweets and write to a file using a naming convention
**************************************************************************************/

import (
	"fmt"
   	_ "github.com/go-sql-driver/mysql"
    "strings"
    "bytes"
    "os"
    "encoding/json"
)

//stores only the group_name of the group_ID
type Groupname struct {
    GroupName	string		`db:"group_name"`
}

//stores the Twitter user name and XpnkID for each GroupMemberID
type GroupMember struct {
	XpnkID			string
    TweeterID		string	`db:"twitter_user"`
    ProfileImg		string	`db:"profile_image"`
}

//stores the tweet object for a tweet
type XpnkTweet struct {
	TwitterUser		string	`db:"twitter_user" 	json:"twitter_user"`
	TweetID			string	`db:"tweet_ID" 		json:"tweet_ID"`
	TweetDate		string	`db:"tweet_date" 	json:"tweet_date"`
	TweetOembed		string	`db:"tweet_oembed" 	json:"tweet_oembed"`
	Twitter_ID		string	`db:"twitter_ID" 	json:"twitter_ID"`
	ProfileImageURL	string	`db:"profile_image" json:"profile_image"`
	TweetMedia		string	`db:"tweet_media" 	json:"tweet_media"`
}

type UserTweets struct {
	XpnkID			string
	TwitterPosts	[]XpnkTweet
}

func Create_group_tweets_json() {
	
	dbmap := initDb()
	defer dbmap.Db.Close()
	
	//get ids for all groups	
	group_ids := Get_Groups()

	var this_group int

	//this is where we iterate over each group to get all the tweets per group
	//everything else happens inside this loop
	for i := 0; i < len(group_ids); i++ {
		this_group = group_ids[i].GroupID
		
		var group_name Groupname
	
		/******
		* converting the group name into a string to use in the filename
		******/
		//get group name to use for the filename	
		err := dbmap.SelectOne(&group_name, "SELECT `group_name` FROM `groups` WHERE `Group_ID`=?", this_group)
	
		if err != nil {fmt.Printf("There was an error ", err)}

		fmt.Printf("\n==========\nGROUP NAME:%+v\n",group_name)
	
		//extract just the group name string from group_name	
		this_name := group_name.GroupName
	
		//convert the group name into a hyphenated string for use in json filename
		this_name = strings.Replace(this_name, " ", "-", -1)	
	
		//convert all characters to lowercase
		this_name = strings.ToLower(this_name)
	
		fmt.Printf("\n==========\nGROUP NAME IS NOW:%+v\n",this_name)

		/******
		* get all the xpnk user_ID's associated with the Group_ID from USER_GROUPS
		******/
		var group_members []GroupMemberID

		_, err = dbmap.Select(&group_members, "SELECT `user_ID` FROM `USER_GROUPS` WHERE `Group_ID`=?", this_group)

		fmt.Printf("\n==========\nMember ID's:%+v\n",group_members)

		checkErr(err, "user_ID Select failed")


		/******
		* get each twitter_user name & profile image from USERS for each user_ID
		******/		
		var group_tweeters []GroupMember

		for i := 0; i < len(group_members); i++ {
	
			tweeter := group_members[i].UserID
			
			var thisTweeter GroupMember
	
			err := dbmap.SelectOne(&thisTweeter,"SELECT `twitter_user`,`profile_image` FROM `USERS` WHERE `user_ID`=?", tweeter)
			
			fmt.Printf("\n==========\nTHIS TWEETER ID:%+v\n==========\n",thisTweeter)
			
			thisTweeter.XpnkID = tweeter
		
			group_tweeters = append(group_tweeters, thisTweeter)
	
			checkErr(err, "twitter_user Select failed")
		}
		
		fmt.Printf("\n==========\nTWEETERS:%+v\n==========\n",group_tweeters)
	
		/******
		* write the twitter_user names to a file using a naming convention
		******/
		this_users, err := os.Create("/home/xapnik/node-v0.12.5/XAPNIK/data/"+this_name+"_users.json")
		
		//convert group_tweeters struct to json
		users_str, err := json.Marshal(group_tweeters)
		if err != nil {
			fmt.Println("Error encoding JSON")
			return
		}
	
			this_users.WriteString(string(users_str))
	
		/*******
		* get all the groups' tweets from the db
		*******/    
		var group_tweets []UserTweets
		
		var this_user  UserTweets

		var user_tweets []XpnkTweet

		for i := 0; i < len(group_tweeters); i++ {
	
			tweeter := group_tweeters[i].TweeterID
			
			this_user.XpnkID = group_tweeters[i].XpnkID
	
			_, err := dbmap.Select(&user_tweets, "SELECT * FROM `TWEETS` WHERE `twitter_user`=?", tweeter)
			
			this_user.TwitterPosts = user_tweets
		
			fmt.Printf("\n==========\nTHIS USERS POSTS:%+v\n==========\n",this_user)
		
			group_tweets = append(group_tweets, this_user)
		
			user_tweets = []XpnkTweet{} 
			//this has to be emptied or it carries over to the next tweeter
	
			checkErr(err, "user_tweets Select failed")
	
		}
		fmt.Printf("\n==========\nGROUP TWEETS:%+v\n==========\n",group_tweets)

	
		/******
		* write the contents of group_tweets to a .json file
		******/
		//create the file according to our naming convention
		this_file, err := os.Create("/home/xapnik/node-v0.12.5/XAPNIK/data/"+this_name+"_tweets.json")
			fmt.Printf("\n==========\nCREATED:%+v\n==========\n",this_file)

		//convert group_tweets to json
		str, err := JSONMarshal(group_tweets, true)
		if err != nil {
			fmt.Println("Error encoding JSON")
			return
		}
	
			this_file.WriteString(string(str))
	}//end group_ids loop		
			
}//end main  



/******
* keep html tags in tact in our json file
******/
func JSONMarshal(v interface{}, safeEncoding bool) ([]byte, error) {
    b, err := json.Marshal(v)

    if safeEncoding {
        b = bytes.Replace(b, []byte("\\u003c"), []byte("<"), -1)
        b = bytes.Replace(b, []byte("\\u003e"), []byte(">"), -1)
        b = bytes.Replace(b, []byte("\\u0026"), []byte("&"), -1)
        //b = bytes.Replace(b, []byte("\\\""), []byte("\""), -1)
    }
    return b, err
} 