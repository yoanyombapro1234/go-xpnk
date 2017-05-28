package xpnk_insertGroup

import  (
	"testing"
	"xpnk-group/xpnk_createGroupInsert"
)	

func TestInsertGroup(t *testing.T) {

	var insert xpnk_createGroupInsert.Group_Insert
	
	insert.GroupName = "dangerladies"
	insert.SourceID = "T12345"
	insert.Source = "Slack"
	
	v := InsertGroup(insert)
	
	if v != "inserted" {
		t.Errorf("Expected INSERTED, got  ", v)
	}
				
}