package auth

import (
	"fmt"
	"net/http"
	 "github.com/gin-gonic/gin"
	 "xpnk_auth"
	 "xpnk_constants"
)

func XPNKReadHeader (c *gin.Context) int{
	var this_header http.Header
	this_header = c.Request.Header
	fmt.Printf("HEADER: %+v", this_header)
	token := this_header["Token"]
	fmt.Printf("TOKEN: %+v", token)
	if len(token) != 0  {
		fmt.Printf("TOKEN ONLY:  %+v", this_header["Token"][0])
		auth := xpnk_auth.ParseToken(token[0], xpnk_constants.SigningKey)
		if auth == 1 {
			return 1
		} else {
			fmt.Printf("INVALID TOKEN:  %+V", this_header["Token"])
			return 0
		}
	} else {
		fmt.Printf("NO TOKEN FOUND:  %+V", this_header["Token"])
		return 0
	}
}