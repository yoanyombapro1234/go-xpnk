package xpnk_instaUser

import  "testing"

func TestGetInstaUserId(t *testing.T) {

	var v string

	v = getInstaUserId("wearewakanda")

	if v != "1399704848" {
		t.Error("Expected 1399704848, go ", v)
	}
}