package xpnk_getSlackTeam

import (
  "testing"
  "xpnk-user/xpnk_createUserInsert"
)  

func TestGetSlackGroup(t *testing.T) {

	var slackToken string 

	slackToken = ""
	
	var v []xpnk_createUserInsert.User_Insert
	
	v = GetSlackTeam(slackToken)
	
    if v[0].SlackID == "" {
		t.Errorf("Expected a slice of Slack Users, got  ", v)
	}
				
}