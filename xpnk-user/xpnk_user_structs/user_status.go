package xpnk_user_structs

type UserStatus struct {
	TwitterLoginNeeded		bool
	InstagramLoginNeeded	bool
	DisqusLoginNeeded		bool
	UserGroups				UserGroups
}
