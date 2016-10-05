package xpnk_insertDisqus

import  (
	"testing"
	"xpnk_disqus/xpnk_createDisqusInsert"
)	

func TestInsertDisqus(t *testing.T) {

	var insert []xpnk_createDisqusInsert.Disqus_Insert
	var thisinsert xpnk_createDisqusInsert.Disqus_Insert
	
	thisinsert.DisqusUser = "TestUser"
	thisinsert.DisqusName = "Test User"
	thisinsert.DisqusUserID = "11111111"
	thisinsert.DisqusPID = "1111111111"
	thisinsert.DisqusPermalink = "http://example.com/2016/08/post/#comment-1111111111"
	thisinsert.DisqusTitle = "Test Post"
	thisinsert.DisqusEmbed = "<p>She has achieved the impossible. I hope her next endeavor is youth serum!</p>"
	thisinsert.DisqusDate = "2016-08-14T01:47:30"
	thisinsert.DisqusAvatar = "https://disqus.com/api/users/avatars/TestUser.jpg"
	thisinsert.DisqusMedia = "//a.disquscdn.com/uploads/mediaembed/images/4093/2864/original.gif"
	thisinsert.DisqusFavicon = "https://disqus.com/api/forums/favicons/blogger.jpg"
	thisinsert.DisqusForum = "blogger"
	
	insert = append(insert, thisinsert)

	v := InsertDisqus(insert)
	
	if v != "inserted" {
		t.Errorf("Expected INSERTED, got  ", v)
	}
				
}