package main

import (
         "github.com/gin-gonic/gin"
         "fmt"
         "net/http"
         "encoding/json"
   		 _ "github.com/go-sql-driver/mysql"
   		 "database/sql"
   		 "strings"
   		 "strconv"
   		 "reflect"
   		 "errors"
   		 "xpnk_constants"
   		 "xpnk_auth"
   		 "xpnk-user/xpnk_checkUserInvite"
   		 "xpnk-user/xpnk_createUserInsert" //v.1 to be deprecated
   		 "xpnk-user/xpnk_createUserObject"
   		 "xpnk-user/xpnk_updateUser"
   		 "xpnk-user/xpnk_insertMultiUsers"
   		 "xpnk-user/xpnk_checkTwitterId"
   		 "xpnk-shared/db_connect"
   		 "xpnk-group/xpnk_createGroupFromSlack"
   		 "xpnk-group/xpnk_createGroup"
   		 "xpnk-group/xpnk_createInvite"
   		 "xpnk_slack"
 )
 
type SlackTeamToken struct {
	TeamToken			string					`form:"team_token" binding:"required"`
	BotToken			string					`form:"bot_token"  binding:"required"`
} 
 
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
	 Slack_avatar		string					`json:"slack_avatar"`
}

type NewSlackAuthInsert struct {
	 Slack_accesstoken	string					`db:"slack_authtoken"`
	 Slack_userid		string					`db:"slack_userid"`
	 Slack_username		string					`db:"slack_name"`
	 Slack_avatar		string					`db:"profile_image"`
	 Xpnk_id			int						`db:"user_ID"`
}

type InviteRequest struct {
	Id				int					`form:"id" binding:"required"`
	Source			string				`form:"source"`
}

type NewUserInvite 		struct {
	Xpnk_token			string					`form:"xpnk_token" binding:"required"`
	Group_name			string					`form:"xpnk_group_name" binding:"required"`
}

type NewGroupMember		struct {
	Group_ID			int						`json:"id"`
	User_ID				int						`json:"userId"`				
}

type NewGroupMemberInsert	struct {
	Group_ID			int						`db:"Group_ID"`
	User_ID				int						`db:"user_ID"`				
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

type UserGroups 		struct {
	Xpnk_id				string			`db:"user_ID"			json:"user_ID"`
	Groups				[]GroupsByUser
}

type GroupsByUser		struct {
	Group_ID			int
	Owner				bool
	Admin				bool
	Name				string			
	Slug				string
}

type GroupOwner 		struct {
	Group_ID			int				`db:"Group_ID"			json:"Group_ID"`
	Owner				sql.NullBool	`db:"group_owner"		json:"group_owner"`
	Admin				sql.NullBool	`db:"group_admin"		json:"group_admin"`
	Name				string			`db:"group_name"		json:"group_name"`
	Slug				string									`json:"group_slug"`		
}

type TwitterID struct {
	 Twttr_userid		string					`form:"id" binding:"required"`
}

type NewTwitterAuth struct {
	 Twttr_userid		string					`json:"twitter_userid"`
}

/*
type NewTwitterAuthInsert struct {
	 Twttr_accesstoken	string					`db:"twitter_authtoken"`
	 Twttr_secret		string					`db:"twitter_secret"`
	 Twttr_userid		string					`db:"twitter_ID"`
	 Twttr_username		string					`db:"twitter_user"`
	 Xpnk_id			string					`db:"user_ID"`
}
*/

type IGID struct {
	 IG_userid			string					`form:"id" binding:"required"`
}
/*
type NewIGAuthInsert struct {
	 Insta_accesstoken	string					`db:"insta_accesstoken"`
	 Insta_userid		string					`db:"insta_userid"`
	 Insta_username		string					`db:"insta_user"`
	 Xpnk_id			string					`db:"user_ID"`
}
*/

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
	mySigningKey = ""
)

