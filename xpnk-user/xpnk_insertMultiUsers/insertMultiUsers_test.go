package xpnk_insertMultiUsers

import  (
	"testing"
	"xpnk-user/xpnk_createUserInsert"
)	

func TestInsertMultiUsers(t *testing.T) {

	testInsert := []xpnk_createUserInsert.User_Insert{
		{SlackID:"XXXXXXXX", SlackName:"testslacker", SlackAvatar:"https://avatars.slack-edge.com/2016-06-03/dajf;aidfjadkjfidjfkjad_original.jpg"}, {SlackID:"YYYYYYYYY", SlackName:"testslacker02", SlackAvatar:""},{SlackID:"ZZZZZZZZZ", SlackName:"testslacker03", SlackAvatar:"https://avatars.slack-edge.com/2015-06-10/xxxxxxxxxxxxxxxxxxxx_original.jpg"},		
	}
	
	v := InsertMultiUsers(testInsert)
	
	if v != "inserted" {
		t.Errorf("Expected INSERTED, got  ", v)
	}
				
}