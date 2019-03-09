package xpnk_createGroupInsert

/**************************************************************************************
Takes a group_name, source and source_id and prepares them for insertion into the 
database by mapping each item to a database field
**************************************************************************************/

import (
	"fmt"
)

//incoming data for new group to be created
type New_Group struct {
	GroupName		string						
	SourceID		string						
	Source			string						
}

//stores New_Group data into struct mapped for insertion into db
type Group_Insert struct {
	GroupName		string						`db:"group_name"`
	SourceID		string						`db:"source_id"`
	Source			string						`db:"source"`
	GroupID			int							`db:"Group_ID"`		
}

func CreateGroupInsert(newGroup New_Group) Group_Insert {
//func CreateGroupInsert(newGroup New_Group) string {

		var this_group_insert Group_Insert
		
		this_group_insert.GroupName			= newGroup.GroupName
		this_group_insert.SourceID			= newGroup.SourceID
		this_group_insert.Source	 		= newGroup.Source
						
		fmt.Printf("\n==========\nTHIS_GROUP_INSERT: \n%+v\n",this_group_insert)
		
		return this_group_insert
}