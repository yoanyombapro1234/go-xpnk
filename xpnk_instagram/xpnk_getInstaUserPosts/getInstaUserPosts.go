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
	
func GetInstaUserPosts(instaUserId string) *instagram.PaginatedMediasResponse{

	instaToken := "192772980.1fb234f.6121c6ef7adb4aaf86923777a8d2c1c2"

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
	
	params := url.Values{}
	params.Set("count", "2")
	params.Set("max_timestamp", "")
	params.Set("min_timestamp",unixtoday)
	
	instaPosts, err := api.GetUserRecentMedia(instaUserId, params)
	
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