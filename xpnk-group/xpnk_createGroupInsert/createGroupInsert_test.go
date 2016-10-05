package xpnk_createGroupInsert

import   (
"testing"
)

func TestCreateGroupInsert(t *testing.T) {

	var newgroup New_Group
	
	newgroup.GroupName = "Test Group"
	newgroup.SourceID  = "T12345"
	newgroup.Source    = "Slack"

	v := CreateGroupInsert(newgroup)

	if v.GroupName != "Test Group" || v.SourceID != "T12345" || v.Source != "Slack" {
		t.Errorf("Expected %c, got %v", v)
	}				
}