package xpnk_createGroup

/**************************************************************************************
Queen file for setting up a new Xapnik group from the Xapnik app
1) Receives request details  
2) Creates a new Xapnik Group from the request data
3) Adds the group id and owner id to the USER_GROUPS table with owner marked as owner
4) Returns the newly created group id, group name, and slug
**************************************************************************************/

import (
	"fmt"
	"xpnk-group/xpnk_createGroupInsert"
	"xpnk-group/xpnk_insertGroup2"
	"xpnk-group/xpnk_addUsersToGroup2"
)

type NewGroup struct {
	Owner					string				`form:"Owner" 		binding:"required"`
	GroupName				string				`form:"GroupName" 	binding:"required"`
	Source					string				`form:"Source"		binding:"required"`
	SourceID				string				`form:"SourceID"`
} 

type NewGroupReturn struct {
	Owner					string				`db:"group_owner" 	json:"group_owner"`
	GroupID					int					`db:"Group_ID" 		json:"group_id"`
	GroupSlug				string								
}

func CreateGroup(newGroup NewGroup) (int, error) {
	
	var groupInsert 		xpnk_createGroupInsert.New_Group
	groupInsert.GroupName	= newGroup.GroupName
	groupInsert.Source		= newGroup.Source
	groupInsert.SourceID	= newGroup.SourceID
	
	insertCreated			:= xpnk_createGroupInsert.CreateGroupInsert(groupInsert)
	
	fmt.Printf("\n==========\nNEW GROUP: \n%+v\n", insertCreated)
		
	inserted, err 			:= xpnk_insertGroup2.InsertGroup2(insertCreated)
	
	var userGroup xpnk_addUsersToGroup2.Group_User
	userGroup.User = newGroup.Owner
	userGroup.Group = inserted
	userGroup.Owner = true
	fmt.Printf("\n==========\nGROUP MEMBER TO ADD: \n%+v\n", userGroup)
		
	ownerAdded, err2		:= xpnk_addUsersToGroup2.AddUsersToGroup(userGroup)
	if err2 != nil {
		fmt.Printf("\n==========\nUSER_GROUPS INSERT ERROR: \n%+v\n", err2)
		return 0, err2
	}
	fmt.Printf("\n==========\nNEW GROUP MEMBER: \n%+v\n", ownerAdded)
	
	return inserted, err
}