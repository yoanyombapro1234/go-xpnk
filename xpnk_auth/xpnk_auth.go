package xpnk_auth

import (
    "fmt"
    "time"

    "github.com/dgrijalva/jwt-go"
)

const (
    mySigningKey = ""
)

func XPNKAuth() {
    createdToken, err := NewToken([]byte(mySigningKey))
    if err != nil {
        fmt.Println("Creating token failed")
    }
    ParseToken(createdToken, mySigningKey)
}

func NewToken(mySigningKey []byte) (string, error) {
    // Create the token
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
    	"foo": "bar",
    	"nbf": time.Now().Unix(),
    } )
    // Sign and get the complete encoded token as a string
    tokenString, err := token.SignedString(mySigningKey)
    fmt.Println(tokenString, err)
    return tokenString, err
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