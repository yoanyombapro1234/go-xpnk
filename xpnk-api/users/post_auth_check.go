package users

import (
	"fmt"
	"net/http"
	 _ "github.com/go-sql-driver/mysql"
	 "github.com/gin-gonic/gin"
	 "xpnk_auth"
	 "xpnk_constants"
)

func XPNKAuthCheck (c *gin.Context) {
	fmt.Printf("HEADER: %+v  END\n", c.Request.Header)
	var this_header http.Header
	this_header = c.Request.Header
	token := this_header["Token"]
	if len(token) != 0  {
		fmt.Printf("TOKEN ONLY:  %+v", this_header["Token"][0])
		auth := xpnk_auth.ParseToken(token[0], xpnk_constants.SigningKey)
		if auth == 1 {
			c.JSON(200, gin.H{"success":"You're clear for take off."})
		} else {
			c.JSON(422, gin.H{"error": "Token can't be authenticated."})
		}
	} else {
		c.JSON(422, gin.H{"error": "No access token was sent."})
	}
}