package xpnk_user_structs

type UserSub struct {
	Id						int				`db:"user_ID"		json:"Id"`
	Endpoint				string			`db:"endpoint"		json:"Endpoint"`
	Type					int				`db:"type"			json:"Type"`
	P256dh					string			`db:"p256dh"		json:"P256dh""`
	Auth					string			`db:"auth"			json:"Auth"`
}
