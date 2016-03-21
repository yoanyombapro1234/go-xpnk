package xpnk_storeInstagramID

import "testing"

func TestStoreInstaUserId(t *testing.T) {

	var v int64

	v = StoreInstaUserId("192772980","kirstenlambertsen")
	//should return 1

	if v != 1 {
		t.Error("Expected 1 , got ", v)
	}
}