package xpnk_checkUserInvite

import   (
"testing"
)

func TestCheckUserInvite(t *testing.T) {
	var v							GroupObj
	token							:= "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJTbGFjayIsImp0aSI6IlUzUjQ3MVhUMyIsIm5iZiI6MTQ5NTk5MTI3N30.0cwIqVYhzB_Z2gnIsm-k62GV5g_iIMVlyK66WplYg5g"
	group							:= "xapnik-testing"
	
	v								= CheckUserInvite(token, group)
	if v.GroupName != group {
		t.Errorf("Expected %g, got %v", group, v)
	}	
}