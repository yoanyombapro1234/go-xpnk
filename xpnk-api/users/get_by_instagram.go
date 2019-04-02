package users

import (
	"fmt"
	 _ "github.com/go-sql-driver/mysql"
	 "github.com/gin-gonic/gin"
	 "xpnk-shared/db_connect"
)

type IGID struct {
	 IG_userid			string					`form:"id" binding:"required"`
}

func UsersByIGID_2 (c *gin.Context) {
	ig_ID					 := c.Param("id")
	var user					XPNKUser
	var err_msg					error
	if ig_ID == "" {
		c.JSON(422, gin.H{"error": "Invalid or missing Instagram user ID."})
	} else {
		user, err_msg 		  = get_user_by_ig(ig_ID)
		if user.Insta_userid != ig_ID {
			c.JSON(400, err_msg.Error())
		} else {
			c.JSON(200, user)
		}
	}
}

func UsersByIGID (c *gin.Context) {
	var igID					IGID
	var user					XPNKUser
	var err_msg					error
	c.Bind(&igID)
	ig_id						:= igID.IG_userid
	fmt.Printf("ig_id: %v \n", ig_id)
	if ig_id == "" {
		c.JSON(422, gin.H{"error": "Invalid or missing Instagram user ID."})
	} else {
		user, err_msg 			= get_user_by_ig(ig_id)
		if user.Insta_userid != ig_id {
			c.JSON(202, user)
			fmt.Printf("error:  %v \n", err_msg)
		} else {
			c.JSON(200, user)
		}
	}
}

func get_user_by_ig(ig_id string) (XPNKUser, error) {
	dbmap := db_connect.InitDb()
	defer dbmap.Db.Close()
	
	var xpnkUser			XPNKUser
	var err_msg				error
	IGId					:= ig_id
	err	:= dbmap.SelectOne(&xpnkUser, "SELECT `user_ID`, `twitter_user`, `twitter_ID`, `twitter_authtoken`, `twitter_secret`, `insta_user`, `insta_userid`, `insta_accesstoken`, `disqus_username`, `disqus_userid`, `disqus_accesstoken`, `profile_image` FROM USERS WHERE insta_userid=?", IGId)
	if err != nil {
		fmt.Printf("\n==========\nget_user_by_ig - Problemz with selecting user by IGId: \n%+v\n",err)
		err_msg = err
	} else {
		fmt.Printf("\n==========\nfound user: \n%+v\n",xpnkUser.User_ID)
	}
	return xpnkUser, err_msg
}
