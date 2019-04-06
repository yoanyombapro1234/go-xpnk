package groups

import (
	"fmt"
	"strconv"
	"database/sql"
	 "github.com/gin-gonic/gin"
	  _ "github.com/go-sql-driver/mysql"
	 "xpnk-shared/db_connect"
)

func GroupsDelete (c *gin.Context) {
	groupid		 			:= 	c.Params.ByName("id")
	ownerid		 			:= 	c.Params.ByName("owner")
	
	groupnum, err 			:=	strconv.Atoi(groupid)
	_, err2					:=  strconv.Atoi(ownerid)
	
	if err != nil || err2 != nil {
		c.JSON(422, gin.H{"error": "One of the ids you sent is missing or wrong."})
	  	return
	}
		
	if groupnum <= 0 {
	  	c.JSON(422, gin.H{"error": "No group_id was sent."})
	  	return
	} else {
		groupdel, err 	:= delGroup(groupid, ownerid)
		if err != nil {
			 fmt.Printf("\nERROR DELETING GROUP: %+v\n", err)
			c.JSON(400, err.Error())
			return
		} else {
			fmt.Printf("\nGROUP DELETED: %+v\n", groupdel)	
			returnstring := "Group deleted: " + c.Params.ByName("id")
			c.JSON(201, returnstring)
		}	
	}		 
}

func delGroupUsers (groupID int) (sql.Result, error) {
	dbmap 					:= db_connect.InitDb()
	defer dbmap.Db.Close()
		
	res, err := dbmap.Exec("delete from USER_GROUPS where Group_ID=?", groupID)
	
	if err != nil {
		fmt.Printf("\n===========\n delGroupUsers error: %+v", err)
	} else {
		fmt.Printf("\n===========\n delGroupUsers response: %+v", res)
	}
	
	return res, err
}
func delGroup (groupID string, ownerID string) (int64, error) {
	type Group struct {
		Group_ID 			int 			`db:"Group_ID"`
		Group_Name 			string 			`db:"group_name"`
		Source				string			`db:"source"`
		Source_ID			string			`db:"source_id"`		
	}
	
	var group_id 			Group 
	groupnum, err := strconv.Atoi(groupID)

	if err != nil {
		return 0, err
	}
	group_id.Group_ID = groupnum
	fmt.Printf("\n==============\n Group_ID to be deleted: %+v", group_id.Group_ID)
	
	ownercheck, err := GroupOwner(groupID, ownerID)
	if err != nil || ownercheck == false {
		var result int64 
		return result , err 
	}
	
	dbmap 					:= db_connect.InitDb()
	defer dbmap.Db.Close()
	
	dbmap.AddTableWithName(Group{}, "GROUPS").SetKeys(true, "Group_ID")
	
	_, err = dbmap.Delete(&group_id)
	fmt.Printf("\n==============\n deleted: %+v", group_id)
	
	count, err := dbmap.SelectInt("select count(*) from GROUPS where Group_ID=?", group_id.Group_ID)
	fmt.Printf("\n==============\n COUNT: %+v", count)
	
	res, err2 := delGroupUsers(group_id.Group_ID)
	if err2 != nil {
		fmt.Printf("\n===========\n delGroup error: %+v", err)
	} else {
		fmt.Printf("\n===========\n delGroup response: %+v", res)
	}
	
	return count, err
}