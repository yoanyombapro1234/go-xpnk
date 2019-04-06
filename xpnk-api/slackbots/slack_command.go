package slackbots

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"xpnk_slack"
	"xpnk_constants"
)

func SlackCommandHandler (c *gin.Context) {
    fmt.Printf("GIN CONTEXT:  %+v", c)
	var command_body xpnk_slack.SlackCommand
	var token string
	c.Bind(&command_body) 
	fmt.Printf("COMMAND BODY: %+v", command_body)
	
	token = command_body.Token

	if token != "" && token == xpnk_constants.SlackCommandTkn { 
		response := xpnk_slack.SlackGroupStatus(command_body)
		c.JSON(200, response)
	}
}