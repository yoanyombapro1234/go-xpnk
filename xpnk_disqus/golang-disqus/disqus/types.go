package disqus

/*********************************************
* Types for the response object from Disqus
*********************************************/

type ListPostsResponse struct {
	Header			*Cursor		`json:"cursor"`
	Code			int			`json:"code"`
	Contents		[]Content	`json:"response"`
}

type Cursor struct {
	Prev			string		`json:"prev"`
	HasNext			bool		`json:"hasNext"`
	Next			string		`json:"next"`
	HasPrev			bool		`json:"hasPrev"`
	Total			string 		`json:"total"`
	Id				string 		`json:"id"`
	More			bool		`json:"more"`
}

type Content struct {
	IsHighlighted	bool		`json:"isHighlighted"`
	IsFlagged		bool		`json:"isFlagged"`
	Forum			string		`json:"forum"`
	Parent			int			`json:"parent"`
	Author			*Author		`json:"author"`
	Medias			[]Media		`json:"media"`
	Points			int			`json:"points"`
	IsApproved		bool		`json:"isApproved"`
	Dislikes		int			`json:"dislikes"`
	RawMessage		string		`json:"raw_message"`
	IsSpam			bool		`json:"isSpam"`
	Thread			string 		`json:"thread"`
	NumReports		int			`json:"numReports"`
	DeletedByAuth	bool		`json:"isDeletedByAuthor"`
	CreatedAt		string		`json:"createdAt"`
	IsEdited		bool		`json:"isEdited"`
	Message			string		`json:"message"`
	Id				string      `json:"id"`
	IsDeleted		bool		`json:"isDeleted"`
	Likes			int			`json:"likes"`
	Permalink		string
	Title			string
	Favicon			string
}

type Author	struct {
	Username		string		`json:"username"`
	About			string		`json:"about"`
	Name			string		`json:"name"`
	D3PT			bool		`json:"disable3rdPartyTrackers"`
	PowerContrib	bool		`json:"isPowerContributor"`
	Joined			string		`json:"joinedAt"`
	Rep				float32		`json:"rep"`
	ProfileUrl		string		`json:"profileUrl"`
	Url				string		`json:"url"`
	Reputation		float32		`json:"reputation"`
	Location		string		`json:"location"`
	IsPrivate		bool		`json:"isPrivate"`
	SignedUrl		string		`json:"signedUrl"`
	IsPrimary		bool		`json:"isPrimary"`
	IsAnonymous		bool		`json:"isAnonymous"`
	Id				string		`json:"id"`
	Avatar			*Avatar		`json:"avatar"`		
}

	type Avatar struct {
		Small			*SmAvatar 	`json:"small"`		
		Large			*LgAvatar 	`json:"large"`
		IsCustom		bool		`json:"isCustom"`
		Permalink		string		`json:"permalink"`
		Cache			string		`json:"cache"`
	}
	
	type SmAvatar struct {
		Permalink		string		`json:"permalink"`
		Cache			string		`json:"cache"`
	}
	type LgAvatar struct {
		Permalink		string		`json:"permalink"`
		Cache			string		`json:"cache"`
	}
	
type Media struct {
	Forum			string		`json:"forum"`
	ThumbUrl		string		`json:"thumbnailUrl"`
	Description		string		`json:"description"`
	Thread			string		`json:"thread"`
	Title			string		`json:"title"`
	Url				string		`json:"url"`
	MediaType		string		`json:"mediaType"`
	Html			string		`json:"html"`
	Location		string		`json:"location"`
	ResolvedUrl		string		`json:"resolvedUrl"`
	Post			string		`json:"post"`
	Thumbnail_URL	string		`json:"thumbnailURL"`
	Type			int			`json:"type"`
	Metadata		*Metadata	`json:"metadata"`
}	
	type Metadata struct {
		Create_method		string	`json:"create_method"`
		Thumbnail			string	`json:"thumbnail"`
	}

type MetaResponse struct {
	Meta *Meta
}
	
type Meta struct {
	Code         int
	ErrorType    string `json:"error_type"`
	ErrorMessage string `json:"error_message"`
}	

type DisqusThreadDetails struct {
	Code					int					`json:"code"`
    Contents				*Details			`json:"response"`
}	

	type Details struct {
		Link				string				`json:"link"`
		Title				string				`json:"title"`
	}