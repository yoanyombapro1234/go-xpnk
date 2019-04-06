package slackbots

import (
	"fmt"
	"strconv"
	"strings"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"xpnk-group/xpnk_createGroupFromSlack"
)

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