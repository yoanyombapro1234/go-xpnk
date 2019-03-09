package twitter_verify

import (
	"fmt"
   	anaconda "github.com/ChimeraCoder/anaconda"
   	"net/url"
   	"xpnk_constants"
)

func AccountVerify (token string, secret string) (string, error) {

	anaconda.SetConsumerKey(xpnk_constants.TwitterKey)
	anaconda.SetConsumerSecret(xpnk_constants.TwitterSec)
	api := anaconda.NewTwitterApi(token, secret)
	
	user, err := api.GetSelf(url.Values{})
	
	if err != nil {
		fmt.Printf("GetSelf in xpnk_twitter account_verify returned an error: %s", err.Error())
	}
	
	screenname := user.ScreenName
	userid := user.IdStr
	
	fmt.Printf("\nTwitter user: %+s\n", screenname)
	fmt.Printf("\nTwitter userid: %+v\n", userid)
	
	return userid, err 
}