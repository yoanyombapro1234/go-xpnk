package xpnk_getDisqusUserPosts

/**************************************************************************************
Takes a Disqus user ID and retrieves today's posts by that user from Disqus API
**************************************************************************************/

import (
  "fmt"
  "net/url"
  "xpnk_disqus/golang-disqus/disqus"
  "xpnk_disqus/xpnk_getDisqusThreadDetails"
)

//only one of these is necessary for an api request for recent posts	
type DisqusUser struct {
    DisqusName				string				`db:"disqus_username"`
    Disqus_accesstoken		string				`db:"disqus_accesstoken"`
    Disqus_maxID			string
}	
	
func GetDisqusUserPosts(disqus_User DisqusUser) []disqus.Content{

	disqusClient := "Cen6kb83THtogsxE5I5cXh2VTgZyNKhH9th5G2kXsiA1UGYkG5NseX9zh4RO9ERx"
	disqusToken := disqus_User.Disqus_accesstoken
	disqusName := disqus_User.DisqusName

	api := disqus.New(disqusClient)
	
	fmt.Println("Successfully created disqus.Api with app credentials")
	
    fmt.Println("disqus_User: %v", disqus_User)
	
	params := url.Values{}
	
	if disqusName != "" {
		params.Set("user:username", disqusName)
	} else if disqusToken != "" {
		params.Set("access_token", disqusToken)
	} else {
		panic("Disqus username or access_token must be provided.")
	}
	
	params.Set("limit", "7")
	
	//params.Set("min_timestamp",unixtoday)
	//this param is commented out in the hope that IG will fix their api some day
	
	fmt.Println("Params: %v", params)
	
	fmt.Println("Getting posts for disqus_User: %v", disqus_User.DisqusName)
	
	disqusPosts, err := api.GetUserRecentComments(params)
	
	if err != nil {
		fmt.Printf("At line 51 of getDisqusUserPosts:  %+v",err)
		fmt.Printf ("At line 51 of getDisqusUserPosts:  ")

	}	
	
	fmt.Printf("Response Object:  %+v\n", disqusPosts.Contents)
	fmt.Printf("\n==================First comment id: %+v\n", disqusPosts.Contents[0].Id)

	//get the link to the post and append the url params to the specific comment and insert the entire link into the disqusPosts.Permalink item

	for i := 0; i < len(disqusPosts.Contents); i++ {
	
		threadID := disqusPosts.Contents[i].Thread
		fmt.Printf("/n=====Thread ID for adding to disqusPosts object: %v\n", threadID)
		threadDetails := xpnk_getDisqusThreadDetails.GetDisqusThreadDetails(threadID)	
		
		permalink := threadDetails.Link + "#comment-" + disqusPosts.Contents[i].Id
		fmt.Printf("/n=====Permalink: %v\n", permalink)
		disqusPosts.Contents[i].Permalink = permalink
		
		title := threadDetails.Title
		fmt.Printf("/n=====Title: %v\n", title)
		disqusPosts.Contents[i].Title = title
	}	
			
	disqusComments := disqusPosts.Contents
		
	return disqusComments
}	