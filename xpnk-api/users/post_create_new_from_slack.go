package users

import (
	"fmt"
	 _ "github.com/go-sql-driver/mysql"
	 "github.com/gin-gonic/gin"
	 "xpnk-shared/db_connect"
)

type NewSlackAuth struct {
	 Slack_accesstoken	string					`json:"access_token"`
	 Slack_userid		string					`json:"slack_userid"`
	 Slack_username		string					`json:"slack_name"`
	 Slack_avatar		string					`json:"slack_avatar"`
}

type NewSlackAuthInsert struct {
	 Slack_accesstoken	string					`db:"slack_authtoken"`
	 Slack_userid		string					`db:"slack_userid"`
	 Slack_username		string					`db:"slack_name"`
	 Slack_avatar		string					`db:"profile_image"`
	 Xpnk_id			int						`db:"user_ID"`
}

func SlackNewMember (c *gin.Context) {
	var slackuser NewSlackAuth
	c.Bind(&slackuser)
	
	if slackuser.Slack_username != "" {	
		xpnkuser := updateSlackUser(slackuser)	
		fmt.Printf("\nXPNKUSER RETURNED BY UPDATESLACKUSER: %+v\n", xpnkuser)	 
		c.JSON(200, xpnkuser)
		 
	} else {
		c.JSON(422, gin.H{"error": "No access token was sent."})
	}		 	
}

func updateSlackUser(new_Slackauth NewSlackAuth) int{
	slack_id := new_Slackauth.Slack_userid
	
	dbmap := db_connect.InitDb()
	defer dbmap.Db.Close()
	dbmap.AddTableWithName(NewSlackAuthInsert{}, "USERS").SetKeys(true, "Xpnk_id")
	
	var user_xpnkid int
	err := dbmap.SelectOne(&user_xpnkid, "SELECT user_ID FROM USERS WHERE slack_userid=?", slack_id)
	
	fmt.Printf("\n==========\nUSER_XPNKID: \n%+v\n",user_xpnkid)
	
	if err == nil {
		var new_Slackauthinsert NewSlackAuthInsert
		new_Slackauthinsert.Xpnk_id = user_xpnkid
		new_Slackauthinsert.Slack_accesstoken = new_Slackauth.Slack_accesstoken
	 	new_Slackauthinsert.Slack_userid = new_Slackauth.Slack_userid	 
	 	new_Slackauthinsert.Slack_username = new_Slackauth.Slack_username
	 	new_Slackauthinsert.Slack_avatar = new_Slackauth.Slack_avatar
		
		_, dberr := dbmap.Update(&new_Slackauthinsert)
		if dberr == nil {
			fmt.Printf("\n==========\nNewSlackAuth Update Success!")
		} else {
			fmt.Printf("\n==========\nProblemz with update: \n%+v\n",dberr)
		}
	} else {
		fmt.Printf("\n==========\nProblemz with select: \n%+v\n",err)
	}
	
	return user_xpnkid
} 