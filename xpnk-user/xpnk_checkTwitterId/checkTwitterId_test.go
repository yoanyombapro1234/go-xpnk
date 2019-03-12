package xpnk_checkTwitterId

import   (
"testing"
"fmt"
)

func TestCheckUserInvite(t *testing.T) {

	token := "1095770999872806913-l3OUMvb8Jt2N1KdBnIbQd5CoFurIiC"
	secret := "gWyHbf3Xskk6jJQO5q8WX9zzJeIUmAoOdRJUpQXPa0PjQ"
	
	//token := "131547767-O0v9F9vnAM1YTsyWL6500oDsXQRHuSoecObwqSM"
	//secret := "ihNLKHTUk9010DQRMOWxISx2WrxZFLYVnXhLVj6ac"
	
	v, err := CheckTwitterId(token, secret)
	
	if err != nil {
		t.Error("Expected a user ID, got ", err)
	}
	
	fmt.Printf("User info: %+s", v)
	
}