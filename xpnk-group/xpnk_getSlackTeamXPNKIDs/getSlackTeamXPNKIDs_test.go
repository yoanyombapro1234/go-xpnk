package xpnk_getSlackTeamXPNKIDs

import   (
"testing"
"xpnk-user/xpnk_createUserInsert"
)

func TestGetSlackUserXPNKID(t *testing.T) {
	
	testBatch := []xpnk_createUserInsert.User_Insert{
		{SlackID:"", SlackName:"", SlackAvatar:""},	
	}

	test_count := len(testBatch)

	var v int 
	v = len(GetSlackUserXPNKID(testBatch))
	
	if v != test_count {
		t.Errorf("Expected %c, got %v", test_count, v)
	}		
}