package get_user_object

import (
	"testing"
	"fmt"
)

func TestGetUserObject (t *testing.T) {

	user_id := "1"

	v, err := GetUserObject(user_id)
	
	if err != nil {
		t.Error("Expected a user object, got ", err)
	}
	
	fmt.Printf("User object: %s", v)
				
}


