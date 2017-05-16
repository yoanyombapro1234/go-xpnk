package xpnk_createGroupFromSlack

import   (
"testing"
"strings"
)

func TestInvite(t *testing.T) {

	var newinvite Slacker
	
	newinvite.Token 		= ""
	newinvite.Slacker  		= ""
	newinvite.SlackGroup    = ""
	newinvite.XpnkGroup		= strings.Replace(newinvite.SlackGroup, " ", "-", -1)
	newinvite.XpnkToken		= ""

	v := Invite(newinvite)

	if v != "Success!" {
		t.Errorf("Expected Success!, got %v", v)
	}	
}