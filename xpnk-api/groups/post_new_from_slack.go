package groups

import (
	"fmt"
	 "github.com/gin-gonic/gin"
	 "xpnk-group/xpnk_createGroupFromSlack"
)

func SlackCreateNewGroup (c *gin.Context) {
	fmt.Printf("GIN CONTEXT:  %+v", c)
	var team_tokens xpnk_createGroupFromSlack.SlackTeamTokens
	var token string
	var bot_token string
	c.Bind(&team_tokens) 
	fmt.Printf("TEAM TOKENS: %+v", team_tokens)
	fmt.Printf("TOKEN ONLY:  %+v", team_tokens.TeamToken)
	fmt.Printf("TEST MODE: %+v", team_tokens.TestToken)
	
	token = team_tokens.TeamToken
	bot_token = team_tokens.BotToken

	if token != "" && bot_token != "" { 
	c.JSON(200, "Well hello, Slack friend! I'll set up your Xapnik group and send the team invitations when it's ready.")
		xpnk_createGroupFromSlack.CreateGroup(team_tokens)
	} else {
		c.JSON(422, gin.H{"error":"I'm sorry, I can't see either or both of the token parameters."})
	}
}