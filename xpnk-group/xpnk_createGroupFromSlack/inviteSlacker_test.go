package xpnk_createGroupFromSlack

import   (
"testing"
)

func TestInvite(t *testing.T) {

	var newinvite Slacker
	//put your group's stuff here
	//newinvite.Token 			= ""
	//newinvite.Slacker  		= ""
	//newinvite.SlackGroup    	= ""
	//newinvite.XpnkGroup		= ""

	v := Invite(newinvite)

	if v != "Success!" {
		t.Errorf("Expected Success!, got %v", v)
	}	
}