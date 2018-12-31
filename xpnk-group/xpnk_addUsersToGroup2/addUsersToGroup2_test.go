package xpnk_addUsersToGroup2

import  (
	"testing"
	"fmt"
)	

func TestAddUsersToGroup(t *testing.T) {

	var testUser Group_User
	testUser.User		= "1"
	testUser.Group		= 1001
	testUser.Owner		= true
	
	v, err := AddUsersToGroup(testUser)
	
	if v.Group <= 0 {
		t.Errorf("Expected an integer, got  ", v.Group)
		fmt.Printf("\nError:  %+v \n", err)
	}
				
}