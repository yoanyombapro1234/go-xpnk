package xpnk_stats

import 
(
	"testing"
	"fmt"
 )
 
func TestGetStats(t *testing.T) {

	groupID		:=	60
	groupName	:= 	"xapnik-testing"

	var v GroupStats
	
	v = GetStats(groupID, groupName)
	
	if 
		//v == nil
		v.GroupName != "" {
		//t.Errorf("Expected %c, got %v", testToken, resp)
		fmt.Printf("This is your SlackName:   %+s/n", v.Stats[0].SlackName)
	} else {
		t.Errorf("Expected an integer, got %v", v)
	}				

}