type Person struct {
	ID        uint   `json:"id"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

	 

func main() {

	r := gin.Default()
	r.Use(Cors())
	r.Static("api/v1/data", "../xpnk-data/")
	r.Static("api/v2/data", "../xpnk-data/")
	
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
      
    v2 := r.Group("api/v2")
		{  
			v2.OPTIONS ("/ping", func(c *gin.Context) {
				c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, PUT")
 				c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
 				c.Next()
			})
			v2.GET ("/ping", func(c *gin.Context) {
				c.String(200, "pong")
			})
			
			v2.OPTIONS ("/users/twitter/:id", func(c *gin.Context) {
				c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT")
 				c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, id, xpnkid, token")
 				c.Next()
			})
			v2.GET("/users/twitter/:id", UsersByTwitterID_2)
			
			v2.OPTIONS ("/users/ig/:id", func(c *gin.Context) {
				c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT")
 				c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, id, xpnkid, token")
 				c.Next()
			})
			v2.GET("/users/ig/:id", UsersByIGID_2)
			
			v2.OPTIONS ("/users/invite", func(c *gin.Context) {
			    c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, PUT")
 				c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, token, xpnkid")
 				c.Next()
			})
			v2.GET ("/users/invite", CheckUserInvite)
			
			v2.OPTIONS ("/users/authSet", func(c *gin.Context) {
				c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, PUT")
 				c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, token, xpnkid")
 				c.Next()
			})
			v2.GET ("/users/authSet", XPNKAuthSet)
			
			v2.GET ("/users/groups/:id", GetUserGroups)
			
			v2.GET ("/users/login/twitter", LoginTwitter)
			
			v2.OPTIONS ("/users/authCheck", func(c *gin.Context) {
				c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, PUT")
 				c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, token, xpnkid")
 				c.Next()
			})
			v2.POST ("/users/authCheck", XPNKAuthCheck)
			
			v2.OPTIONS ("/users", func(c *gin.Context) {
				c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, PUT")
 				c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, id, xpnkid, token")
 				c.Next()
			})
			v2.POST("/users", UsersNew_2)
			
			v2.PUT("/users/:id", UsersUpdate_2)
			v2.DELETE("/users/:id", UsersDelete)
			
			v2.OPTIONS ("/groups", func(c *gin.Context) {
				c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, PUT")
 				c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, token, xpnkid")
 				c.Next()
			})
			
			v2.OPTIONS ("/groups/:id/members", func(c *gin.Context) {
				c.Writer.Header().Set("Access-Control-Allow-Methods", "PUT")
 				c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, token, xpnkid")
 				c.Next()
			})
			v2.GET ("/groups/:id/members", GroupsByID)
			v2.POST("/groups/", GroupsNew)
			v2.GET ("/groups/:id/invite/:source", GroupsInvite)
			
			v2.OPTIONS ("/groups/:id/owner/:owner", func(c *gin.Context) {
				c.Writer.Header().Set("Access-Control-Allow-Methods", "PUT, DELETE")
 				c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, token, xpnkid")
 				c.Next()
			})
			v2.DELETE("/groups/:id/owner/:owner", GroupsDelete)
			
			v2.OPTIONS ("groups/:id/user/:user/owner/:owner", func(c *gin.Context) {
				c.Writer.Header().Set("Access-Control-Allow-Methods", "PUT, DELETE")
 				c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, token, xpnkid")
 				c.Next()
			})
			v2.DELETE("groups/:id/user/:user/owner/:owner", GroupsMemberDelete)
			
		}

/*****************************************
* V1
*****************************************/
	
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

/*****************************************
* V2
*****************************************/

func UsersByTwitterID_2 (c *gin.Context) {
	twitter_id					:= c.Param("id")
	var user					XPNKUser 	
	var err_msg					error		
	fmt.Printf("twitter_id:  %v \n", twitter_id)
	if twitter_id == "" {
		c.JSON(422, gin.H{"error": "Invalid or missing Twitter user ID."})
	} else {
		user, err_msg 			= get_user_by_twitter(twitter_id)
		if user.Twitter_ID != twitter_id {
			c.JSON(400, err_msg.Error())
		} else {
			c.JSON(200, user)
		}
	}
}

func UsersByIGID_2 (c *gin.Context) {
	ig_ID					 := c.Param("id")
	var user					XPNKUser
	var err_msg					error
	if ig_ID == "" {
		c.JSON(422, gin.H{"error": "Invalid or missing Twitter user ID."})
	} else {
		user, err_msg 		  = get_user_by_ig(ig_ID)
		if user.Insta_userid != ig_ID {
			c.JSON(400, err_msg.Error())
		} else {
			c.JSON(200, user)
		}
	}
}

func GetUserGroups (c *gin.Context) { 
	var groups					[]GroupOwner
	var group_trim				GroupsByUser
	var	groups_trim				[]GroupsByUser
	var user_groups				UserGroups
	var err_msg					error
	user_id						:= c.Params.ByName("id")
	
	if user_id == "" {
		c.JSON(422, gin.H{"error": "Invalid or missing user ID."})
		return
	} else {
		groups, err_msg 		  = get_user_groups(user_id)
		if err_msg != nil {
			c.JSON(400, err_msg.Error())
		} else {	
			for i := 0; i < len(groups); i++ {
				var this_group GroupOwner
				this_group = groups[i]
				group_name				:= strings.ToLower(this_group.Name)
				group_path	 			:= strings.Replace(group_name, " ", "-", -1)
				
				group_trim.Group_ID 	= this_group.Group_ID
				group_trim.Owner 		= this_group.Owner.Bool
				group_trim.Admin		= this_group.Admin.Bool
				group_trim.Name			= this_group.Name
				group_trim.Slug 		= group_path
				groups_trim 			= append(groups_trim, group_trim)
			}	
			user_groups.Xpnk_id   = user_id
			user_groups.Groups 	  = groups_trim
			c.JSON(200, user_groups)
		}
	}
} 

func LoginTwitter (c *gin.Context) {
	type TwitterCreds struct {
		Token				string			`form:"token" binding:"required"`
		Secret				string			`form:"secret" binding:"required"`
	}
	var twitter_creds TwitterCreds
	var err error
	c.Bind(&twitter_creds)
	token := twitter_creds.Token
	secret := twitter_creds.Secret
	if twitter_creds.Token == "" || twitter_creds.Secret == "" {
		c.JSON(400, "A user token and secret are required. One or both are missing.")
		return
	}
	
	user_groups, err := xpnk_checkTwitterId.CheckTwitterId(token, secret)
	if err !=  nil {
		c.JSON(400, err.Error())	
	} else {
		c.JSON(200, user_groups)
	}
}

func UsersNew_2 (c *gin.Context) {
	var newUser					xpnk_createUserObject.User_Object
	var err_msg					error
	c.Bind(&newUser)
	fmt.Printf("newUser to add:  %+v \n", newUser)
	if newUser.TwitterID == "" && newUser.InstaUserID == "" {
		c.JSON(400, "Need either a Twitter user ID or a Instagram user ID to create a new user.")
		return
	}
	var userInsert				[]xpnk_createUserObject.User_Object
	userInsert 				 =  append(userInsert, newUser)
	
	newID, err_msg 			:=  xpnk_insertMultiUsers.InsertMultiUsers_2(userInsert)
	if err_msg != nil {
		c.JSON(400, err_msg.Error())	
	} else {
		c.JSON(200, newID)
	}
}

func UsersUpdate_2(c *gin.Context) {

	var thisUser xpnk_createUserObject.User_Object
	var err_msg error
	id, err := strconv.Atoi(c.Params.ByName("id"))
	if err != nil {
		fmt.Printf("\nError:  %+v \n", err) 
		c.JSON(400, err)
		return
	}	
	c.BindJSON(&thisUser)
	if thisUser.InstaUserID == "" && thisUser.TwitterID == "" {
		fmt.Printf("You must send either an InstaUserID or TwitterID param, or both. If you passed an int value for either, please change it to a string.")
		c.JSON(400, "You must send either an InstaUserID or TwitterID param, or both. If you passed an int value for either, please change it to a string.")	
		return
	}
		
	if (reflect.TypeOf(thisUser.InstaUserID).String()) != "string" || (reflect.TypeOf(thisUser.TwitterID).String()) != "string" {
		c.JSON(400, "TwitterID and InstaUserID values must be strings. Please check.")
		return
	}
	
	thisUser.Id = id
	
	update_thisUser, err_msg 	:=  xpnk_updateUser.UpdateUser_2(thisUser)
	if err_msg != nil {
		fmt.Printf(err_msg.Error())
		c.JSON(400, err_msg.Error())
		return
	}
	
	if update_thisUser == 1 {
		fmt.Printf("\nthisUser updated:  %+v \n ID: %n\n", thisUser, id)
		c.JSON(200, thisUser)
	}
}

func UsersDelete (c *gin.Context) {
	userid, err 			:= 	strconv.Atoi(c.Params.ByName("id"))
	if err != nil {
		c.JSON( 400, err.Error())
		return
	}
	
	if userid <= 0 {
	  	c.JSON(422, gin.H{"error": "No User_id was sent."})
	  	return
	} else {
		userdel, err 	:= delUser(userid)
		if err != nil {
			 fmt.Printf("\nERROR DELETING USER: %+v\n", err)
			c.JSON(400, err.Error())
			return
		} else {
			fmt.Printf("\nUSER DELETED: %+v\n", userdel)	
			returnstring := "User deleted: " + c.Params.ByName("id")
			c.JSON(201, returnstring)
		}	
	}		 
}

func GroupsNew (c *gin.Context) {
	var newGroup				xpnk_createGroup.NewGroup
	var err_msg					error
	c.Bind(&newGroup)
	fmt.Printf("newGroup to add:  %+v \n", newGroup)
	if newGroup.GroupName == "" {
		c.JSON(400, "A group name is required to create a new group.")
		return
	}
	
	newID, err_msg 			:=  xpnk_createGroup.CreateGroup(newGroup)
	if err_msg != nil {
		c.JSON(400, err_msg.Error())	
	} else {
		c.JSON(200, newID)
	}
}

func GroupsInvite (c *gin.Context) {
	id, err 				:= strconv.Atoi(c.Params.ByName("id"))
	if err != nil {
		c.JSON( 400, err.Error())
		return
	}
	source		 			:= c.Params.ByName("source")	
	response, err			:= xpnk_createInvite.CreateInvite(id, source, "")
	if err != nil {
		c.JSON( 400, err.Error())
		return
	}
	c.JSON(200, response)
}

func GroupsDelete (c *gin.Context) {
	groupid		 			:= 	c.Params.ByName("id")
	ownerid		 			:= 	c.Params.ByName("owner")
	
	groupnum, err 			:=	strconv.Atoi(groupid)
	_, err2					:=  strconv.Atoi(ownerid)
	
	if err != nil || err2 != nil {
		c.JSON(422, gin.H{"error": "One of the ids you sent is missing or wrong."})
	  	return
	}
		
	if groupnum <= 0 {
	  	c.JSON(422, gin.H{"error": "No group_id was sent."})
	  	return
	} else {
		groupdel, err 	:= delGroup(groupid, ownerid)
		if err != nil {
			 fmt.Printf("\nERROR DELETING GROUP: %+v\n", err)
			c.JSON(400, err.Error())
			return
		} else {
			fmt.Printf("\nGROUP DELETED: %+v\n", groupdel)	
			returnstring := "Group deleted: " + c.Params.ByName("id")
			c.JSON(201, returnstring)
		}	
	}		 
}

func GroupsMemberDelete (c *gin.Context) {
	groupid 			:= 	c.Params.ByName("id")
	userid				:= 	c.Params.ByName("user")
	ownerid				:= 	c.Params.ByName("owner")
	
	groupnum, err		:= strconv.Atoi(groupid)
	usernum, err2		:= strconv.Atoi(userid)
	ownernum, err3		:= strconv.Atoi(ownerid)
	
	if err != nil || err2 != nil || err3 != nil {
		c.JSON(422, gin.H{"error": "One of the ids you sent is missing or wrong."})
	  	return
	}
	
	if groupnum <= 0 || usernum <= 0 || ownernum <= 0 {
	  	c.JSON(422, gin.H{"error": "One of the ids you sent is missing or wrong."})
	  	return
	} else {
		memberdel, err 	:= delMember(groupid, userid, ownerid)
		if err != nil {
			 fmt.Printf("\nERROR DELETING MEMBER: %+v\n", err)
			c.JSON(400, err.Error())
			return
		} else {
			fmt.Printf("\nUSER REMOVED FROM GROUP: %+v\n", memberdel)	
			returnstring := "User removed: " + c.Params.ByName("user")
			c.JSON(201, returnstring)
		}	
	}		 
}

/*****************************************
* V1
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
		c.JSON(422, gin.H{"error":"I'm sorry, I can't see either or both of the token parameters."})
	}
}

func SlackResponseHandler (c *gin.Context) {
    slack_response := c.PostForm("payload")
    var slack_json map[string]interface{}
    json.Unmarshal([]byte(slack_response), &slack_json)
    var token string
    var group_id string
    var callback_id string
    var test_mode string
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
    		test_mode = s[3]
    	} else {
    		group_id = ""
    	}
    }
    if group_id != "" && token != "" {
        SlackInviteGroup(token, group_id, test_mode)
    } 
    
    fmt.Printf("PAYLOAD: \n %s \n", string(slack_response))
    fmt.Printf("group_id: \n %s \n", token)
    fmt.Printf("slacker_token: \n %s \n", group_id)
    //fmt.Printf("SLACK RESPONSE: \n %s \n", string(slack_response))
    c.JSON(200, "Thanks!")
}

func SlackInviteGroup (token string, group_id string, test_mode string) {
        groupID, err := strconv.Atoi(group_id)
        if err != nil {
                fmt.Printf("Couldn't convert group_id to string in api line 399.")
        } else {
        	if test_mode == "false" || test_mode == "" {
				fmt.Printf("I'm inviting the group!")
				xpnk_createGroupFromSlack.InviteGroup (token, groupID, "false")
			} else if test_mode == "true" {
				fmt.Printf("I'm inviting the group!")
				xpnk_createGroupFromSlack.InviteGroup (token, groupID, "true")
			}	
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
		response := xpnk_slack.SlackGroupStatus(command_body)
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
		c.JSON(400, user_invite_check)
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
	var err_msg					error			
	c.Bind(&twitterId)
	twitter_id					:= twitterId.Twttr_userid
	fmt.Printf("twitter_id:  %v \n", twitter_id)
	if twitter_id == "" {
		c.JSON(422, gin.H{"error": "Invalid or missing Twitter user ID."})
	} else {
		user, err_msg 			= get_user_by_twitter(twitter_id)
		if user.Twitter_ID != twitter_id {
			c.JSON(202, user)
			fmt.Printf("error:  %v \n", err_msg)
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
		c.JSON(422, gin.H{"error": "Invalid or missing Twitter user ID."})
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

func getUserByTwitter(twitter_id string) int {
	dbmap := db_connect.InitDb()
	defer dbmap.Db.Close()
	this_twitter_id := twitter_id
	var xpnk_id int
	err := dbmap.SelectOne(&xpnk_id, "SELECT user_ID FROM USERS WHERE twitter_ID=?", this_twitter_id)
	var return_id int
	if err == nil {
		return_id = xpnk_id
	} else {
		fmt.Printf("\n==========\nProblemz with selecting xpnk_id: \n%+v\n",err)
	}
	
	return return_id
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

func get_user_by_twitter(twitter_id string) (XPNKUser, error) {
	dbmap := db_connect.InitDb()
	defer dbmap.Db.Close()
	
	var xpnkUser			XPNKUser
	var err_msg				error
	twitterId				:= twitter_id
	
	err	:= dbmap.SelectOne(&xpnkUser, "SELECT `user_ID`, `twitter_user`, `twitter_ID`, `twitter_authtoken`, `twitter_secret`, `insta_user`, `insta_userid`, `insta_accesstoken`, `disqus_username`, `disqus_userid`, `disqus_accesstoken`, `profile_image` FROM USERS WHERE twitter_ID=?", twitterId)
	if err != nil {
		fmt.Printf("\n==========\nget_user_by_twitter - Problemz with selecting user by twitterID: \n%+v\n",err)
		err_msg = err
		fmt.Printf("\n==========\nget_user_by_twitter - Problemz with selecting user by twitterID: \n%+v\n",err_msg)
	} else {
		fmt.Printf("\n==========\nfound user: \n%+v\n",xpnkUser.User_ID)
	}
	return xpnkUser, err_msg
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

func get_user_groups(user_id string) ([]GroupOwner, error) {
	dbmap := db_connect.InitDb()
	defer dbmap.Db.Close()

	var groupOwners			[]GroupOwner
	var err_msg				error
	
	id						:= user_id
	
	_, err := dbmap.Select(&groupOwners, "SELECT `USER_GROUPS`.`Group_ID`, `USER_GROUPS`.`group_owner`, `USER_GROUPS`.`group_admin`, `GROUPS`.`group_name` FROM USER_GROUPS INNER JOIN GROUPS ON `USER_GROUPS`.`Group_ID` = `GROUPS`.`Group_ID` WHERE `USER_GROUPS`.`user_ID` =?", id)
		
	/*
	SELECT `USER_GROUPS`.`Group_ID`, `USER_GROUPS`.`group_owner`, `USER_GROUPS`.`group_admin`, `groups`.`group_name` FROM USER_GROUPS INNER JOIN groups ON `USER_GROUPS`.`Group_ID` = `groups`.`Group_ID` WHERE `USER_GROUPS`.`user_ID` = 1;
	*/
	
	if err != nil {
		err_msg				= err
	} 
	
	return groupOwners, err_msg
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

func delUser (userID int) (int64, error) {
	type User struct {
		User_ID 			int 			`db:"user_ID"`
	}
	
	var user_id 			User 
	user_id.User_ID = userID 
	fmt.Printf("\n==============\n User_ID to be deleted: %+v", user_id.User_ID)
	
	dbmap 					:= db_connect.InitDb()
	defer dbmap.Db.Close()
	
	dbmap.AddTableWithName(User{}, "USERS").SetKeys(true, "user_ID")
	
	_, err := dbmap.Delete(&user_id)
	fmt.Printf("\n==============\n deleted: %+v", user_id)
	
	count, err := dbmap.SelectInt("select count(*) from USERS where user_ID=?", user_id.User_ID)
	fmt.Printf("\n==============\n COUNT: %+v", count)
	
	res, err2 := delUserGroups(user_id.User_ID)
	if err2 != nil {
		fmt.Printf("\n===========\n delUserGroups error: %+v", err)
	} else {
		fmt.Printf("\n===========\n delUserGroups response: %+v", res)
	}
	
	return count, err
}

func delUserGroups (userID int) (sql.Result, error) {
	dbmap 					:= db_connect.InitDb()
	defer dbmap.Db.Close()
		
	res, err := dbmap.Exec("delete from USER_GROUPS where User_ID=?", userID)
	
	if err != nil {
		fmt.Printf("\n===========\n delUserGroups error: %+v", err)
	} else {
		fmt.Printf("\n===========\n delUserGroups response: %+v", res)
	}
	
	return res, err
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

func groupOwner (groupID string, ownerID string) (bool, error) {
	dbmap 					:= db_connect.InitDb()
	defer dbmap.Db.Close()
	
	query := "select group_owner from USER_GROUPS WHERE Group_ID=" + groupID + " AND user_ID=" + ownerID + " AND group_owner=1"

	var owner_check int
	err := dbmap.SelectOne(&owner_check, query)
	if err != nil {
		owner_error := errors.New("Only group owner can delete group or member. Owner ID passed is not group owner. Did not find groups owner id as owner: " + err.Error())
		return false, owner_error
	}
	if owner_check != 1 {
		fmt.Printf("\n===========\n owner_check: %+v", strconv.Itoa(owner_check))
		err1 := errors.New("Only group owner can delete group member. Owner ID passed is not group owner.")
		return false, err1
	}
	return true, err
}

func delGroup (groupID string, ownerID string) (int64, error) {
	type Group struct {
		Group_ID 			int 			`db:"Group_ID"`
		Group_Name 			string 			`db:"group_name"`
		Source				string			`db:"source"`
		Source_ID			string			`db:"source_id"`		
	}
	
	var group_id 			Group 
	groupnum, err := strconv.Atoi(groupID)

	if err != nil {
		return 0, err
	}
	group_id.Group_ID = groupnum
	fmt.Printf("\n==============\n Group_ID to be deleted: %+v", group_id.Group_ID)
	
	ownercheck, err := groupOwner(groupID, ownerID)
	if err != nil || ownercheck == false {
		var result int64 
		return result , err 
	}
	
	dbmap 					:= db_connect.InitDb()
	defer dbmap.Db.Close()
	
	dbmap.AddTableWithName(Group{}, "GROUPS").SetKeys(true, "Group_ID")
	
	_, err = dbmap.Delete(&group_id)
	fmt.Printf("\n==============\n deleted: %+v", group_id)
	
	count, err := dbmap.SelectInt("select count(*) from GROUPS where Group_ID=?", group_id.Group_ID)
	fmt.Printf("\n==============\n COUNT: %+v", count)
	
	res, err2 := delGroupUsers(group_id.Group_ID)
	if err2 != nil {
		fmt.Printf("\n===========\n delGroup error: %+v", err)
	} else {
		fmt.Printf("\n===========\n delGroup response: %+v", res)
	}
	
	return count, err
}

func delMember (groupID string, userID string, ownerID string) (sql.Result, error) {	

	ownercheck, err := groupOwner(groupID, ownerID)
	if err != nil || ownercheck == false {
		var result sql.Result 
		return result , err 
	}
	
	dbmap 					:= db_connect.InitDb()
	defer dbmap.Db.Close()

	delete_query := "delete from USER_GROUPS where Group_ID=" + groupID + " AND user_ID=" + userID 
	res, err := dbmap.Exec(delete_query)
	return res, err 
}

func delGroupUsers (groupID int) (sql.Result, error) {
	dbmap 					:= db_connect.InitDb()
	defer dbmap.Db.Close()
		
	res, err := dbmap.Exec("delete from USER_GROUPS where Group_ID=?", groupID)
	
	if err != nil {
		fmt.Printf("\n===========\n delGroupUsers error: %+v", err)
	} else {
		fmt.Printf("\n===========\n delGroupUsers response: %+v", res)
	}
	
	return res, err
}
	
//Maps are not inherently safe for concurrency - will have to use sync.RWMutex	
