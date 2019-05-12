package users

import  (
	"testing"
	"xpnk-user/xpnk_user_structs"
)	

func TestSaveSub(t *testing.T) {
	
	testInsert := xpnk_user_structs.UserSub {
		Id 			: 1,
		Endpoint	: "https://api.pushservice.com/somethingunique",
		Type		: 0,
		P256dh		: "BIPUL12DLfytvTajnryr2PRdAgXS3HGKiLqndGcJGabyhHheJYlNGCeXl1dn18gSJ1WAkAPIxr4gK0_dQds4yiI=",
		Auth		: "FPssNDTKnInHVndSTdbKFw==",
	}

	
	v := SaveSub(testInsert)
	
	if v < 1 {
		t.Errorf("Expected a user ID, got  ", v)
	}
				
}