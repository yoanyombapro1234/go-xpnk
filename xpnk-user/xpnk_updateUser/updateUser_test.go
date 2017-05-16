package xpnk_updateUser

import   (
"testing"
)

func TestUpdateUser(t *testing.T) {

	var userupdate User_Update
	
	userupdate.XpnkID 			= 207
	userupdate.TwitterUser 		= "Test"
	userupdate.InstaUser 		= "Test"
	userupdate.TwitterID 		= "abcdef123456"
	userupdate.InstaUserID 		= "abcdef123456"
	userupdate.InstaAccessToken = "abcdef123456xxxxx.sajfidjafdij"
	
	

	v := UpdateUser(userupdate)

	if v != 1 {
		t.Errorf("Expected 1, got %v", v)
	}				
}