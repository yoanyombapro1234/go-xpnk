package get_user_by_twitter

import (
	"testing"
	"fmt"
)

func TestGetUserByTwitter (t *testing.T) {

	twitter_id := "0000000"
	//twitter_id := "131547767"

	v, err := GetUserByTwitter(twitter_id)
	
	if err != nil {
		t.Error("Expected a username!, got ", err)
	}
	
	fmt.Printf("User's Xapnik id: %s", v)
				
}


