package xpnk_getSlackGroup

import  "testing"

func TestGetSlackGroup(t *testing.T) {

	var slackToken string 

	slackToken = ""
	
	v := GetSlackGroup(slackToken)
	
    if v.SourceID != "" {
		t.Errorf("Expected true, got  ", v)
	}
				
}