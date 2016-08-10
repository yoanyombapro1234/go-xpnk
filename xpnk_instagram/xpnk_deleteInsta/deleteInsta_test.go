package xpnk_deleteInsta

import  (
	"testing"
)	

func TestDeleteInsta(t *testing.T) {

	var thisUser string
	
	thisUser = "192772980"

	v := DeleteInsta(thisUser)
	
	if v != "deleted" {
		t.Errorf("Expected DELETED, got  ", v)
	}
				
}