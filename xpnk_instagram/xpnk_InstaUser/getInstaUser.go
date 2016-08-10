package xpnk_instaUser

/**************************************************************************************
Takes an Instagram Username and requests the user ID from Instagram API
**************************************************************************************/

import (
  "fmt"
  "github.com/yanatan16/golang-instagram/instagram"
  "net/url"
)

func getInstaUserId(instaName string) string{

	instaToken := ""

	api := instagram.New("", instaToken)
	
	//if ok, err := api.VerifyCredentials(); !ok {
	//	panic(err)
	//	return
	//} 
	// this is occasionally returning 'not enough arguments to return' even when 
	// the credentials are indeed verified by Instagram
	
	fmt.Println("Successfully created instagram.Api with user credentials")

	params := url.Values{}
	params.Set("count", "1")
	params.Set("q", instaName)
	
	usersResponse, err := api.GetUserSearch(params)
	if err != nil {
		panic(err)
	}	
	
	fmt.Printf("Users Object:  %v\n", usersResponse)
	
	instaUsers := usersResponse.Users
	fmt.Printf("User object:  %v \n", instaUsers)
	
	instaUserId := usersResponse.Users[0].Id
	fmt.Printf("UserID:  %v \n", instaUserId)
	
	return instaUserId
}