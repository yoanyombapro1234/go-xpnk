package xpnk_deleteDisqus

import  (
	"testing"
)	

func TestDeleteDisqus(t *testing.T) {

	var thisUser string
	
	thisUser = "TestUser"

	v := DeleteDisqus(thisUser)
	
	if v != "deleted" {
		t.Errorf("Expected DELETED, got  ", v)
	}
				
}