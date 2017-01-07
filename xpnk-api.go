package main

import (
         "github.com/gin-gonic/gin"
         "fmt"
         "net/http"
   		 _ "github.com/go-sql-driver/mysql"
   		 "xpnk_auth"
   		 "xpnk-shared/db_connect"
 )
 
type Usertoken struct {
	Token				string					`json:"token"`
} 

type Slack_ID struct {
	Slack_id			string					`json:"slack_id"`
}
 
type NewSlackAuth struct {
	 Slack_accesstoken	string					`json:"access_token"`
	 Slack_userid		string					`json:"slack_userid"`
	 Slack_username		string					`json:"slack_name"`
}

type NewSlackAuthInsert struct {
	 Slack_accesstoken	string					`db:"slack_authtoken"`
	 Slack_userid		string					`db:"slack_userid"`
	 Slack_username		string					`db:"slack_name"`
	 Xpnk_id			int						`db:"user_ID"`
}


type NewTwitterAuth struct {
	 Twttr_accesstoken	string					`json:"access_token"`
	 Twttr_secret		string					`json:"user_secret"`
	 Twttr_userid		string					`json:"twitter_userid"`
	 Xpnk_id			string						
}

type NewTwitterAuthInsert struct {
	 Twttr_accesstoken	string					`db:"twitter_authtoken"`
	 Twttr_secret		string					`db:"twitter_secret"`
	 Twttr_userid		string					`db:"twitter_ID"`
	 Xpnk_id			string					`db:"user_ID"`
}

type NewIGAuth struct {
	 Insta_accesstoken	string					`json:"access_token"`
	 Insta_userid		string					`json:"insta_userid"`
	 Insta_username		string					`json:"insta_username"`
	 Xpnk_id			string
}

type NewIGAuthInsert struct {
	 Insta_accesstoken	string					`db:"insta_accesstoken"`
	 Insta_userid		string					`db:"insta_userid"`
	 Insta_username		string					`db:"insta_user"`
	 Xpnk_id			string					`db:"user_ID"`
}

type NewDisqusAuth struct {
	 Disqus_accesstoken	string					`json:"access_token"`
	 Disqus_userid		string					`json:"disqus_userid"`
	 Disqus_username	string					`json:"disqus_username"`
	 Xpnk_id			string
}

type NewDisqusAuthInsert struct {
	 Disqus_accesstoken	string					`db:"disqus_accesstoken"`
	 Disqus_userid		string					`db:"disqus_userid"`
	 Disqus_username	string					`db:"disqus_username"`
	 Xpnk_id			string					`db:"user_ID"`
}

const (
	mySigningKey = "lakdjfiafjeoijaldknamnf823984udkafdjasdf"
)
	 

func main() {

	r := gin.Default()
	r.Use(Cors())
	
/*****************************************
*
* Endpoints
*
*****************************************/
	
	r.GET("/ping", func(c *gin.Context) {
        c.String(200, "Hello")
      })  
	
	v1 := r.Group("api/v1")
		{
			v1.OPTIONS ("/ping", func(c *gin.Context) {
				c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, PUT")
 				c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
 				c.Next()
			})
			v1.GET ("/ping", func(c *gin.Context) {
				c.String(200, "Hello there.")
			})
			
			v1.OPTIONS ("/xpnk_auth_set", func(c *gin.Context) {
				c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, PUT")
 				c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, token, xpnkid")
 				c.Next()
			})
			v1.GET ("/xpnk_auth_set", XPNKAuthSet)
			
			v1.OPTIONS ("/xpnk_auth_check", func(c *gin.Context) {
				c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, PUT")
 				c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, token, xpnkid")
 				c.Next()
			})
			v1.POST ("/xpnk_auth_check", XPNKAuthCheck)
			
			v1.OPTIONS ("/get_xpnkid/slack/:id", func(c *gin.Context) {
				c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, PUT")
 				c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, token, xpnkid")
 				c.Next()
			})
			v1.GET ("/get_xpnkid/slack/:id", GetXPNK_ID)
			
			/*
			v1.OPTIONS ("/xpnk_read_header", func(c *gin.Context) {
				c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, PUT")
 				c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, token, xpnkid")
 				c.Next()
			})
			v1.POST ("/xpnk_read_header", XPNKReadHeader)
			*/
			
			/*
			V1.OPTIONS ("/slack_new_group", func(c *gin.Context) {
			    c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, PUT")
 				c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
 				c.Next()
			})
			V1.POST ("slack_new_group", SlackNewGroup)
			*/
			
			v1.OPTIONS ("/slack_new_member", func(c *gin.Context) {
			    c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, PUT")
 				c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, token, xpnkid")
 				c.Next()
			})
			v1.POST ("slack_new_member", SlackNewMember)
			
			v1.OPTIONS ("/twitter_auth", func(c *gin.Context) {
				c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, PUT")
 				c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, token, xpnkid")
 				c.Next()
			})
			v1.POST("/twitter_auth", PostTwttrAuth)
			
			v1.OPTIONS ("/ig_auth", func(c *gin.Context) {
				c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, PUT")
 				c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, token, xpnkid")
 				c.Next()
			})
			v1.POST("/ig_auth", PostIGAuth)
			
			v1.OPTIONS ("/disqus_auth", func(c *gin.Context) {
				c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, PUT")
 				c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, token, xpnkid")
 				c.Next()
			})
			v1.POST("/disqus_auth", PostDisqusAuth)
		}
		
	r.Run(":9090")
}	

