package xpnk_createMultiUserInsert

import   (
"testing"
"xpnk-user/xpnk_createUserInsert"
)

func TestCreateDisqusInsert(t *testing.T) {
	
	testBatch := []xpnk_createUserInsert.User_Insert{
		{SlackID:"XXXXXXXX", SlackName:"testname", SlackAvatar:"https://avatars.slack-edge.com/2016-06-03/xxxxxxxxxxxxxxxxxxxxx_original.jpg"}, {SlackID:"XXXXXXX02", SlackName:"testname02", SlackAvatar:""},{SlackID:"XXXXXXXXX03", SlackName:"testuser03", SlackAvatar:"https://avatars.slack-edge.com/2015-06-10/xnnnnnnnnnnnnnnnnnnnnnnnnn_original.jpg"},		
	}

	test_count := len(testBatch)

	var v int 
	v = len(CreateMultiUserInsert(testBatch))
	
	if v != test_count {
		t.Errorf("Expected %c, got %v", test_count, v)
	}		
}