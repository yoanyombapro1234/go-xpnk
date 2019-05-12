package xpnk_group_structs

type NewestPosts struct {
	GroupId					int				`db:"Group_ID"`
	Tweet					string			`db:"tweet"`
	Instagram				int				`db:"instagram"`
	Disqus					string			`db:"disqus"`
}
