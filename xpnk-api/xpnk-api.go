package main

import (
         "github.com/gin-gonic/gin"
   		 _ "github.com/go-sql-driver/mysql"
   		 "xpnk_constants"
   		 "xpnk-api/users"
   		 "xpnk-api/groups"
   		 "xpnk-api/slackbots"
 )

const (
	mySigningKey = xpnk_constants.SigningKey
)
	 
func main() {

	r := gin.Default()
	r.Use(Cors())
	r.Static("api/v1/data", "../xpnk-data/")
	r.Static("api/v2/data", "../xpnk-data/")
	
	
	r.GET("/ping", func(c *gin.Context) {
        c.String(200, "Hello")
      })
      
    r.POST("/ping", func(c *gin.Context) {
        c.String(200, "Hello")
      })    
      
    v2 := r.Group("api/v2")
		{  
			v2.GET ("/ping", func(c *gin.Context) {
				c.String(200, "pong")
			})
			
			v2.GET("/users/twitter/:id", users.UsersByTwitterID_2)
			v2.GET("/users/ig/:id", users.UsersByIGID_2)
			v2.GET ("/users/invite", users.CheckUserInvite)
			v2.GET ("/users/authSet", users.XPNKAuthSet)
			v2.GET ("/users/groups/:id", users.GetGroups)
			v2.GET ("/users/login/twitter", users.LoginTwitter)
			v2.GET ("/users/login/insta", users.LoginInsta)
			
			v2.POST("/users", users.UsersNew_2)
			v2.POST ("/users/authCheck", users.XPNKAuthCheck)
			
			v2.PUT("/users", users.UsersUpdate_2)
			
			v2.DELETE("/users/:id", users.UsersDelete)
			
			v2.GET ("/xpnk_auth_set", users.XPNKAuthSet)
			
			
			v2.GET ("/groups/:id/members", groups.GroupsByID)
			v2.GET ("/groups/:id/invite/:source", groups.GroupsInvite)
			
			v2.POST("/groups/", groups.GroupsNew)
			v2.POST("/groups/add", groups.GroupsAddMember)
			
			v2.DELETE("/groups/:id/owner/:owner", groups.GroupsDelete)
			v2.DELETE("groups/:id/user/:user/owner/:owner", groups.GroupsMemberDelete)
			
		}

/*****************************************
* V1
*****************************************/
	
	v1 := r.Group("api/v1")
		{
			v1.GET ("/ping", func(c *gin.Context) {
				c.String(200, "pong")
			})
						
			v1.GET("/users/twitter", users.UsersByTwitterID)
			v1.GET("/users/ig", users.UsersByIGID)
			v1.GET ("/get_xpnkid/slack/:id", users.GetXPNK_ID)
			v1.POST("/users/new", users.UsersNew)
			v1.POST("/users/update", users.UsersUpdate)
			
			v1.GET ("/groups/members/:id", groups.GroupsByID)
			v1.GET ("/groups/id/:name", groups.GroupID)
			v1.POST("/groups/add", groups.GroupsAddMember)
			
			v1.GET ("/slack_response", slackbots.SlackResponseHandler)
			v1.GET ("/slack_response/command", slackbots.SlackCommandHandler)
			v1.POST ("slack_new_member", users.SlackNewMember)			
			v1.POST ("/slack_new_group", groups.SlackCreateNewGroup)
			v1.POST ("/slack_response", slackbots.SlackResponseHandler)
			v1.POST ("/slack_response/command", slackbots.SlackCommandHandler)
			
			v1.GET ("/check_user_invite", users.CheckUserInvite)
			v1.GET ("/xpnk_auth_set", users.XPNKAuthSet)
			v1.POST ("/xpnk_auth_check", users.XPNKAuthCheck)
				
			v1.POST("/twitter_auth", users.PostTwttrAuth)
		}
		
	r.Run(":9090")
}	

func Cors() gin.HandlerFunc {
 return func(c *gin.Context) {
 c.Writer.Header().Add("Access-Control-Allow-Origin", "*")
 c.Next()
 }
}