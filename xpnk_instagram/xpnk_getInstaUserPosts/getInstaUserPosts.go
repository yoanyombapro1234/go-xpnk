package xpnk_getInstaUserPosts

/**************************************************************************************
Takes an Instagram user ID and retrieves today's posts by that user from Instagram API
**************************************************************************************/

import (
  "fmt"
  "github.com/yanatan16/golang-instagram/instagram"
  "net/url"
  "time"
  "strconv"
)
	
type InstaUser struct {
    InstaID				string				`db:"insta_userid"`
    Insta_accesstoken	string				`db:"insta_accesstoken"`
    Insta_maxID			string				
}	
	
func GetInstaUserPosts(insta_User InstaUser) *instagram.PaginatedMediasResponse{

	instaToken := insta_User.Insta_accesstoken

	api := instagram.New("", instaToken)
	
	//if ok, err := api.VerifyCredentials(); !ok {
	//	panic(err)
	//	return
	//} 
	// this is occasionally returning 'not enough arguments to return' even when 
	// the credentials are indeed verified by Instagram
	
	fmt.Println("Successfully created instagram.Api with user credentials")
	
	//TODO calculate unix timestamp for 24 hrs ago and insert into min_timestamp
	now := time.Now()
    secs := uint64(now.Unix() - (60 * 60 * 24))
    unixtoday := strconv.FormatUint(secs, 10)
    fmt.Println("Unix timestamp %v", unixtoday)
    fmt.Println("insta_User: %v", insta_User)
	
	params := url.Values{}
	params.Set("count", "10")
	params.Set("scope", "public_content")
	//params.Set("min_timestamp",unixtoday)
	//params.Set("min_id", insta_User.Insta_maxID)
	//these params are commented out in the hope that IG will fix their api some day
	
	fmt.Println("Params: %v", params)
	
	fmt.Println("Getting posts for insta_User: %v", insta_User.InstaID)
	
	instaPosts, err := api.GetUserRecentMedia(insta_User.InstaID, params)
	
	if err != nil {
		panic(err)
	}	
	
	fmt.Printf("Medias Object:  %+v\n", instaPosts)

	for i := 0; i < len(instaPosts.Medias); i++ {
	
		instaPostID := instaPosts.Medias[i].Link
		fmt.Printf("Media Link: %v\n", instaPostID)
	
	}		
	return instaPosts
}	