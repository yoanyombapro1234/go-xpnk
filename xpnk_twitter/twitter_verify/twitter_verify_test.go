package twitter_verify

import (
	"testing"
	"fmt"
)

func TestAccountVerify (t *testing.T) {

	//token := "7413522-AoRGbJk9QxqTwqiJZORwZK41WhKi5MTbNiNbkxpHVg"
	//secret := "p4rlKiuyieL2fer2QbOcPDk4veGoXqeYCChsvq3A9YIFR"

	token := "131547767-O0v9F9vnAM1YTsyWL6500oDsXQRHuSoecObwqSM"
	secret := "ihNLKHTUk9010DQRMOWxISx2WrxZFLYVnXhLVj6ac"

	v, s, err := AccountVerify(token, secret)
	
	if err != nil {
		t.Error("Expected a username!, got ", err.Error())
	}
	
	fmt.Printf("Twitter userid: %s", v)
	fmt.Printf("Twitter screen name: %s", s)
				
}


