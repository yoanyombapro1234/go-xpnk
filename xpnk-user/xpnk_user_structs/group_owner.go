package xpnk_user_structs

import (
	"database/sql"
)

type GroupOwner 		struct {
	Group_ID			int				`db:"Group_ID"			json:"Group_ID"`
	Owner				sql.NullBool	`db:"group_owner"		json:"group_owner"`
	Admin				sql.NullBool	`db:"group_admin"		json:"group_admin"`
	Name				string			`db:"group_name"		json:"group_name"`
	Slug				string									`json:"group_slug"`		
}