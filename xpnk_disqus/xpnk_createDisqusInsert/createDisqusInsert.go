package xpnk_createDisqusInsert

/**************************************************************************************
Takes a slice of disqusPosts and prepares them for insertion into the datatbase by mapping each item to a database field
**************************************************************************************/

import (
	"fmt"
	"xpnk_disqus/golang-disqus/disqus"
	"html"
)

//stores Disqus post data from json into struct mapped for insertion into db
type Disqus_Insert struct {
	DisqusForum		string						`db:"disqus_forum"`
	DisqusFavicon	string						`db:"disqus_favicon"`
	DisqusUser		string						`db:"disqus_user"`
	DisqusName		string						`db:"disqus_name"`
	DisqusUserID	string						`db:"disqus_userid"`
	DisqusPID		string						`db:"disqus_pid"`
	DisqusPermalink	string						`db:"disqus_permalink"`
	DisqusTitle		string						`db:"disqus_title"`
	DisqusEmbed		string						`db:"disqus_embed"`
	DisqusDate		string						`db:"disqus_date"`
	DisqusAvatar	string						`db:"disqus_avatar"`
	DisqusMedia		string						`db:"disqus_media"`
}

//store each Disqus comment and user name in a slice of structs to then insert into db
//1) store the single post data in a Disqus_Insert struct
//2) append the Disqus_Insert into a slice of Disqus_Inserts

func CreateDisqusInsert(disqusComments []disqus.Content) []Disqus_Insert {
	var Disqus_Inserts []Disqus_Insert

	for i := 0; i < len(disqusComments); i++ {
		//put stuff from disqusComments into Disqus_Insert
		//then get the oembed and put into Disqus_Insert
		//Get the post ID for each Disqus post

		var this_disqus_insert Disqus_Insert
		
		this_disqus_insert.DisqusForum		= disqusComments[i].Forum
		this_disqus_insert.DisqusFavicon	= "https://disqus.com/api/forums/favicons/" + this_disqus_insert.DisqusForum + ".jpg" 
		this_disqus_insert.DisqusUser 		= disqusComments[i].Author.Username
		this_disqus_insert.DisqusName 		= disqusComments[i].Author.Name
		this_disqus_insert.DisqusUserID 	= disqusComments[i].Author.Id
		this_disqus_insert.DisqusPID		= disqusComments[i].Id
		this_disqus_insert.DisqusPermalink 	= disqusComments[i].Permalink
		this_disqus_insert.DisqusTitle	 	= html.UnescapeString(disqusComments[i].Title)
		this_disqus_insert.DisqusEmbed		= html.UnescapeString(
											disqusComments[i].Message)
		this_disqus_insert.DisqusAvatar		= disqusComments[i].Author.Avatar.Permalink
		this_disqus_insert.DisqusDate 		= disqusComments[i].CreatedAt
		
		if len(disqusComments[i].Medias) > 0 { 
			this_disqus_insert.DisqusMedia	= "http:" + disqusComments[i].Medias[0].ThumbUrl
		}
				
		//this_disqus_insert.DisqusEmbed = emoji.UnicodeToTwemoji(embed.Html, 16, false)
		//TODO see if emoji translate properly in Disqus posts
				
		fmt.Printf("\n==========\nTHIS_INSERT: \n%+v\n",this_disqus_insert)

		Disqus_Inserts = append(Disqus_Inserts, this_disqus_insert)
	}//end looping through all posts
	
	fmt.Printf("\n==========\nTHIS_BATCH: \n%+v\n",Disqus_Inserts)
	
	return Disqus_Inserts
	
}//end createDisqusInsert	