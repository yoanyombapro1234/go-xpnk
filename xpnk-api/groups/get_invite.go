package groups

import (
	"strconv"
	"github.com/gin-gonic/gin"
	"xpnk-group/xpnk_createInvite"
)

func GroupsInvite (c *gin.Context) {
	id, err 				:= strconv.Atoi(c.Params.ByName("id"))
	if err != nil {
		c.JSON( 400, err.Error())
		return
	}
	source		 			:= c.Params.ByName("source")	
	response, err			:= xpnk_createInvite.CreateInvite(id, source, "")
	if err != nil {
		c.JSON( 400, err.Error())
		return
	}
	c.JSON(200, response)
}