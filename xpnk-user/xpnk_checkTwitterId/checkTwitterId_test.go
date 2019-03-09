package xpnk_checkTwitterId

import   (
"testing"
"fmt"
)

func TestCheckUserInvite(t *testing.T) {

	token := "131547767-uKMLWtiEQfQ65CyeRvMd4JlNb3JDH87iiQDcnhax"
	secret := "mR9Je9jw64JyFDj5y0aTV4yC6zIN3Tkpr5GYTlm2WOXQu"
	
	//token := "131547767-O0v9F9vnAM1YTsyWL6500oDsXQRHuSoecObwqSM"
	//secret := "ihNLKHTUk9010DQRMOWxISx2WrxZFLYVnXhLVj6ac"
	
	v, err := CheckTwitterId(token, secret)
	
	if err != nil {
		t.Error("Expected a user ID, got ", err)
	}
	
	fmt.Printf("User info: %+s", v)
	
}