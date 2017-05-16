package xpnk_auth

import   (
"testing"
)

func TestUnpackToken(t *testing.T) {
	var v							interface {}
	token							:= ""
	v								= UnpackToken(token, mySigningKey)
	if v != "Success!" {
		t.Errorf("Expected Success!, got %v", v)
	}	
}