func Cors() gin.HandlerFunc {
 return func(c *gin.Context) {
 c.Writer.Header().Add("Access-Control-Allow-Origin", "*")
 c.Next()
 }
}

/*****************************************
*
* Endpoint functions
* 
*****************************************/

func XPNKAuthSet (c *gin.Context)  {
	usertoken, err := xpnk_auth.NewToken([]byte(mySigningKey))
	if usertoken != "" {
		response := usertoken
		c.JSON(201, response)
	} 	else {
		fmt.Printf("ERROR: %v+", err)
		c.JSON(422,gin.H{"error":"No access token created."})
	}
}

func XPNKReadHeader (c *gin.Context) int{
	var this_header http.Header
	this_header = c.Request.Header
	fmt.Printf("HEADER: %+v", this_header)
	token := this_header["Token"]
	fmt.Printf("TOKEN: %+v", token)
	if len(token) != 0  {
		fmt.Printf("TOKEN ONLY:  %+v", this_header["Token"][0])
		auth := xpnk_auth.ParseToken(token[0], mySigningKey)
		if auth == 1 {
			return 1
		} else {
			fmt.Printf("INVALID TOKEN:  %+V", this_header["Token"])
			return 0
		}
	} else {
		fmt.Printf("NO TOKEN FOUND:  %+V", this_header["Token"])
		return 0
	}
}

func XPNKAuthCheck (c *gin.Context) {
	fmt.Printf("HEADER: %+v", c.Request.Header)
	var this_header http.Header
	this_header = c.Request.Header
	token := this_header["Token"]
	if len(token) != 0  {
		fmt.Printf("TOKEN ONLY:  %+v", this_header["Token"][0])
		auth := xpnk_auth.ParseToken(token[0], mySigningKey)
		if auth == 1 {
			c.JSON(200, gin.H{"success":"You're clear for take off."})
		} else {
			c.JSON(422, gin.H{"error": "Token can't be authenticated."})
		}
	} else {
		c.JSON(422, gin.H{"error": "No access token was sent."})
	}
}

func GetXPNK_ID (c *gin.Context) {
	var slackuserid string
	slackuserid = c.Param("id")
	
	if slackuserid != "" {
		xpnkuserid := getXPNKUser(slackuserid)	
		fmt.Printf("\nXPNKUSERID RETURNED BY GetXPNK_ID: %+v\n", xpnkuserid)	 
		c.JSON(201, xpnkuserid)
		 
	} else {
		c.JSON(422, gin.H{"error": "No slackid was sent."})
	}		 	
}

func SlackNewMember (c *gin.Context) {
	var slackuser NewSlackAuth
	c.Bind(&slackuser)
	
	if slackuser.Slack_username != "" {	
		xpnkuser := updateSlackUser(slackuser)	
		fmt.Printf("\nXPNKUSER RETURNED BY UPDATESLACKUSER: %+v\n", xpnkuser)	 
		c.JSON(200, xpnkuser)
		 
	} else {
		c.JSON(422, gin.H{"error": "No access token was sent."})
	}		 	
}

func PostTwttrAuth(c *gin.Context) {
	
	auth := XPNKReadHeader(c)
	fmt.Printf("\n==========\nauth: %v \n", auth)
	if auth != 1 {
		c.JSON(422, gin.H{"error": "Invalid or missing xapnik token."})
	} else {

		fmt.Printf("\n==========\nPostTwttrAuth engaged \n") 
		
		var this_header http.Header
		this_header = c.Request.Header
		xpnk_id_object := this_header["Xpnkid"]
		var xpnk_id string
		if len(xpnk_id_object) == 0 {
			c.JSON(422, gin.H{"error" : "Missing or zero-value xpnkid in header."})
		} else {
			xpnk_id = xpnk_id_object[0]
		}	
		
		var twitterauth NewTwitterAuth
		c.Bind(&twitterauth)
				
		twitterauth.Xpnk_id = xpnk_id
		
		fmt.Printf("\n==========\ntwitterauth object:  %+v \n", twitterauth)
		
		if twitterauth.Twttr_accesstoken != "" && twitterauth.Twttr_secret != "" {
				 
			 c.JSON(201, "Updating Twitter info to user account.")
			 updateTwitter(twitterauth)
		} else {
			c.JSON(422, gin.H{"error": "No access token was sent."})
		}
	}			 	
}

