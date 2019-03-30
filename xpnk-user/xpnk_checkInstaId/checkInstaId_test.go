package xpnk_checkInstaId

import   (
"testing"
"fmt"
"xpnk-user/xpnk_createUserObject"
)

func TestCheckUserInvite(t *testing.T) {

	var insta_user xpnk_createUserObject.User_Object
	insta_user.InstaUserID = /*"192772980"*/ "0000000000000"
	
	v, err := CheckInstaId(insta_user)
	
	if err != nil {
		t.Error("Expected a user ID, got ", err)
	}
	
	fmt.Printf("\nUser info: %+s\n", v)
	
}