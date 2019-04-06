package groups

import (
	"fmt"
	"strings"
	"github.com/gin-gonic/gin"
	 "xpnk-shared/db_connect"
)

func GroupID (c *gin.Context) {
	var groupname 	string
	groupname 		= c.Param("name")
	
	if groupname != "" {
		groupid 	:= getGroupID(groupname)	
		fmt.Printf("\nGROUPID RETURNED BY GetGroupID: %+v\n", groupid)	 
		c.JSON(201, groupid) 
	} else {
		c.JSON(422, gin.H{"error": "No groupname was sent."})
	}		 
}

func getGroupID (groupName string) int{
	var group_name			= strings.Replace(groupName, "-", " ", -1)
	var groupID				int
	dbmap 					:= db_connect.InitDb()
	defer dbmap.Db.Close()
	
	err := dbmap.SelectOne(&groupID, "SELECT Group_ID FROM GROUPS WHERE group_name=?", group_name)
	
	if err == nil {
		return groupID
	} else {
		fmt.Printf("\n==========\n getGroupID - Problemz with getting Group_ID for group api line 35: \n%+v\n",err)
		return groupID
	}
}