package groups

import (
	"fmt"
	"strconv"
	 "github.com/gin-gonic/gin"
	  _ "github.com/go-sql-driver/mysql"
   	 "database/sql"
	 "xpnk-shared/db_connect"
)

func GroupsMemberDelete (c *gin.Context) {
	groupid 			:= 	c.Params.ByName("id")
	userid				:= 	c.Params.ByName("user")
	ownerid				:= 	c.Params.ByName("owner")
	
	groupnum, err		:= strconv.Atoi(groupid)
	usernum, err2		:= strconv.Atoi(userid)
	ownernum, err3		:= strconv.Atoi(ownerid)
	
	if err != nil || err2 != nil || err3 != nil {
		c.JSON(422, gin.H{"error": "One of the ids you sent is missing or wrong."})
	  	return
	}
	
	if groupnum <= 0 || usernum <= 0 || ownernum <= 0 {
	  	c.JSON(422, gin.H{"error": "One of the ids you sent is missing or wrong."})
	  	return
	} else {
		memberdel, err 	:= delMember(groupid, userid, ownerid)
		if err != nil {
			 fmt.Printf("\nERROR DELETING MEMBER: %+v\n", err)
			c.JSON(400, err.Error())
			return
		} else {
			fmt.Printf("\nUSER REMOVED FROM GROUP: %+v\n", memberdel)	
			returnstring := "User removed: " + c.Params.ByName("user")
			c.JSON(201, returnstring)
		}	
	}		 
}

func delMember (groupID string, userID string, ownerID string) (sql.Result, error) {	

	ownercheck, err := GroupOwner(groupID, ownerID)
	if err != nil || ownercheck == false {
		var result sql.Result 
		return result , err 
	}
	
	dbmap 					:= db_connect.InitDb()
	defer dbmap.Db.Close()

	delete_query := "delete from USER_GROUPS where Group_ID=" + groupID + " AND user_ID=" + userID 
	res, err := dbmap.Exec(delete_query)
	return res, err 
}