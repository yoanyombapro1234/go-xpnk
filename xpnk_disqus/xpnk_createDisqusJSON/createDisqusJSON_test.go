package xpnk_createDisqusJSON

import  "testing"

func TestCreateDisqusJSON(t *testing.T) {

	v := CreateDisqusJSON(1)
	
	if v != "File created!" {
		t.Error("Expected File created!, got ", v)
	}
				
}