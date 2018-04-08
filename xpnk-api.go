package main

import (
         "github.com/gin-gonic/gin"
         "fmt"
         "net/http"
         "encoding/json"
   		 _ "github.com/go-sql-driver/mysql"
   		 "database/sql"
   		 //"io/ioutil"
   		 "strings"
   		 "strconv"
   		 "xpnk_constants"
   		 "xpnk_auth"
   		 "xpnk-user/xpnk_checkUserInvite"
   		 "xpnk-user/xpnk_createUserInsert"
   		 "xpnk-user/xpnk_updateUser"
   		 "xpnk-user/xpnk_insertMultiUsers"
   		 "xpnk-shared/db_connect"
   		 "xpnk-group/xpnk_createGroupFromSlack"
   		 "xpnk_slack"
 )
 
type SlackTeamToken struct {
	TeamToken			string		`form:"team_token" binding:"required"`
	BotToken			string		`form:"bot_token"  binding:"required"`
} 
 
type Usertoken struct {
	Token				string		`json:"token"`
} 

type Slack_ID struct {
	Slack_id			string		`json:"slack_id"`
}
 
type NewSlackAuth struct {
	 Slack_accesstoken	string		`json:"access_token"`
	 Slack_userid		string		`json:"slack_userid"`
	 Slack_username		string		`json:"slack_name"`
	 Slack_avatar		string		`json:"slack_avatar"`
}

type NewSlackAuthInsert struct {
	 Slack_accesstoken	string		`db:"slack_authtoken"`
	 Slack_userid		string		`db:"slack_userid"`
	 Slack_username		string		`db:"slack_name"`
	 Slack_avatar		string		`db:"profile_image"`
	 Xpnk_id			int			`db:"user_ID"`
}

type NewUserInvite 		struct {
	Xpnk_token			string		`form:"xpnk_token" binding:"required"`
	Group_name			string		`form:"xpnk_group_name" binding:"required"`
}

type NewGroupMember		struct {
	Group_ID			int			`json:"id"`
	User_ID				int			`json:"userId"`				
}

type NewGroupMemberInsert	struct {
	Group_ID			int			`db:"Group_ID"`
	User_ID				int			`db:"user_ID"`				
}

type XPNKUser 			struct {
	User_ID				int			   `db:"user_ID"			json:"user_ID"`
	Slack_userid		string		   `db:"slack_userid"		json:"slack_userid"`
	Slack_name			string		   `db:"slack_name"			json:"slack_name"`
	Twitter_user		string		   `db:"twitter_user"		json:"twitter_user"`
	Twitter_ID			string		   `db:"twitter_ID"			json:"twitter_ID"`
	Twitter_token		string		   `db:"twitter_authtoken"	json:"twitter_authtoken"`
	Twitter_secret		string		   `db:"twitter_secret"		json:"twitter_secret"`
	Insta_user			string		   `db:"insta_user"			json:"insta_user"`
	Insta_userid		string		   `db:"insta_userid"		json:"insta_userid"`
	Insta_token			string		   `db:"insta_accesstoken"	json:"insta_accesstoken"`
	Disqus_username		sql.NullString `db:"disqus_username"	json:"disqus_username"`
	Disqus_userid		sql.NullString `db:"disqus_userid"		json:"disqus_userid"`
	Disqus_token		string		   `db:"disqus_accesstoken"	json:"disqus_accesstoken"`
	Profile_image		string		   `db:"profile_image"		json:"profile_image"`
}

type TwitterID struct {
	 Twttr_userid		string			`form:"id" binding:"required"`
}

type NewTwitterAuth struct {
	 Twttr_userid		string			`json:"twitter_userid"`
}

/*
type NewTwitterAuthInsert struct {
	 Twttr_accesstoken	string			`db:"twitter_authtoken"`
	 Twttr_secret		string			`db:"twitter_secret"`
	 Twttr_userid		string			`db:"twitter_ID"`
	 Twttr_username		string			`db:"twitter_user"`
	 Xpnk_id			string			`db:"user_ID"`
}
*/

type IGID struct {
	 IG_userid			string			`form:"id" binding:"required"`
}
/*
type NewIGAuthInsert struct {
	 Insta_accesstoken	string			`db:"insta_accesstoken"`
	 Insta_userid		string			`db:"insta_userid"`
	 Insta_username		string			`db:"insta_user"`
	 Xpnk_id			string			`db:"user_ID"`
}
*/

type NewDisqusAuth struct {
	 Disqus_accesstoken	string			`json:"access_token"`
	 Disqus_userid		string			`json:"disqus_userid"`
	 Disqus_username	string			`json:"disqus_username"`
	 Xpnk_id			string
}

