package groups

import (
	"fmt"
	 "github.com/gin-gonic/gin"
	 "xpnk-group/xpnk_createGroup"
)

func GroupsNew (c *gin.Context) {
	var newGroup				xpnk_createGroup.NewGroup
	var err_msg					error
	c.Bind(&newGroup)
	fmt.Printf("newGroup to add:  %+v \n", newGroup)
	if newGroup.GroupName == "" {
		c.JSON(400, "A group name is required to create a new group.")
		return
	}
	
	newID, err_msg 			:=  xpnk_createGroup.CreateGroup(newGroup)
	if err_msg != nil {
		c.JSON(400, err_msg.Error())	
	} else {
		c.JSON(200, newID)
	}
}