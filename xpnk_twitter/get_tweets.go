package xpnk_twitter

//TODO - take profile image out of tweets table and associate it with user in user table instead -- only done the *first time* the user's tweets are fetched

import (
	"fmt"
	"strings"
    "database/sql"
   	_ "github.com/go-sql-driver/mysql"
   	"github.com/gopkg.in/gorp.v1"
   	anaconda "github.com/ChimeraCoder/anaconda"
   	"net/url"
   	"time"
   	"github.com/chnlr/emoji"

)

//stores only the user_ID's of each member of a group
type GroupMemberID struct {
    UserID		string				`db:"user_ID"`
}

//stores the twitter_user names for each GroupMemberID
type GroupTweeter struct {
    TwitterUser	string				`db:"twitter_user"`
    ProfileImage string				`db:"profile_image" json:"profile_image"`
}

//stores a twitter_user and its last_tweet id
type Last_Tweet struct {
    TwitterUser	string					
    LastTweet	string					
}

//stores tweet and associated twitter_user name for insertion into db
type Tweet_Insert struct {
	TwitterUser		string				`db:"twitter_user"`
	TweetID			string				`db:"tweet_ID"`
	TweetDate		time.Time			`db:"tweet_date"`
	TweetOembed		string				`db:"tweet_oembed"`
	Twitter_ID		string				`db:"twitter_ID"`
	ProfileImageURL	string				`db:"profile_image"`
	TweetMedia		string				`db:"tweet_media"`
}

