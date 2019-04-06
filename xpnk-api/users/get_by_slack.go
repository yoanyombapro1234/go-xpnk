package users

import (
	"fmt"
	 "github.com/gin-gonic/gin"
	 "xpnk-shared/db_connect"
)

func GetXPNK_ID (c *gin.Context) {
	var slackuserid string
	slackuserid = c.Param("id")
	
	if slackuserid != "" {
		xpnkuserid := getXPNKUser(slackuserid)	
		fmt.Printf("\nXPNKUSERID RETURNED BY GetXPNK_ID: %+v\n", xpnkuserid)	 
		c.JSON(201, xpnkuserid) 
	} else {
		c.JSON(422, gin.H{"error": "No slackid was sent."})
	}		 	
}

func getXPNKUser(slackuserid string) int {
	dbmap := db_connect.InitDb()
	defer dbmap.Db.Close()
	slackid := slackuserid
	var user_xpnkid int
	err := dbmap.SelectOne(&user_xpnkid, "SELECT user_ID FROM USERS WHERE slack_userid=?", slackid)
	var xpnkid int
	if err == nil {
		xpnkid = user_xpnkid
	} else {
		fmt.Printf("\n==========\nProblemz with selecting user_xpnkid: \n%+v\n",err)
	}
	
	return xpnkid
}