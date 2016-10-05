package xpnk_createGroupFromSlack

import   (
"testing"
)

func TestCreateGroup(t *testing.T) {

	var token string
	token = ""
	
	v := CreateGroup(token)

	if v != "Check the database!" {
		t.Errorf("Expected %c, got %v", v)
	}				
}