package get_user_by_insta

import (
	"testing"
	"fmt"
)

func TestGetUserByInsta (t *testing.T) {

	insta_id := "192772980" /*"0000000000"*/

	v, err := GetUserByInsta(insta_id)
	
	if err != nil {
		t.Error("Expected a user id, got ", err.Error())
	}
	
	fmt.Printf("User's Xapnik id: %s", v)
				
}


