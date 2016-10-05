package xpnk_createDisqusInsert

import   (
"testing"
"xpnk_disqus/xpnk_getDisqusUserPosts"
)

func TestCreateDisqusInsert(t *testing.T) {

	var testUser xpnk_getDisqusUserPosts.DisqusUser
	testUser.DisqusName = "TestUser"
	testUser.Disqus_accesstoken = ""
	
	resp := xpnk_getDisqusUserPosts.GetDisqusUserPosts(testUser)
	
	resp_count := len(resp)
	
	var v int

	v = len(CreateDisqusInsert(resp))

	if v != resp_count {
		t.Errorf("Expected %c, got %v", resp_count, v)
	}				
}