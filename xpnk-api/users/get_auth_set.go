package users

import (
	"fmt"
	 _ "github.com/go-sql-driver/mysql"
	 "github.com/gin-gonic/gin"
	 "xpnk_auth"
	 "xpnk_constants"
)

func XPNKAuthSet (c *gin.Context)  {
	usertoken, err := xpnk_auth.NewToken([]byte(xpnk_constants.SigningKey), "", "")
	if usertoken != "" {
		response := usertoken
		c.JSON(201, response)
	} 	else {
		fmt.Printf("ERROR: %v+", err)
		c.JSON(422,gin.H{"error":"No access token created."})
	}
}