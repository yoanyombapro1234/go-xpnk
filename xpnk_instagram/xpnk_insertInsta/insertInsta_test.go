package xpnk_insertInsta

import  (
	"testing"
	"xpnk_instagram/xpnk_createInstaInsert"
)	

func TestInsertInsta(t *testing.T) {

	var insert []xpnk_createInstaInsert.Instagram_Insert
	var thisinsert xpnk_createInstaInsert.Instagram_Insert
	
	thisinsert.InstagramUser = "khloekardashian"
	thisinsert.InstagramName = "Khloé"
	thisinsert.InstagramUserID = "208560325"
	thisinsert.InstagramUrl = "https://www.instagram.com/p/BDboYEghRgR/"
	thisinsert.InstagramOembed = "<blockquote class=\"instagram-media\" data-instgrm-version=\"6\" style=\" background:#FFF; border:0; border-radius:3px; box-shadow:0 0 1px 0 rgba(0,0,0,0.5),0 1px 10px 0 rgba(0,0,0,0.15); margin: 1px; max-width:658px; padding:0; width:99.375%; width:-webkit-calc(100% - 2px); width:calc(100% - 2px);\"><div style=\"padding:8px;\"> <div style=\" background:#F8F8F8; line-height:0; margin-top:40px; padding:37.3148148148% 0; text-align:center; width:100%;\"> <div style=\" background:url(data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAACwAAAAsCAMAAAApWqozAAAAGFBMVEUiIiI9PT0eHh4gIB4hIBkcHBwcHBwcHBydr+JQAAAACHRSTlMABA4YHyQsM5jtaMwAAADfSURBVDjL7ZVBEgMhCAQBAf//42xcNbpAqakcM0ftUmFAAIBE81IqBJdS3lS6zs3bIpB9WED3YYXFPmHRfT8sgyrCP1x8uEUxLMzNWElFOYCV6mHWWwMzdPEKHlhLw7NWJqkHc4uIZphavDzA2JPzUDsBZziNae2S6owH8xPmX8G7zzgKEOPUoYHvGz1TBCxMkd3kwNVbU0gKHkx+iZILf77IofhrY1nYFnB/lQPb79drWOyJVa/DAvg9B/rLB4cC+Nqgdz/TvBbBnr6GBReqn/nRmDgaQEej7WhonozjF+Y2I/fZou/qAAAAAElFTkSuQmCC); display:block; height:44px; margin:0 auto -44px; position:relative; top:-22px; width:44px;\"></div></div><p style=\" color:#c9c8cd; font-family:Arial,sans-serif; font-size:14px; line-height:17px; margin-bottom:0; margin-top:8px; overflow:hidden; padding:8px 0 7px; text-align:center; text-overflow:ellipsis; white-space:nowrap;\"><a href=\"https://www.instagram.com/p/BDboYEghRgR/\" style=\" color:#c9c8cd; font-family:Arial,sans-serif; font-size:14px; font-style:normal; font-weight:normal; line-height:17px; text-decoration:none;\" target=\"_blank\">A photo posted by Khloé (@khloekardashian)</a> on <time style=\" font-family:Arial,sans-serif; font-size:14px; line-height:17px;\" datetime=\"2016-03-26T21:42:44+00:00\">Mar 26, 2016 at 2:42pm PDT</time></p></div></blockquote><script async defer src=\"//platform.instagram.com/en_US/embeds.js\f\"></script>"
	thisinsert.InstagramDate = "1459028564"
	thisinsert.InstagramAvatar = "https://scontent.cdninstagram.com/t51.2885-19/s150x150/12751281_569756159856278_1434520857_a.jpg"
	
	insert = append(insert, thisinsert)

	v := InsertInsta(insert)
	
	if v != "inserted" {
		t.Errorf("Expected INSERTED, got  ", v)
	}
				
}