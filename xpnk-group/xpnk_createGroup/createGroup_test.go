package xpnk_createGroup

import   (
"testing"
"time"
"fmt"
)

func TestCreateGroup(t *testing.T) {

	time					:= time.Now()
	timeString				:= time.String()
	testGroupName			:= "UnitTestGroup" + timeString
	var request NewGroup
	request.Owner			= "1"
	request.Source			= "xapnik-app"
	request.GroupName		= testGroupName
	
	v, err := CreateGroup(request)

	if v >= 0 {
		t.Errorf("Expected %c, got %v", v)
		fmt.Printf("\nError:  %+v \n", err)
	}				
}