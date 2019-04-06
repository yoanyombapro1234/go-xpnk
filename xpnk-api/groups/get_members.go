package groups

import (
	"fmt"
	"github.com/gin-gonic/gin"
	  _ "github.com/go-sql-driver/mysql"
	 "xpnk-shared/db_connect"
)

func GroupsByID (c *gin.Context) {
	var groupid 	string
	groupid 		= c.Param("id")
	
	if groupid != "" {
		groupmems 	:= getGroup(groupid)	
		fmt.Printf("\nGROUPMEMS RETURNED BY GetGROUP: %+v\n", groupmems)	 
		c.JSON(201, groupmems) 
	} else {
		c.JSON(422, gin.H{"error": "No group_id was sent."})
	}		 
}

func getGroup (groupID string) []int{
	var groupUsers			[]int
	dbmap := db_connect.InitDb()
	defer dbmap.Db.Close()
	
	err, _ := dbmap.Select(&groupUsers, "SELECT user_ID FROM USER_GROUPS WHERE Group_ID=?", groupID)
	
	if err == nil {
		return groupUsers
	} else {
		fmt.Printf("\n==========\n getGroup - Problemz with getting users for group api line 830: \n%+v\n",err)
		return groupUsers
	}
}