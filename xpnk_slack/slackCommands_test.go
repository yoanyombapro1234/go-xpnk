package xpnk_slack

import  (
	"testing"
	"fmt"
)	

func TestSlackGroupStatus(t *testing.T) {

	var 	thisCommand		SlackCommand
	
	thisCommand.Token		= "BCx5C0HMEmWfcpPjxw637Wkt"
	thisCommand.TeamID		= "T3R3D2ERW"
	thisCommand.TeamDomain	= ""		
	thisCommand.ChannelID	= ""
	thisCommand.ChannelName	= ""
	thisCommand.UserID		= "U066LDRB7"
	thisCommand.UserName	= "mspseudolus"
	thisCommand.Command		= "/xapnik"
	thisCommand.Text		= ""
	thisCommand.ResponseURL	= "https://hooks.slack.com/commands/1234/5678"
	thisCommand.TriggerID	= ""
	
	v 			:= SlackGroupStatus(thisCommand)

	if v != "T3R3D2ERW" {
		t.Errorf("Expected Success!, got %v", v)
	}	else {
		fmt.Printf("This is your team ID:   %s/n", v)
	}
}