package xpnk_groupTokens

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
   	"xpnk-shared/db_connect"
   	"xpnk_auth"
)

/**************************************************************************************
*
*Takes a group_id and an integer.
*Creates as many xpnk_tokens as indicated by the integer. 
*Inserts group_id and xpnk_tokens as a row in group_tokens table.
*
**************************************************************************************/

type GroupCount struct {
	XpnkGroup				int
	//MemberCount			int
	Source				string
	Identifier			string
}

type GroupTokenInsert struct {
	XpnkToken			string				`db:"xpnk_token"`
	XpnkGroup			int					`db:"group_id"`
}

func SaveGroupToken (group_count GroupCount) string {

	dbmap := db_connect.InitDb()
	defer dbmap.Db.Close()
	dbmap.AddTableWithName(GroupTokenInsert{}, "group_tokens")

	var group_token_insert 					GroupTokenInsert

	source								:= group_count.Source
	identifier							:= group_count.Identifier
	group_token_insert.XpnkGroup		= group_count.XpnkGroup
	group_token_insert.XpnkToken 		= xpnk_auth.GetNewGroupToken(source, identifier)
	
	err := dbmap.Insert(&group_token_insert)
	if err != nil {fmt.Printf("There was an error at line 44 of groupTokens.go", err)

	}
	
	return "Success!"
}	