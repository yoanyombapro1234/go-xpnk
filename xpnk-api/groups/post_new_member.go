package groups

import (
	"fmt"
	 "github.com/gin-gonic/gin"
	 "xpnk-shared/db_connect"
)

type NewGroupMember		struct {
	Group_ID			int						`form:"id"		json:"id"`
	User_ID				int						`form:"userId" 	json:"userId"`				
}

type NewGroupMemberInsert	struct {
	Group_ID			int						`db:"Group_ID"`
	User_ID				int						`db:"user_ID"`				
}

func GroupsAddMember (c *gin.Context) {
	var new_groupMember		NewGroupMember
	c.Bind(&new_groupMember)
	fmt.Printf("\n new_groupMember.Group_ID:  %v \n", new_groupMember.Group_ID)
	fmt.Printf("new_groupMember.User_ID:  %v \n", new_groupMember.User_ID)
	if new_groupMember.Group_ID < 1 || new_groupMember.User_ID < 1 {
		c.JSON(422, gin.H{"error": "Group id or user id or both are missing."})
	} else {
		var new_member_insert   NewGroupMemberInsert
		new_member_insert.Group_ID	= new_groupMember.Group_ID
		new_member_insert.User_ID	= new_groupMember.User_ID
	
		insert_new_member 			:= InsertNewGroupMember(new_member_insert)
	
		if insert_new_member == 1 {
			c.JSON(201, "User added!")
		} else {
			c.JSON(422, gin.H{"error": insert_new_member })
		}
	}
}

func InsertNewGroupMember(new_GroupMember NewGroupMemberInsert) int {
	var returnVal 				int
	dbmap := db_connect.InitDb()
	defer dbmap.Db.Close()
	dbmap.AddTableWithName(NewGroupMemberInsert{}, "USER_GROUPS")
		
	err := dbmap.Insert(&new_GroupMember)	
		if err == nil {
			fmt.Printf("\n==========\nInsertNewGroupMember Update Success!")
			returnVal = 1
		} else {
			fmt.Printf("\n==========\nProblemz with InsertNewGroupMember line 700 in api: \n%+v\n",err)
			returnVal = 0
		}
	return returnVal	
}