package xpnk_createInstaJSON

import  "testing"

func TestCreateInstaJSON(t *testing.T) {

	v := CreateInstaJSON(5)
	
	if v != "File created!" {
		t.Error("Expected File created!, got ", v)
	}
				
}