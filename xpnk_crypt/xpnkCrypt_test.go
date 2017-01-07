package xpnk_crypt

import 
(
	"testing"
	"fmt"
 )
 
func TestDecrypt(t *testing.T) {

	var testToken string
	testToken = "OhNo.NotMe.INeverLostControl.You'reFaceToFaceWithTheManWhoSoldTheWorld."
	
	encrypted := Encrypt(testToken)
	
	resp := Decrypt(encrypted)
	
	fmt.Printf("The test token is:  %v", resp)
		
	if resp != testToken {
		t.Errorf("Expected %c, got %v", testToken, resp)
	}				

}