func PostIGAuth(c *gin.Context) {
	
	auth := XPNKReadHeader(c)
	fmt.Printf("\n==========\nauth: %v \n", auth)
	if auth != 1 {
		c.JSON(422, gin.H{"error": "Invalid or missing xapnik token."})
	} else {
	
		var this_header http.Header
		this_header = c.Request.Header
		xpnk_id_object := this_header["Xpnkid"]
		var xpnk_id string
		if len(xpnk_id_object) == 0 {
			c.JSON(422, gin.H{"error" : "Missing or zero-value xpnkid in header."})
		} else {
			xpnk_id = xpnk_id_object[0]
		}	

		var igauth NewIGAuth
		c.Bind(&igauth)
		
		igauth.Xpnk_id  = xpnk_id
		
		if igauth.Insta_accesstoken != "" {
					 
			 c.JSON(201, "Insta access token received.")
			 updateIG(igauth)
		} else {
			c.JSON(422, gin.H{"error": "No access token was sent."})
		}
	}			 	
}

func PostDisqusAuth(c *gin.Context) {

	auth := XPNKReadHeader(c)
	fmt.Printf("\n==========\nauth: %v \n", auth)
	if auth != 1 {
		c.JSON(422, gin.H{"error": "Invalid or missing xapnik token."})
	} else {
	
		var this_header http.Header
		this_header = c.Request.Header
		xpnk_id_object := this_header["Xpnkid"]
		var xpnk_id string
		if len(xpnk_id_object) == 0 {
			c.JSON(422, gin.H{"error" : "Missing or zero-value xpnkid in header."})
		} else {
			xpnk_id = xpnk_id_object[0]
		}	
	
		var disqusauth NewDisqusAuth
		c.Bind(&disqusauth)
		
		disqusauth.Xpnk_id = xpnk_id
	
		if disqusauth.Disqus_accesstoken != "" {
			 c.JSON(201, "Disqus access token received.")
			 updateDisqus(disqusauth)
		} else {
			c.JSON(422, gin.H{"error": "No access token was sent."})
		}	
	}		 	
}


/**********************************************
*
*DATABASE INSERT FUNCTIONS
*
**********************************************/

func updateSlackUser(new_Slackauth NewSlackAuth) int{
	slack_id := new_Slackauth.Slack_userid
	
	dbmap := db_connect.InitDb()
	defer dbmap.Db.Close()
	dbmap.AddTableWithName(NewSlackAuthInsert{}, "USERS").SetKeys(true, "Xpnk_id")
	
	var user_xpnkid int
	err := dbmap.SelectOne(&user_xpnkid, "SELECT user_ID FROM USERS WHERE slack_userid=?", slack_id)
	
	fmt.Printf("\n==========\nUSER_XPNKID: \n%+v\n",user_xpnkid)
	
	if err == nil {
		var new_Slackauthinsert NewSlackAuthInsert
		new_Slackauthinsert.Xpnk_id = user_xpnkid
		new_Slackauthinsert.Slack_accesstoken = new_Slackauth.Slack_accesstoken
	 	new_Slackauthinsert.Slack_userid = new_Slackauth.Slack_userid	 
	 	new_Slackauthinsert.Slack_username = new_Slackauth.Slack_username
		
		_, dberr := dbmap.Update(&new_Slackauthinsert)
		if dberr == nil {
			fmt.Printf("\n==========\nNewSlackAuth Update Success!")
		} else {
			fmt.Printf("\n==========\nProblemz with update: \n%+v\n",dberr)
		}
	} else {
		fmt.Printf("\n==========\nProblemz with select: \n%+v\n",err)
	}
	
	return user_xpnkid
} 

func getXPNKUser(slackuserid string) int {
	dbmap := db_connect.InitDb()
	defer dbmap.Db.Close()
	slackid := slackuserid
	var user_xpnkid int
	err := dbmap.SelectOne(&user_xpnkid, "SELECT user_ID FROM USERS WHERE slack_userid=?", slackid)
	var xpnkid int
	if err == nil {
		xpnkid = user_xpnkid
	} else {
		fmt.Printf("\n==========\nProblemz with selecting user_xpnkid: \n%+v\n",err)
	}
	
	return xpnkid
}

