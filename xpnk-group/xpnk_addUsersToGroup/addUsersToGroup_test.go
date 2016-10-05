package xpnk_addUsersToGroup

import  (
	"testing"
)	

func TestAddUsersToGroup(t *testing.T) {

	testUsers := []int{ 84, 85, 86,		
	}
	testGroup := 9
	
	v := AddUsersToGroup(testGroup,testUsers)
	
	if v != "inserted" {
		t.Errorf("Expected INSERTED, got  ", v)
	}
				
}