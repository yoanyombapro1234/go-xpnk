package xpnk_twitter

//TODO insert TwitterUser and ProfileImageURL into users table instead of tweets table

//take a list of twitter_user names and get their most recent posts, save to db

import (
	"fmt"
	"strings"
   	_ "github.com/go-sql-driver/mysql"
   	anaconda "github.com/ChimeraCoder/anaconda"
   	"net/url"
   	"time"

)

//stores tweet and associated twitter_user name for insertion into db
type New_Tweet_Insert struct {
	TwitterUser		string				`db:"twitter_user"`
	TweetID			string				`db:"tweet_ID"`
	TweetDate		time.Time			`db:"tweet_date"`
	TweetOembed		string				`db:"tweet_oembed"`
	Twitter_ID		string				`db:"twitter_ID"`
	ProfileImageURL	string				`db:"profile_image"`
}

func add_new_user(u string) {

	dbmap := initDb()
	defer dbmap.Db.Close()
	
	Tweeter := u
	
    //fmt.Printf("\n==========\nTweeter:%v\n",Tweeter)
    
		//using PERSONAL keys and secrets for development - CHANGE to KURATUR when go live
		/*
		*anaconda settings for the Twitter keys and secrets
		*/
		anaconda.SetConsumerKey("v209uvVUUfWtFHKXTrcFQ")
		anaconda.SetConsumerSecret("AQkRQfl7k7rwt1Zbh6RbCeKHTyQwNaD7mH78Yz1g")
		api := anaconda.NewTwitterApi("131547767-O0v9F9vnAM1YTsyWL6500oDsXQRHuSoecObwqSM", "ihNLKHTUk9010DQRMOWxISx2WrxZFLYVnXhLVj6ac")

		//build the Twitter query	
		username := Tweeter
		fetchcount := ""
		const trimuser = "true"
	
		v := url.Values{}
		v.Set("from", username)
		v.Set("count", fetchcount)
		//v.Set("trim_user", trimuser)
		fmt.Printf("\n==========\nVALUES: \n%+v\n",v)

		usertimeline, err := api.GetSearch("",v)
		if err != nil {
			fmt.Printf("GetSearch returned error: %s", err.Error())
		}
	
		fmt.Printf("\n==========\nTIMELINE: \n%+v\n",usertimeline)
	
		//TODO extract only the needed fields and format into Twitter oEmbed html format?
		//1) store the single tweet data in a Tweet_Insert struct
		//2) append the Tweet_Insert into a slice of Tweet_Inserts
		
		var New_Tweet_Inserts []New_Tweet_Insert
		
		for i := 0; i < len(usertimeline.Statuses); i++ {
			//first put stuff from usertimeline into Tweet_Insert
			//then get the oembed and put into Tweet_Insert
			//Get the tweet ID for each tweet
			
			var this_tweet_insert New_Tweet_Insert
			
			this_tweet_insert.TwitterUser = usertimeline.Statuses[i].User.ScreenName
			this_tweet_insert.TweetID = usertimeline.Statuses[i].IdStr
			this_tweet_insert.Twitter_ID = usertimeline.Statuses[i].User.IdStr

			
			//convert Twitter's created_at time format to time.Time format
			this_created_at := usertimeline.Statuses[i].CreatedAt
			this_created_date, _ := time.Parse(time.RubyDate,this_created_at)
			
			this_tweet_insert.TweetDate = this_created_date
						
			//remove '_normal' from end of profile image filename in order to get full-size avatar
			r := strings.NewReplacer("_normal", "")
			p := usertimeline.Statuses[i].User.ProfileImageURL
			avatar := r.Replace(p)
			
			this_tweet_insert.ProfileImageURL = avatar
			
			getoembed := usertimeline.Statuses[i].IdStr
			fmt.Printf("\n==========\nIDSTRING: \n%+v\n",getoembed)			
		
			//Get the oembed code for each tweet has to be queried separately :(
			vals := url.Values{}
			vals.Set("id", getoembed)
			vals.Set("omit_script", "true")
			embed, err := api.GetOEmbed(vals)
			if err != nil {
				fmt.Printf("GetUserTimeline returned error: %s", err.Error())
			}
	
			fmt.Printf("\n==========\nEMBED: \n%+v\n",embed.Html)
			
			this_tweet_insert.TweetOembed = embed.Html

			fmt.Printf("\n==========\nTHIS_INSERT: \n%+v\n",this_tweet_insert)
			
			New_Tweet_Inserts = append(New_Tweet_Inserts, this_tweet_insert)

    	}
    	    	
    	fmt.Printf("\n==========\nTHIS_BATCH: \n%+v\n",New_Tweet_Inserts)
    	
    	new_doinsert(New_Tweet_Inserts)
    }	
    
    
    //end Twitter query routine
    //fuckin' a, girl, it works!!!!


//end of func add_new_user


/***************************
*DATABASE INSERT FUNCTION
***************************/

type new_twitterposts []struct{}

func new_doinsert(new_twitterposts []New_Tweet_Insert) string{

	//Initialize a map variable to hold all our Tweet_Insert structs (tweets)
	var set map[int]New_Tweet_Insert

	dbmap := initDb()
	defer dbmap.Db.Close()
	
//map the []Tweet_Insert struct to the 'tweets' db table
	dbmap.AddTableWithName(New_Tweet_Insert{}, "TWEETS")
	
//Create the map that will contain all our structs from Posts
	set = make(map[int]New_Tweet_Insert)
		
	for i := 0; i < len(new_twitterposts); i++ {
		set[i] = new_twitterposts[i]
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