func updateTwitter(new_TwitterAuth NewTwitterAuth) string{
	
	dbmap := db_connect.InitDb()
	defer dbmap.Db.Close()
	dbmap.AddTableWithName(NewTwitterAuthInsert{}, "USERS").SetKeys(true, "Xpnk_id")
	
	var user_xpnkid string
	err := dbmap.SelectOne(&user_xpnkid, "SELECT user_ID FROM USERS WHERE user_ID=?", new_TwitterAuth.Xpnk_id)
	fmt.Printf("\n==========\nXPNK_ID: %+v", user_xpnkid)
	
	if err == nil && user_xpnkid == new_TwitterAuth.Xpnk_id {
		var new_Twttrauthinsert NewTwitterAuthInsert
		new_Twttrauthinsert.Xpnk_id = user_xpnkid
		new_Twttrauthinsert.Twttr_accesstoken = new_TwitterAuth.Twttr_accesstoken
	 	new_Twttrauthinsert.Twttr_secret = new_TwitterAuth.Twttr_secret	 
	 	new_Twttrauthinsert.Twttr_userid = new_TwitterAuth.Twttr_userid
		
		_, dberr := dbmap.Update(&new_Twttrauthinsert)
		if dberr == nil {
			fmt.Printf("\n==========\nNewTwitterAuth Update Success!")
		} else {
			fmt.Printf("\n==========\nProblemz with update: \n%+v\n",dberr)
		}
	} else {
		fmt.Printf("\n==========\nProblemz with select: \n%+v\n",err)
	}
	
	return "updated"
} 

func updateIG(new_IGauth NewIGAuth) string{
	
	dbmap := db_connect.InitDb()
	defer dbmap.Db.Close()
	dbmap.AddTableWithName(NewIGAuthInsert{}, "USERS").SetKeys(true, "Xpnk_id")
	
	var user_xpnkid string
	err := dbmap.SelectOne(&user_xpnkid, "SELECT user_ID FROM USERS WHERE user_ID=?", new_IGauth.Xpnk_id)
	fmt.Printf("\n==========\nXPNK_ID: %+v", user_xpnkid)
	
	if err == nil && user_xpnkid == new_IGauth.Xpnk_id {
		var new_IGauthinsert NewIGAuthInsert
		new_IGauthinsert.Xpnk_id = user_xpnkid
		new_IGauthinsert.Insta_accesstoken = new_IGauth.Insta_accesstoken
	 	new_IGauthinsert.Insta_userid = new_IGauth.Insta_userid	 
	 	new_IGauthinsert.Insta_username = new_IGauth.Insta_username
		
		_, dberr := dbmap.Update(&new_IGauthinsert)
		if dberr == nil {
			fmt.Printf("\n==========\nNewIGAuth Update Success!")
		} else {
			fmt.Printf("\n==========\nProblemz with update: \n%+v\n",dberr)
		}
	} else {
		fmt.Printf("\n==========\nProblemz with select: \n%+v\n",err)
	}
	
	return "updated"
} 

func updateDisqus(new_Disqusauth NewDisqusAuth) string{
	
	dbmap := db_connect.InitDb()
	defer dbmap.Db.Close()
	dbmap.AddTableWithName(NewDisqusAuthInsert{}, "USERS").SetKeys(true, "Xpnk_id")
	
	var user_xpnkid string
	err := dbmap.SelectOne(&user_xpnkid, "SELECT user_ID FROM USERS WHERE user_ID=?", new_Disqusauth.Xpnk_id)
	
	if err == nil && user_xpnkid == new_Disqusauth.Xpnk_id {
		var new_Disqusauthinsert NewDisqusAuthInsert
		new_Disqusauthinsert.Xpnk_id = user_xpnkid
		new_Disqusauthinsert.Disqus_accesstoken = new_Disqusauth.Disqus_accesstoken
	 	new_Disqusauthinsert.Disqus_userid = new_Disqusauth.Disqus_userid	 
	 	new_Disqusauthinsert.Disqus_username = new_Disqusauth.Disqus_username
		
		_, dberr := dbmap.Update(&new_Disqusauthinsert)
		if dberr == nil {
			fmt.Printf("\n==========\nNewDisqusAuth Update Success!")
		} else {
			fmt.Printf("\n==========\nProblemz with update: \n%+v\n",dberr)
		}
	} else {
		fmt.Printf("\n==========\nProblemz with select: \n%+v\n",err)
	}
	
	return "updated"
} 

//end db insert routine	
//Maps are not inherently safe for concurrency - will have to use sync.RWMutex	
