package xpnk_createInstaInsert

import   (
"testing"
"xpnk_instagram/xpnk_getInstaUserPosts"
)

func TestCreateInstaInsert(t *testing.T) {

	resp := xpnk_getInstaUserPosts.GetInstaUserPosts("208560325")
	
	resp_count := len(resp.Medias)
	
	var v int

	v = len(CreateInstaInsert(resp))

	if v != resp_count {
		t.Errorf("Expected %c, got %v", resp_count, v)
	}
				
}