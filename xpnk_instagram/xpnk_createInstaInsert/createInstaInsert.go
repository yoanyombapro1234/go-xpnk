package xpnk_createInstaInsert

/**************************************************************************************
Takes a slice of instaPosts and prepares them for insertion into the datatbase by mapping each item to a database field
**************************************************************************************/

import (
	"fmt"
	"xpnk_instagram/xpnk_getInstaEmbed"
	"github.com/yanatan16/golang-instagram/instagram"
	//"time"
	"github.com/chnlr/emoji"
)

//stores Instagram post data from json into struct mapped for insertion into db
type Instagram_Insert struct {
	InstagramUser	string						`db:"insta_user"`
	InstagramName	string						`db:"insta_name"`
	InstagramUserID	string						`db:"insta_userid"`
	InstagramUrl	string						`db:"instagram_url"`
	InstagramOembed	string						`db:"instagram_oembed"`
	InstagramDate	string						`db:"instagram_date"`
	InstagramAvatar	string						`db:"instagram_avatar"`
}

//store each Instagram and user name in a slice of structs to then insert into db
//1) store the single post data in an Instagram_Insert struct
//2) append the Instagram_Insert into a slice of Instagram_Inserts

func CreateInstaInsert(instaPosts *instagram.PaginatedMediasResponse) []Instagram_Insert {
	var Instagram_Inserts []Instagram_Insert

	for i := 0; i < len(instaPosts.Medias); i++ {
		//first put stuff from instaPosts into Instagram_Insert
		//then get the oembed and put into Instagram_Insert
		//Get the post ID for each Instagram post

		var this_insta_insert Instagram_Insert

		this_insta_insert.InstagramUser 	= instaPosts.Medias[i].User.Username
		this_insta_insert.InstagramName 	= instaPosts.Medias[i].User.FullName
		this_insta_insert.InstagramUserID 	= instaPosts.Medias[i].User.Id
		this_insta_insert.InstagramUrl 		= instaPosts.Medias[i].Link
		this_insta_insert.InstagramAvatar	= instaPosts.Medias[i].User.ProfilePicture
	
		//convert created_time time format to time.Time format
		//this_created_at := instaPosts.Medias[i].CreatedTime
		//this_created_date, _ := time.Parse(time.RubyDate,this_created_at)

		this_insta_insert.InstagramDate 	= string(instaPosts.Medias[i].CreatedTime)
				
		//Get the oembed code for each post, has to be queried separately		
		getoembed := instaPosts.Medias[i].Link
		fmt.Printf("\n==========\nIDSTRING: \n%+v\n",getoembed)		

		embed := xpnk_getInstaEmbed.GetInstaEmbed(getoembed)
		//if err != nil {
		//	fmt.Printf("GetOembed returned error: %s", err.Error())
		//}
		//fmt.Printf("\n==========\nEMBED: \n%+v\n",embed.Html)
	
		this_insta_insert.InstagramOembed = emoji.UnicodeToTwemoji(embed.Html, 16, false)
		fmt.Printf("\n==========\nTHIS_INSERT: \n%+v\n",this_insta_insert)

		Instagram_Inserts = append(Instagram_Inserts, this_insta_insert)
	}//end looping through all posts
	
	fmt.Printf("\n==========\nTHIS_BATCH: \n%+v\n",Instagram_Inserts)
	
	return Instagram_Inserts
	
}//end createInstaInsert	