type NewDisqusAuthInsert struct {
	 Disqus_accesstoken	string			`db:"disqus_accesstoken"`
	 Disqus_userid		string			`db:"disqus_userid"`
	 Disqus_username	string			`db:"disqus_username"`
	 Xpnk_id			string			`db:"user_ID"`
}

const (
	mySigningKey = ""
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
      
    r.POST("/ping", func(c *gin.Context) {
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
				c.String(200, "pong")
			})
						
			v1.OPTIONS ("/slack_new_group", func(c *gin.Context) {
			    c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, PUT")
 				c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, token, xpnkid")
 				c.Next()
			})
			v1.POST ("/slack_new_group", SlackCreateNewGroup)
			
			v1.OPTIONS ("/slack_response", func(c *gin.Context) {
			    c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, PUT")
 				c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, token, xpnkid")
 				c.Next()
			})
			v1.GET ("/slack_response", SlackResponseHandler)
			v1.POST ("/slack_response", SlackResponseHandler)
			
			
			v1.OPTIONS ("/slack_response/command", func(c *gin.Context) {
			    c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, PUT")
 				c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, token, xpnkid")
 				c.Next()
			})
			v1.GET ("/slack_response/command", SlackCommandHandler)
			v1.POST ("/slack_response/command", SlackCommandHandler)
			
			
			v1.OPTIONS ("/check_user_invite", func(c *gin.Context) {
			    c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, PUT")
 				c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, token, xpnkid")
 				c.Next()
			})
			v1.GET ("/check_user_invite", CheckUserInvite)
			
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
			
			v1.OPTIONS ("/slack_new_member", func(c *gin.Context) {
			    c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, PUT")
 				c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, token, xpnkid")
 				c.Next()
			})
			v1.POST ("slack_new_member", SlackNewMember)
			
			v1.OPTIONS ("/users", func(c *gin.Context) {
				c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT")
 				c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, id, xpnkid")
 				c.Next()
			})
			//v1.GET("/users/:id", UsersXPNKID)
			
			
			v1.OPTIONS ("/users/new", func(c *gin.Context) {
				c.Writer.Header().Set("Access-Control-Allow-Methods", "POST")
 				c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, id, xpnkid, token")
 				c.Next()
			})
			v1.POST("/users/new", UsersNew)
			
			v1.OPTIONS ("/users/update", func(c *gin.Context) {
				c.Writer.Header().Set("Access-Control-Allow-Methods", "POST")
 				c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, id, xpnkid, token")
 				c.Next()
			})
			v1.POST("/users/update", UsersUpdate)
			
			v1.OPTIONS ("/users/twitter", func(c *gin.Context) {
				c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT")
 				c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, id, xpnkid, token")
 				c.Next()
			})
			v1.GET("/users/twitter", UsersByTwitterID)
						
			v1.OPTIONS ("/twitter_auth", func(c *gin.Context) {
				c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, PUT")
 				c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, twitter_userid, xpnkid")
 				c.Next()
			})
			v1.POST("/twitter_auth", PostTwttrAuth)
			
			v1.OPTIONS ("/users/ig", func(c *gin.Context) {
				c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT")
 				c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, id, xpnkid, token")
 				c.Next()
			})
			v1.GET("/users/ig", UsersByIGID)
									
			v1.OPTIONS ("/disqus_auth", func(c *gin.Context) {
				c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, PUT")
 				c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, token, xpnkid")
 				c.Next()
			})
			v1.POST("/disqus_auth", PostDisqusAuth)
/*			
			v1.OPTIONS ("/groups", func(c *gin.Context) {
				c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT")
 				c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, id, userId, xpnkid,")
 				c.Next()
			})
*/			
			v1.OPTIONS ("/groups/members/:id", func(c *gin.Context) {
				c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, PUT")
 				c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, token, xpnkid")
 				c.Next()
			})
			v1.GET ("/groups/members/:id", GroupsByID)
			
			//v1.GET("/groups/:userId", GroupsByUser)
			//v1.GET("/groups/:name", GroupsByGroupName)
			
			v1.OPTIONS ("/groups/add", func(c *gin.Context) {
				c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT")
 				c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, id, userId, xpnkid, token")
 				c.Next()
			})
			v1.POST("/groups/add", GroupsAddMember)
			
			v1.OPTIONS ("/groups/id/:name", func(c *gin.Context) {
				c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT")
 				c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, token, xpnkid")
 				c.Next()
			})
			v1.GET ("/groups/id/:name", GroupID)
			
			//v1.PUT("/groups/updateOwner/:id/:ownerId", GroupsUpdateOwner)
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

func SlackCreateNewGroup (c *gin.Context) {
	fmt.Printf("GIN CONTEXT:  %+v", c)
	var team_tokens xpnk_createGroupFromSlack.SlackTeamTokens
	var token string
	var bot_token string
	c.Bind(&team_tokens) 
	fmt.Printf("TEAM TOKENS: %+v", team_tokens)
	fmt.Printf("TOKEN ONLY:  %+v", team_tokens.TeamToken)
	fmt.Printf("TEST MODE: %+v", team_tokens.TestToken)
	
	token = team_tokens.TeamToken
	bot_token = team_tokens.BotToken

	if token != "" && bot_token != "" { 
	c.JSON(200, "Well hello, Slack friend! I'll set up your Xapnik group and send the team invitations when it's ready.")
		xpnk_createGroupFromSlack.CreateGroup(team_tokens)
	} else {
		c.JSON(422, gin.H{"error":"I'm sorry, I seem to have lost my mind."})
	}
}

func SlackResponseHandler (c *gin.Context) {
    slack_response := c.PostForm("payload")
    var slack_json map[string]interface{}
    json.Unmarshal([]byte(slack_response), &slack_json)
    var token string
    var group_id string
    var callback_id string
    for key, value := range slack_json {
    	switch key {
    	    case "callback_id":
                callback_id = value.(string)   
    	}
    }	
    if callback_id != "" {
    	s := strings.Split(callback_id, ",")
    	if s[0] == "invites" {
    		group_id = s[1]	
    		token = s[2]
    	} else {
    		group_id = ""
    	}
    }
    if group_id != "" && token != "" {
        SlackInviteGroup(token, group_id)
    }
    
    fmt.Printf("PAYLOAD: \n %s \n", string(slack_response))
    fmt.Printf("group_id: \n %s \n", token)
    fmt.Printf("slacker_token: \n %s \n", group_id)
    //fmt.Printf("SLACK RESPONSE: \n %s \n", string(slack_response))
    c.JSON(200, "Thanks!")
}

func SlackInviteGroup (token string, group_id string) {
        groupID, err := strconv.Atoi(group_id)
        if err != nil {
                fmt.Printf("Couldn't convert group_id to string in api line 399.")
        } else {
            fmt.Printf("I'm inviting the group!")
            xpnk_createGroupFromSlack.InviteGroup (token, groupID)
        }
}

func SlackCommandHandler (c *gin.Context) {
    fmt.Printf("GIN CONTEXT:  %+v", c)
	var command_body xpnk_slack.SlackCommand
	var token string
	c.Bind(&command_body) 
	fmt.Printf("COMMAND BODY: %+v", command_body)
	
	token = command_body.Token

	if token != "" && token == xpnk_constants.SlackCommandTkn { 
		var response string
		response = xpnk_slack.SlackGroupStatus(command_body)
		fmt.Printf("\nAPI Response: %+v\n", response)
		c.JSON(200, response)
	}
}	

func CheckUserInvite (c *gin.Context) {
	var user_invite			NewUserInvite
	var user_invite_check	xpnk_checkUserInvite.GroupObj
	c.Bind(&user_invite)
	fmt.Printf("CheckUserInvite Xpnk_token:  %v \n", user_invite.Xpnk_token)
	fmt.Printf("CheckUserInvite Group_name:  %v \n", user_invite.Group_name)
	
	user_invite_check		= xpnk_checkUserInvite.CheckUserInvite(user_invite.Xpnk_token, user_invite.Group_name)
	if user_invite_check.GroupName == user_invite.Group_name {
		c.JSON(201, user_invite_check)
	} else {
		c.JSON(422, gin.H{"error": user_invite_check })
	}
}

func GroupsByID (c *gin.Context) {
	var groupid 	string
	groupid 		= c.Param("id")
	
	if groupid != "" {
		groupmems 	:= getGroup(groupid)	
		fmt.Printf("\nGROUPMEMS RETURNED BY GetGROUP: %+v\n", groupmems)	 
		c.JSON(201, groupmems) 
	} else {
		c.JSON(422, gin.H{"error": "No group_id was sent."})
	}		 
}

func GroupID (c *gin.Context) {
	var groupname 	string
	groupname 		= c.Param("name")
	
	if groupname != "" {
		groupid 	:= getGroupID(groupname)	
		fmt.Printf("\nGROUPID RETURNED BY GetGroupID: %+v\n", groupid)	 
		c.JSON(201, groupid) 
	} else {
		c.JSON(422, gin.H{"error": "No groupname was sent."})
	}		 
}

func GroupsAddMember (c *gin.Context) {
	var new_groupMember		NewGroupMember
	c.Bind(&new_groupMember)
	fmt.Printf("\n new_groupMember.Group_ID:  %v \n", new_groupMember.Group_ID)
	fmt.Printf("new_groupMember.User_ID:  %v \n", new_groupMember.User_ID)
	
	var new_member_insert   NewGroupMemberInsert
	new_member_insert.Group_ID	= new_groupMember.Group_ID
	new_member_insert.User_ID	= new_groupMember.User_ID
	
	insert_new_member 			:= InsertNewGroupMember(new_member_insert)
	
	if insert_new_member == 1 {
		c.JSON(201, "User added!")
	} else {
		c.JSON(422, gin.H{"error": insert_new_member })
	}
}

func XPNKAuthSet (c *gin.Context)  {
	usertoken, err := xpnk_auth.NewToken([]byte(mySigningKey), "", "")
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
	fmt.Printf("HEADER: %+v  END\n", c.Request.Header)
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

func UsersNew (c *gin.Context) {
	var newUser					xpnk_createUserInsert.User_Insert
	c.Bind(&newUser)
	fmt.Printf("newUser to add:  %+v \n", newUser)
	var userInsert				[]xpnk_createUserInsert.User_Insert
	userInsert 				 =  append(userInsert, newUser)
	
	insertUser 				:=  xpnk_insertMultiUsers.InsertMultiUsers(userInsert)
	if insertUser == "inserted" {
		c.JSON(200, "New user inserted.")
	}	else {
		c.JSON(202, "New user was not inserted. Check the API logs.")
	}		
}

func UsersUpdate (c *gin.Context) {
	var thisUser				xpnk_updateUser.User_Update
	c.BindJSON(&thisUser)
	fmt.Printf("thisUser to add:  %+v \n", thisUser)
	update_thisUser 			:=  xpnk_updateUser.UpdateUser(thisUser)
	if update_thisUser == 1 {
		c.JSON(200, "User updated.")
	}	else {
		c.JSON(202, "User not updated. Check the API logs.")
	}	
}

func UsersByTwitterID (c *gin.Context) {
	var twitterId				TwitterID
	var user					XPNKUser 			
	c.Bind(&twitterId)
	twitter_id					:= twitterId.Twttr_userid
	fmt.Printf("twitter_id:  %v \n", twitter_id)
	if twitter_id == "" {
		c.JSON(422, gin.H{"error": "Invalid or missing Twitter user ID."})
	} else {
		user 					= get_user_by_twitter(twitter_id)
		if user.Twitter_ID != twitter_id {
			c.JSON(202, user)
		} else {
			c.JSON(200, user)
		}
	}
}

func UsersByIGID (c *gin.Context) {
	var igID					IGID
	var user					XPNKUser
	c.Bind(&igID)
	ig_id						:= igID.IG_userid
	fmt.Printf("ig_id: %v \n", ig_id)
	if ig_id == "" {
		c.JSON(422, gin.H{"error": "Invalid or missing Twitter user ID."})
	} else {
		user 					= get_user_by_ig(ig_id)
		if user.Insta_userid != ig_id {
			c.JSON(202, user)
		} else {
			c.JSON(200, user)
		}
	}
}

func PostTwttrAuth(c *gin.Context) {
	var twitter_user			NewTwitterAuth
	c.Bind(&twitter_user)
	twitterId					:= twitter_user.Twttr_userid
	fmt.Printf("twitterId:  %v \n", twitterId)
	if twitterId == ""{
		c.JSON(422, gin.H{"error": "Invalid or missing Twitter user ID."})
	} else {
	
		twitter_user_check		:= check_twitter_id(twitterId)
		if twitter_user_check != 0 {
			fmt.Printf("\nXPNKUSERID RETURNED BY check_twitter_id: %+v\n", twitter_user_check)
			c.JSON(201, twitter_user_check)
		} else {
			c.JSON(422, gin.H{"error": twitter_user_check })
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
*DATABASE FUNCTIONS
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
	 	new_Slackauthinsert.Slack_avatar = new_Slackauth.Slack_avatar
		
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

func get_user_by_twitter(twitter_id string) XPNKUser {
	dbmap := db_connect.InitDb()
	defer dbmap.Db.Close()
	
	var xpnkUser			XPNKUser
	twitterId				:= twitter_id
	err	:= dbmap.SelectOne(&xpnkUser, "SELECT * FROM USERS WHERE twitter_ID=?", twitterId)
	if err != nil {
		fmt.Printf("\n==========\nget_user_by_twitter - Problemz with selecting user by twitterID: \n%+v\n",err)
	} else {
		fmt.Printf("\n==========\nfound user: \n%+v\n",xpnkUser.User_ID)
	}
	return xpnkUser
}

func get_user_by_ig(ig_id string) XPNKUser {
	dbmap := db_connect.InitDb()
	defer dbmap.Db.Close()
	
	var xpnkUser			XPNKUser
	IGId					:= ig_id
	err	:= dbmap.SelectOne(&xpnkUser, "SELECT * FROM USERS WHERE insta_userid=?", IGId)
	if err != nil {
		fmt.Printf("\n==========\nget_user_by_ig - Problemz with selecting user by IGId: \n%+v\n",err)
	} else {
		fmt.Printf("\n==========\nfound user: \n%+v\n",xpnkUser.User_ID)
	}
	return xpnkUser
}

func check_twitter_id(twitter_id string) int {
	dbmap := db_connect.InitDb()
	defer dbmap.Db.Close()
	twitterId				:= twitter_id
	var user_xpnkid 		int
	var xpnkid				int
	err	:= dbmap.SelectOne(&user_xpnkid, "SELECT user_ID FROM USERS WHERE twitter_ID=?", twitterId)
	if err == nil {
		xpnkid 				= user_xpnkid
	} else {
		fmt.Printf("\n==========\ncheck_twitter_id - Problemz with selecting user_xpnkid: \n%+v\n",err)
		xpnkid = 0
	}
	return xpnkid
}

/*
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
	 	new_Twttrauthinsert.Twttr_username = new_TwitterAuth.Twttr_username
		
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
*/

func check_ig_id (ig_id string) int {
	dbmap := db_connect.InitDb()
	defer dbmap.Db.Close()
	IGId					:= ig_id
	var user_xpnkid 		int
	var xpnkid				int
	err	:= dbmap.SelectOne(&user_xpnkid, "SELECT user_ID FROM USERS WHERE insta_userid	=?", IGId)
	if err == nil {
		xpnkid 				= user_xpnkid
	} else {
		fmt.Printf("\n==========\ncheck_ig_id - Problemz with selecting user_xpnkid: \n%+v\n",err)
		xpnkid = 0
	}
	return xpnkid
}

func check_user_ig (xpnk_id int) string {
	dbmap := db_connect.InitDb()
	defer dbmap.Db.Close()
	xpnkid					:= xpnk_id
	var user_ig				string
	var check_user_ig		string
	err	:= dbmap.SelectOne(&user_ig, "SELECT insta_userid FROM USERS WHERE user_ID	=?", xpnkid)
	if err == nil {
		check_user_ig 		= user_ig
	} else {
		fmt.Printf("\n==========\ncheck_user_ig - Problemz with checking user's Instagram ID: \n%+v\n",err)
		check_user_ig = "0"
	}
	return check_user_ig
} 

/*
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
*/
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

func InsertNewGroupMember(new_GroupMember NewGroupMemberInsert) int {
	var returnVal 				int
	dbmap := db_connect.InitDb()
	defer dbmap.Db.Close()
	dbmap.AddTableWithName(NewGroupMemberInsert{}, "USER_GROUPS")
		
	err := dbmap.Insert(&new_GroupMember)	
		if err == nil {
			fmt.Printf("\n==========\nInsertNewGroupMember Update Success!")
			returnVal = 1
		} else {
			fmt.Printf("\n==========\nProblemz with InsertNewGroupMember line 700 in api: \n%+v\n",err)
			returnVal = 0
		}
	return returnVal	
}

func getGroup (groupID string) []int{
	var groupUsers			[]int
	dbmap := db_connect.InitDb()
	defer dbmap.Db.Close()
	
	err, _ := dbmap.Select(&groupUsers, "SELECT user_ID FROM USER_GROUPS WHERE Group_ID=?", groupID)
	
	if err == nil {
		return groupUsers
	} else {
		fmt.Printf("\n==========\n getGroup - Problemz with getting users for group api line 830: \n%+v\n",err)
		return groupUsers
	}
}

func getGroupID (groupName string) int{
	var group_name			= strings.Replace(groupName, "-", " ", -1)
	var groupID				int
	dbmap 					:= db_connect.InitDb()
	defer dbmap.Db.Close()
	
	err := dbmap.SelectOne(&groupID, "SELECT Group_ID FROM GROUPS WHERE group_name=?", group_name)
	
	if err == nil {
		return groupID
	} else {
		fmt.Printf("\n==========\n getGroupID - Problemz with getting Group_ID for group api line 8872: \n%+v\n",err)
		return groupID
	}
}
	
//Maps are not inherently safe for concurrency - will have to use sync.RWMutex	
