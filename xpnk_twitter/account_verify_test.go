package xpnk_twitter

import (
	"testing"
	"fmt"
)

func TestAccountVerify (t *testing.T) {

	token := "131547767-O0v9F9vnAM1YTsyWL6500oDsXQRHuSoecObwqSM"
	secret := "ihNLKHTUk9010DQRMOWxISx2WrxZFLYVnXhLVj6ac"

	v := AccountVerify(token, secret)
	
	if v == " " {
		t.Error("Expected a username!, got ", v)
	}
	
	fmt.Printf("Twitter userid: %s", v)
				
}


