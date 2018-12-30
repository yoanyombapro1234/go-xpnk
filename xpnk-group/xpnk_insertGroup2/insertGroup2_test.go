package xpnk_insertGroup2

import  (
	"testing"
	"xpnk-group/xpnk_createGroupInsert"
	"time"
	"fmt"
)	

func TestInsertGroup2(t *testing.T) {

	var insert xpnk_createGroupInsert.Group_Insert
	
	time					:= time.Now()
	timeString				:= time.String()
	testGroupName			:= "UnitTestGroup" + timeString
	insert.GroupName = testGroupName
	insert.SourceID = ""
	insert.Source = "xapnik-app"
	
	v, err := InsertGroup2(insert)
	
	if v <= 0 {
		t.Errorf("Expected an int, got  ", v)
		fmt.Printf("\nError:  %+v \n", err)
	}
				
}