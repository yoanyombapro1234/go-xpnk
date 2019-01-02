package xpnk_createInvite

import   (
"testing"
"fmt"
)

func TestCreateInvite(t *testing.T) {

	var group_id int
	group_id = 60
	
	v	 := CreateInvite(group_id, "xapnik-app", "")

	if v == "" {
		t.Errorf("Expected a token, got %v", v)
	}
	
	fmt.Println("Invite URL: %v", v)				
}