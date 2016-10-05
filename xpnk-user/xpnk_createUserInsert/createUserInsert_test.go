package xpnk_createUserInsert

import   (
"testing"
)

func TestCreateUserInsert(t *testing.T) {

	var newuser User_Insert
	
	newuser.SlackID = "JJJJJJJJJ01" 
	newuser.SlackName = "testuser01"
	newuser.ProfileImage = "https://avatars.slack-edge.com/2016-06-03/48132819874_jjjjjjjjjjjjjjjjjjjjjjj_original.jpg"
	newuser.SlackAvatar = "https://avatars.slack-edge.com/2016-06-03/jjjjjjjjjjjjjjjjjjjjjjj_original.jpg"

	v := CreateUserInsert(newuser)

	if v.SlackID != "KKKKKKKKK01" || v.SlackName != "testuser02" || v.ProfileImage != "https://avatars.slack-edge.com/2016-06-03/kkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkkk.jpg" {
		t.Errorf("Expected %c, got %v", v)
	}				
}