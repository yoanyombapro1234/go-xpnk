package xpnk_insertMultiUsers

import  (
	"testing"
	"xpnk-user/xpnk_createUserInsert"
)	

func TestInsertMultiUsers(t *testing.T) {

/*
	testInsert := []xpnk_createUserInsert.User_Insert{
		{SlackID:"XXXXXXXX", SlackName:"testslacker", SlackAvatar:"https://avatars.slack-edge.com/2016-06-03/48132819874_f6a67138d2a7a35ca9e9_original.jpg"}, {SlackID:"YYYYYYYYY", SlackName:"testslacker02", SlackAvatar:""},{SlackID:"ZZZZZZZZZ", SlackName:"testslacker03", SlackAvatar:"https://avatars.slack-edge.com/2015-06-10/6237358659_962602bad8ad1e816ed4_original.jpg"},		
	}
*/	
	testInsert := []xpnk_createUserInsert.User_Insert{
			{
			TwitterUser: "Test",
			InstaUser:	"Test",
			TwitterID:	"abcdef123456",
			InstaUserID: "abcdef123456",
			InstaAccessToken: "abcdef123456xxxxx.sajfidjafdij",
			DisqusUserName: "Test",
			DisqusUserID: "abcdef123456",
			DisqusAccessToken:	"abcdef123456xxxxx.sajfidjafdij",
			DisqusRefreshToken:	"abcdef123456xxxxx.sajfidjafdij",
			ProfileImage:	"https://avatars.slack-edge.com/2015-06-10/6237358659_962602bad8ad1e816ed4_original.jpg",
			},	
	}

	
	v := InsertMultiUsers(testInsert)
	
	if v != "inserted" {
		t.Errorf("Expected INSERTED, got  ", v)
	}
				
}