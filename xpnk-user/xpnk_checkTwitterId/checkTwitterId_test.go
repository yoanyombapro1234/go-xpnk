package xpnk_checkTwitterId

import   (
"testing"
"fmt"
"xpnk-user/xpnk_createUserObject"
)

func TestCheckUserInvite(t *testing.T) {

	var twitter_user xpnk_createUserObject.User_Object
	twitter_user.TwitterID = "131547767"

	//token := "131547767-O0v9F9vnAM1YTsyWL6500oDsXQRHuSoecObwqSM"
	//secret := "ihNLKHTUk9010DQRMOWxISx2WrxZFLYVnXhLVj6ac"
	
	v, err := CheckTwitterId(twitter_user)
	
	if err != nil {
		t.Error("Expected a user ID, got ", err)
	}
	
	fmt.Printf("\nUser info: %+s\n", v)
	
}