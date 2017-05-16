package xpnk_auth

import (
    "fmt"
    "time"

    "github.com/dgrijalva/jwt-go"
)

const (
    MySigningKey = ""
)

func XPNKAuth() {
    createdToken, err := NewToken([]byte(MySigningKey), "", "")
    if err != nil {
        fmt.Println("Creating token failed")
    }
    ParseToken(createdToken, MySigningKey)
}

func NewToken(MySigningKey []byte, source string, identifier string) (string, error) {
    // Create the token
    if source == "" { source = "none" }
    if identifier == "" { identifier = "none" }
    
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
    	"iss": source,
    	"jti": identifier,
    	"nbf": time.Now().Unix(),
    } )
    // Sign and get the complete encoded token as a string
    tokenString, err := token.SignedString(MySigningKey)
    fmt.Println(tokenString, err)
    return tokenString, err
}

func GetNewGroupToken(source string, identifier string) string {
	thistoken, err 					:= NewToken([]byte(MySigningKey), source, identifier)
	var	NewGroupToken				string
	if err != nil {
		NewGroupToken 				= "There was an error creating the token."
		fmt.Println(err)
	}	else {
		NewGroupToken 				= thistoken
	}
	return NewGroupToken
}			

func ParseToken(myToken string, myKey string) int {
    token, err := jwt.Parse(myToken, func(token *jwt.Token) (interface{}, error) {
        return []byte(myKey), nil
    })

    if err == nil && token.Valid {
        fmt.Println("Your token is valid.  I like your style.")
        return 1
    } else {
        fmt.Println("This token is terrible!  I cannot accept this: %e", err)
        return 0
    }
}

func UnpackToken(myToken string, myKey string) interface{} {
    token, err := jwt.Parse(myToken, func(token *jwt.Token) (interface{}, error) {
        return []byte(myKey), nil
    })
	
	if err == nil && token.Valid {
		fmt.Println("Unpacked the token:  %+v\n", token)
		return "Success!"
	} else {
		fmt.Println("Couldn't unpack the token:  %+v\n", err)
		return err
	}	
}