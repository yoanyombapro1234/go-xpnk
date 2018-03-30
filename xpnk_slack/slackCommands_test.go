package xpnk_slack

import  (
	"testing"
	"fmt"
)	

func TestSlackGroupStatus(t *testing.T) {

	var 	thisCommand		SlackCommand
	
	thisCommand.Token		= ""
	thisCommand.TeamID		= ""
	thisCommand.TeamDomain	= ""		
	thisCommand.ChannelID	= ""
	thisCommand.ChannelName	= ""
	thisCommand.UserID		= ""
	thisCommand.UserName	= ""
	thisCommand.Command		= "/xapnik"
	thisCommand.Text		= ""
	thisCommand.ResponseURL	= ""
	thisCommand.TriggerID	= ""
	
	v 			:= SlackGroupStatus(thisCommand)

	if v != "T3R3D2ERW" {
		t.Errorf("Expected Success!, got %v", v)
	}	else {
		fmt.Printf("This is your team ID:   %s/n", v)
	}
}