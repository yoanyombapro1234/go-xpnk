package xpnk_getDisqusUserPosts

import  "testing"

func TestGetDisqusUserPosts(t *testing.T) {

	var disqusUser DisqusUser 

	disqusUser.DisqusName = "TestUSer"
	disqusUser.Disqus_accesstoken = "1j1j1j1j1j1jj1j1j1j1j1j1j1j1j1j1j1j1j1j1j1j1"
	
	GetDisqusUserPosts(disqusUser)
				
}