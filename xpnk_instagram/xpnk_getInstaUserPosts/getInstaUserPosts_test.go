package xpnk_getInstaUserPosts

import  "testing"

func TestGetInstaUserPosts(t *testing.T) {

	var instaUser InstaUser 

	instaUser.InstaID = ""
	instaUser.Insta_accesstoken = ""
	
	GetInstaUserPosts(instaUser)
				
}
