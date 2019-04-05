package users

import (
	"strings"
	 _ "github.com/go-sql-driver/mysql"
	 "github.com/gin-gonic/gin"
	 "xpnk-shared/db_connect"
	 "xpnk-user/xpnk_user_structs"
)

func GetGroups (c *gin.Context) { 
	var groups					[]xpnk_user_structs.GroupOwner
	var group_trim				xpnk_user_structs.GroupsByUser
	var	groups_trim				[]xpnk_user_structs.GroupsByUser
	var user_groups				xpnk_user_structs.UserGroups
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
				var this_group xpnk_user_structs.GroupOwner
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
			user_groups.Xpnk_id   		= user_id
			xpnk_name, err 				:= GetUserName(user_id)
			if err != nil {user_groups.ScreenName = ""} else {
			user_groups.ScreenName 		= xpnk_name }
			user_groups.Groups 	  		= groups_trim
			c.JSON(200, user_groups)
		}
	}
} 

func get_user_groups(user_id string) ([]xpnk_user_structs.GroupOwner, error) {
	dbmap := db_connect.InitDb()
	defer dbmap.Db.Close()

	var groupOwners			[]xpnk_user_structs.GroupOwner
	var err_msg				error
	
	id						:= user_id
	
	_, err := dbmap.Select(&groupOwners, "SELECT `USER_GROUPS`.`Group_ID`, `USER_GROUPS`.`group_owner`, `USER_GROUPS`.`group_admin`, `GROUPS`.`group_name` FROM USER_GROUPS INNER JOIN GROUPS ON `USER_GROUPS`.`Group_ID` = `GROUPS`.`Group_ID` WHERE `USER_GROUPS`.`user_ID` =?", id)
			
	if err != nil {
		err_msg				= err
	} 
	
	return groupOwners, err_msg
} 