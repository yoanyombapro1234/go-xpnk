package xpnk_get_groups

import (
	"testing"
	"fmt"
)

func TestGetGroups (t *testing.T) {

	user_id := "1"

	v, err := GetGroups(user_id)
	
	if err != nil {
		t.Error("Expected a a UserGroups object, got ", err)
	}
	
	fmt.Printf("User's Groups: %+s", v)
				
}


