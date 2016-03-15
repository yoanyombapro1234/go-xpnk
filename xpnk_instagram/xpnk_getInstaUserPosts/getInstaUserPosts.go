package xpnk_instaUserPosts

import (
  "fmt"
  "github.com/yanatan16/golang-instagram/instagram"
  "net/url"
)
	
func getInstaUserPosts(instaUserId string) {

	instaToken := "192772980.1fb234f.6121c6ef7adb4aaf86923777a8d2c1c2"

	api := instagram.New("", instaToken)
	
	//if ok, err := api.VerifyCredentials(); !ok {
	//	panic(err)
	//	return
	//} 
	// this is occasionally returning 'not enough arguments to return' even when 
	// the credentials are indeed verified by Instagram
	
	fmt.Println("Successfully created instagram.Api with user credentials")
	
	params := url.Values{}
	params.Set("count", "2")
	params.Set("max_timestamp", "")
	params.Set("min_timestamp","")
	
	instaPosts, err := api.GetUserRecentMedia(instaUserId, params)
	
	if err != nil {
		panic(err)
	}	
	
	fmt.Printf("Medias Object:  %+v\n", instaPosts)

	for i := 0; i < len(instaPosts.Medias); i++ {
	
		instaPostID := instaPosts.Medias[i].Link
		fmt.Printf("Media Link: %v\n", instaPostID)
	
	}		
}	