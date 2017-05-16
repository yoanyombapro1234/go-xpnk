package xpnk_groupTokens

import   (
"testing"
)

func TestSaveGroupTokens(t *testing.T) {
	
	var group_count					GroupCount
	var v							string

	ids 							:= []string {}
		
	for i := 0; i <len(ids); i++ {
	
		group_count.XpnkGroup		= 
		group_count.Source			= "Slack"
		group_count.Identifier		= ids[i]
	
		v = SaveGroupToken(group_count)
	}	

	if v != "Success!" {
		t.Errorf("Expected Success!, got %v", v)
	}	
}