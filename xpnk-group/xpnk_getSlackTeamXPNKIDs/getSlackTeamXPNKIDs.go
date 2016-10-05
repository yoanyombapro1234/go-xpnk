package xpnk_getSlackTeamXPNKIDs

/**************************************************************************************
*
*Takes a slice of xpnk_createUserInsert.User_Insert. For each SlackID queries db for 
*the XPNK User_ID. Returns an array of SlackID : User_ID.
*
**************************************************************************************/

import (
   	"xpnk-shared/db_connect"
   	"xpnk-user/xpnk_createUserInsert"
	"fmt"
   	"log"
)

type Slack_User struct {
	SlackID				string			`db:"slack_userid"`
	XPNK_ID				int				`db:"user_ID"`
}

func GetSlackUserXPNKID(team []xpnk_createUserInsert.User_Insert) []Slack_User{

	fmt.Printf("\n==========\nteam is now:%+v\n",team)
	
	var thisTeam []Slack_User
	
	for i :=0; i <len(team); i++ {
		var thisMember Slack_User
	
		thisMember.SlackID = team[i].SlackID
		
		dbmap := db_connect.InitDb()
		defer dbmap.Db.Close()
		
		var xpnk_ID int
		err := dbmap.SelectOne(&xpnk_ID, "SELECT user_ID FROM USERS WHERE slack_userid='" + thisMember.SlackID + " ' ")
		if err == nil {
	    	fmt.Printf("\n==========\nXPNK_ID: %+v", xpnk_ID)
	    	fmt.Printf("\n==========\nSlack_ID: %+v", thisMember.SlackID)
		} else {
			fmt.Printf("\n==========\nProblemz with select: \n%+v\n",err)
		}
		
		thisMember.XPNK_ID = xpnk_ID	
		
		thisTeam = append(thisTeam, thisMember)	
	}
	
	fmt.Printf("\n==========\nthisTeam is now:%+v\n",thisTeam)

	
		return thisTeam	
}
//end db insert routine	
//Maps are not inherently safe for concurrency - will have to use sync.RWMutex	

func checkErr(err error, msg string) {
  if err != nil {
  log.Fatalln(msg, err)
	}	
} 
