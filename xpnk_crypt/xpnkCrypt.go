package xpnk_crypt

import 
(
	 "fmt"
	 "crypto/aes"
	 "crypto/cipher"
 )

const (
    key = "" 
)

type InitObject struct {
	Ciphertext		[]byte
	Block			cipher.Block
	IV				[]byte
}

func Initialize() InitObject {
	var init InitObject
	
	block, err := aes.NewCipher([]byte(key))
		if err != nil {
		  panic(err)
		}

	init.Ciphertext = []byte("") 
    init.Block = block
    init.IV = init.Ciphertext[:aes.BlockSize]
    
    return init
}
 
func Encrypt(this_token string)  []byte {
	
	init := Initialize()
	
    str := []byte(this_token)    

    // encrypt

    encrypter := cipher.NewCFBEncrypter(init.Block, init.IV)

    encrypted := make([]byte, len(str))
    encrypter.XORKeyStream(encrypted, str)

    fmt.Printf("%s encrypted to %v\n", str, encrypted)
    
    return encrypted
}    

func Decrypt(this_encrypted []uint8) string{
    // decrypt
    
    init := Initialize()
    
    decrypter := cipher.NewCFBDecrypter(init.Block, init.IV) // simple!

    decrypted := make([]byte, len(this_encrypted))
    decrypter.XORKeyStream(decrypted, this_encrypted)
    
    n := len(decrypted)
    decrypt_string := string(decrypted[:n])

    fmt.Printf("%v decrypt to %s\n", this_encrypted, decrypted)
    
    return decrypt_string
}