package xpnk_checkUserInvite

import   (
"testing"
)

func TestCheckUserInvite(t *testing.T) {
	var v							GroupObj
	token							:= ""
	group							:= ""
	
	v								= CheckUserInvite(token, group)
	if v.GroupName != group {
		t.Errorf("Expected %g, got %v", group, v)
	}	
}