package users

import  (
	"testing"
	"xpnk-user/xpnk_user_structs"
)	

func TestSaveSub(t *testing.T) {
	
	testInsert := xpnk_user_structs.UserSub {
		Id 			: 1,
		FirebaseKey	: "FPssNDTKnInHVndSTdbKFw==",
	}

	
	v := SaveFireBaseKey(testInsert)
	
	if v < 1 {
		t.Errorf("Expected a user ID, got  ", v)
	}
				
}