func Get_tweets() {

	dbmap := initDb()
	defer dbmap.Db.Close()
	
	group_ids := Get_Groups()

	var this_group int

	//this is where we iterate over each group to get all the tweets per group
	//everything else happens inside this loop
	for i := 0; i < len(group_ids); i++ {
		this_group = group_ids[i].GroupID
		

		//get all user_ID's for the current Group_ID from USER_GROUPS
		var group_members []GroupMemberID

		_, err := dbmap.Select(&group_members, "SELECT `user_ID` FROM `USER_GROUPS` WHERE `Group_ID`=?", this_group)

		checkErr(err, "Select for users of group failed")
	
		//now let's get the twitter_user name from USERS for each user_ID in group_members
		var Group_tweeters []GroupTweeter
	
		for i := 0; i < len(group_members); i++ {
		
			tweeter := group_members[i].UserID
		
			_, err := dbmap.Select(&Group_tweeters, "SELECT `twitter_user` FROM `USERS` WHERE `user_ID`=?", tweeter)
		
			checkErr(err, "Select for twitter_user names failed")
	
		}
	
		//get last_tweet for each TwitterUser in Group_tweeters & put in a slice of Last_Tweet
	
		var This_Twitter_Fetch []Last_Tweet

		for i := 0; i < len(Group_tweeters); i++ {
	
			var this_twitter_user string

			this_twitter_user = Group_tweeters[i].TwitterUser
		
			fmt.Printf("\n==========\nNow Doing:%+v\n",this_twitter_user)
		
			this_last_tweet, err := dbmap.SelectStr("SELECT MAX(`tweet_ID`) FROM `TWEETS` WHERE `twitter_user`=?", this_twitter_user)
			
			//handle case where this_twitter_user isn't found, or no tweet_IDs
			if (err != nil) {
			//run add_new_users.go for this user
			
			fmt.Printf("\n==========\nNo tweets were found for :%v. Running add_new_user.",this_twitter_user)	
			
			add_new_user(this_twitter_user)
			} else {	
			
			fmt.Printf("\n==========\nFound:%+v\n",this_last_tweet)

			checkErr(err, "Select for last_tweet failed" + this_twitter_user)
			
			//append this_twitter_user & this_last_tweet to slice of Last_Tweet
			var this_tweeter Last_Tweet
			this_tweeter.TwitterUser = this_twitter_user
			this_tweeter.LastTweet = this_last_tweet
			This_Twitter_Fetch = append(This_Twitter_Fetch, this_tweeter)
			}
		
		}
	
			fmt.Printf("\n==========\nTweeters and Last Tweets:%+v\n",This_Twitter_Fetch)

	
		//iterate through This_Twitter_Fetch and query Twitter for tweets since last_tweet
	
		for i := 0; i < len(This_Twitter_Fetch); i++ {
			/*
			*anaconda settings for the Twitter keys and secrets
			*/
			anaconda.SetConsumerKey("")
			anaconda.SetConsumerSecret("")
			api := anaconda.NewTwitterApi("")

			//build the Twitter query	
			username := This_Twitter_Fetch[i].TwitterUser
			fetchcount := ""
			lasttweet := This_Twitter_Fetch[i].LastTweet
			const trimuser = "true"	
			v := url.Values{}
			v.Set("from", username)
			v.Set("count", fetchcount)
			v.Set("since_id", lasttweet)
			
			fmt.Printf("\n==========\nVALUES: \n%+v\n",v)

			usertimeline, err := api.GetSearch("",v)
			if err != nil {
				fmt.Printf("GetSearch returned error: %s", err.Error())
			}
	
			fmt.Printf("\n==========\nTIMELINE: \n%+v\n",usertimeline)
	
			//store each tweet & twitter_name in a slice of struct to then insert into db
			//1) store the single tweet data in a Tweet_Insert struct
			//2) append the Tweet_Insert into a slice of Tweet_Inserts
		
			var Tweet_Inserts []Tweet_Insert
		
			for i := 0; i < len(usertimeline.Statuses); i++ {
				//first put stuff from usertimeline into Tweet_Insert
				//then get the oembed and put into Tweet_Insert
				//Get the tweet ID for each tweet
			
				var this_tweet_insert Tweet_Insert
			
				this_tweet_insert.TwitterUser = usertimeline.Statuses[i].User.ScreenName
				this_tweet_insert.TweetID = usertimeline.Statuses[i].IdStr
				this_tweet_insert.Twitter_ID = usertimeline.Statuses[i].User.IdStr
				
				
				//If this is a retweet or retweet of a retweet, we need to grab the url 
				//for any associated media like images or videos				
				rtentities := len(usertimeline.Statuses[i].RetweetedStatus.Entities.Media) 
				
				fmt.Printf("\n==========\nRT MEDIA OBJECTS: \n%+v\n",rtentities)
				
				qtentities := len(usertimeline.Statuses[i].RetweetedStatus.QuotedStatus.Entities.Media)
				
				fmt.Printf("\n==========\nQT MEDIA OBJECTS: \n%+v\n",qtentities)
				
				//crazy if else is crazy -- use switch?
				if rtentities == 0 && qtentities == 0 {
					this_tweet_insert.TweetMedia = ""
					
					fmt.Printf("\n==========\nMEDIA URLs ARE EMPTY")
					
				} else if rtentities == 0 && qtentities != 0 {
				this_tweet_insert.TweetMedia = usertimeline.Statuses[i].RetweetedStatus.QuotedStatus.Entities.Media[0].Media_url
				
				fmt.Printf("\n==========\nMEDIA URL: \n%+v\n",this_tweet_insert.TweetMedia)
				
				} else if rtentities != 0 {
				this_tweet_insert.TweetMedia = usertimeline.Statuses[i].RetweetedStatus.Entities.Media[0].Media_url
				
				fmt.Printf("\n==========\nMEDIA URL: \n%+v\n",this_tweet_insert.TweetMedia)
				
				} else {
				this_tweet_insert.TweetMedia = ""
					
				fmt.Printf("\n==========\nNO CASES RETURNED TRUE FOR MEDIA URL'S")
				}
				
				//convert Twitter's created_at time format to time.Time format
				this_created_at := usertimeline.Statuses[i].CreatedAt
				this_created_date, _ := time.Parse(time.RubyDate,this_created_at)
				this_tweet_insert.TweetDate = this_created_date
						
				//remove '_normal' from end of profile image filename
				r := strings.NewReplacer("_normal", "")
				p := usertimeline.Statuses[i].User.ProfileImageURL
				avatar := r.Replace(p)
				this_tweet_insert.ProfileImageURL = avatar
				getoembed := usertimeline.Statuses[i].IdStr
				
				fmt.Printf("\n==========\nIDSTRING: \n%+v\n",getoembed)	
						
				//Get the oembed code for each tweet, has to be queried separately
				vals := url.Values{}
				vals.Set("id", getoembed)
				vals.Set("omit_script", "true")
				embed, err := api.GetOEmbed(vals)
				if err != nil {
					fmt.Printf("GetUserTimeline returned error: %s", err.Error())
				}
	
				fmt.Printf("\n==========\nEMBED: \n%+v\n",embed.Html)
				
				this_tweet_insert.TweetOembed = emoji.UnicodeToTwemoji(embed.Html, 16, false)

				fmt.Printf("\n==========\nTHIS_INSERT: \n%+v\n",this_tweet_insert)
			
				Tweet_Inserts = append(Tweet_Inserts, this_tweet_insert)

			}
				
			fmt.Printf("\n==========\nTHIS_BATCH: \n%+v\n",Tweet_Inserts)
		
			doinsert(Tweet_Inserts)
		}	
   } //end group_ids loop
    
    //end Twitter query routine
    //fuckin' a, girl, it works!!!!

}
//end of func main


/***************************
*DATABASE INSERT FUNCTION
***************************/

type twitterposts []struct{}

func doinsert(twitterposts []Tweet_Insert) string{

	//delete tweets older than 24 hours, func from delete-old-tweets.go
	defer Dodelete() 

	//Initialize a map variable to hold all our Tweet_Insert structs (tweets)
	var set map[int]Tweet_Insert

	dbmap := initDb()
	defer dbmap.Db.Close()
	
//map the []Tweet_Insert struct to the 'tweets' db table
	dbmap.AddTableWithName(Tweet_Insert{}, "TWEETS")
	
//Create the map that will contain all our structs from Posts
	set = make(map[int]Tweet_Insert)
		
	for i := 0; i < len(twitterposts); i++ {
		set[i] = twitterposts[i]
	}
		
	fmt.Printf("\n==========\nset is now:%+v\n",set)
	
	
	//Insert the the tweets!	
	for _, v := range set {	
			
		//db insert function 
		err := dbmap.Insert(&v)
		if err != nil {fmt.Printf("There was an error ", err)
			
		}
	}	
		return "inserted"
		
}
//end db insert routine	
//Maps are not inherently safe for concurrency - will have to use sync.RWMutex	


/***************************
* db connection config
***************************/	
func initDb() *gorp.DbMap {
db, err := sql.Open("mysql",
	"")
checkErr(err, "sql.Open failed")

dbmap := &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{}}

return dbmap
}