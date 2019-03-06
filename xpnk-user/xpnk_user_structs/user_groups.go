package xpnk_user_structs

type UserGroups 		struct {
	Xpnk_id				string			`db:"user_ID"			json:"user_ID"`
	Groups				[]GroupsByUser
}