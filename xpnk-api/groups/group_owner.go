package groups

import (
	"fmt"
	"errors"
	"strconv"
	  _ "github.com/go-sql-driver/mysql"
	 "xpnk-shared/db_connect"
)

func GroupOwner (groupID string, ownerID string) (bool, error) {
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