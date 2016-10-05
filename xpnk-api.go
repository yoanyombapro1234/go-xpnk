package main

import (
         "github.com/gin-gonic/gin"
         "fmt"
         "log"
         "database/sql"
   		 _ "github.com/go-sql-driver/mysql"
   		 "github.com/gopkg.in/gorp.v1"
 )
  
type NewIGAuth struct {
	 Insta_accesstoken	string					`json:"access_token"`
	 Insta_userid		string					`json:"insta_userid"`
	 Insta_username		string					`json:"insta_username"`
}

type NewIGAuthInsert struct {
	 Insta_accesstoken	string					`db:"insta_accesstoken"`
	 Insta_userid		string					`db:"insta_userid"`
	 Insta_username		string					`db:"insta_user"`
	 Xpnk_id			int						`db:"user_ID"`
}

type NewDisqusAuth struct {
	 Disqus_accesstoken	string					`json:"access_token"`
	 Disqus_userid		string					`json:"disqus_userid"`
	 Disqus_username	string					`json:"disqus_username"`
}

type NewDisqusAuthInsert struct {
	 Disqus_accesstoken	string					`db:"disqus_accesstoken"`
	 Disqus_userid		string					`db:"disqus_userid"`
	 Disqus_username	string					`db:"disqus_username"`
	 Xpnk_id			int						`db:"user_ID"`
}

func main() {
         
	r := gin.Default()
	r.Use(Cors())
	
/*****************************************
*
* These are the endpoints. The functions 
* they use are in the next section.
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
			
			/*
			V1.OPTIONS ("/slack_new_group", func(c *gin.Context) {
			    c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, PUT")
 				c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
 				c.Next()
			})
			V1.POST ("slack_new_group", SlackNewGroup)
			*/
			
			/*
			V1.OPTIONS ("/slack_new_member", func(c *gin.Context) {
			    c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, PUT")
 				c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
 				c.Next()
			})
			V1.POST ("slack_new_member", SlackNewMember)
			*/
			
			v1.OPTIONS ("/ig_auth", func(c *gin.Context) {
				c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, PUT")
 				c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
 				c.Next()
			})
			v1.POST("/ig_auth", PostIGAuth)
			
			v1.OPTIONS ("/disqus_auth", func(c *gin.Context) {
				c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, PUT")
 				c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
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
* Functions used by endpoints
* 
*****************************************/

/*
func SlackNewMember (c *gin.Context) {
	var slackuser SlackNewUser
	c.Bind(&slackuser)
	
	if slackuser.

}
*/

func PostIGAuth(c *gin.Context) {
	
	var igauth NewIGAuth
	c.Bind(&igauth)
	
	fmt.Printf("\n==========\nIGAUTH\n%+v\n",igauth)
	fmt.Printf("\n==========\nUSERNAME: %+v", igauth.Insta_username)
	fmt.Printf("\n==========\nUSERID: %+v", igauth.Insta_userid)
	fmt.Printf("\n==========\nACCESS TOKEN: %+v", igauth.Insta_accesstoken)
	
	//content := igauth
	
	//c.JSON(201, content)

	
	if igauth.Insta_accesstoken != "" {
	
		 content := &NewIGAuthInsert{
		 	 Insta_accesstoken: igauth.Insta_accesstoken,
		 	 Insta_userid: igauth.Insta_userid,
		 	 Insta_username: igauth.Insta_username,
		 }
		 
		 fmt.Printf("\n==========\nNewIGAuthInsert: \n%+v\n",content)	
		 
		 c.JSON(201, content)
		 updateIG(igauth)
	} else {
		c.JSON(422, gin.H{"error": "No access token was sent."})
	}		 	
}

func PostDisqusAuth(c *gin.Context) {
	
	var disqusauth NewDisqusAuth
	c.Bind(&disqusauth)
	
	fmt.Printf("\n==========\nDISQUSAUTH\n%+v\n",disqusauth)
	
	//content := disqusauth
	
	//c.JSON(201, content)

	
	if disqusauth.Disqus_accesstoken != "" {
	
		 content := &NewDisqusAuthInsert{
		 	 Disqus_accesstoken: disqusauth.Disqus_accesstoken,
		 	 Disqus_userid: disqusauth.Disqus_userid,
		 	 Disqus_username: disqusauth.Disqus_username,
		 }
		 
		 fmt.Printf("\n==========\nNewDisqusAuthInsert: \n%+v\n",content)	
		 
		 c.JSON(201, content)
		 updateDisqus(disqusauth)
	} else {
		c.JSON(422, gin.H{"error": "No access token was sent."})
	}		 	
}

/***************************
*DATABASE INSERT FUNCTIONS
***************************/

func updateIG(new_IGauth NewIGAuth) string{
	ig_username := new_IGauth.Insta_username
	
	dbmap := initDb()
	defer dbmap.Db.Close()
	dbmap.AddTableWithName(NewIGAuthInsert{}, "USERS").SetKeys(true, "Xpnk_id")
	
	var user_xpnkid int
	err := dbmap.SelectOne(&user_xpnkid, "SELECT user_ID FROM USERS WHERE insta_user=?", ig_username)
	fmt.Printf("\n==========\nXPNK_ID: %+v", user_xpnkid)
	
	if err == nil {
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
	disqus_name := new_Disqusauth.Disqus_username
	
	dbmap := initDb()
	defer dbmap.Db.Close()
	dbmap.AddTableWithName(NewDisqusAuthInsert{}, "USERS").SetKeys(true, "Xpnk_id")
	
	var user_xpnkid int
	err := dbmap.SelectOne(&user_xpnkid, "SELECT user_ID FROM USERS WHERE disqus_username=?", disqus_name)
	
	if err == nil {
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


/***************************
* db connection config
***************************/

func checkErr(err error, msg string) {
    if err != nil {
        log.Fatalln(msg, err)
	}
}	
	
func initDb() *gorp.DbMap {
db, err := sql.Open("mysql",
	"")
checkErr(err, "sql.Open failed")

dbmap := &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{}}

return dbmap
}