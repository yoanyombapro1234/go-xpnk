package xpnk_createMultiUserInsert

/**************************************************************************************
Takes a slice of User_Insert and prepares them for batch insertion into the USERS table
**************************************************************************************/

import (
	"fmt"
	"xpnk-user/xpnk_createUserInsert"
)

func CreateMultiUserInsert(userBatch []xpnk_createUserInsert.User_Insert) []xpnk_createUserInsert.User_Insert {

	var thisBatchInsert []xpnk_createUserInsert.User_Insert

	for i := 0; i < len(userBatch); i++ {
	
		//run each user through xpnk_createUserInsert
		//take returned struct from xpnk_createUserInsert and append to thisBatchInsert

		var this_user xpnk_createUserInsert.User_Insert
		this_user = userBatch[i]
	
		this_user_insert := xpnk_createUserInsert.CreateUserInsert(this_user)
						
		fmt.Printf("\n==========\nTHIS_INSERT: \n%+v\n",this_user_insert)

		thisBatchInsert = append(thisBatchInsert, this_user_insert)
	}//end looping through all posts

	fmt.Printf("\n==========\nTHIS_BATCH: \n%+v\n",thisBatchInsert)

	return thisBatchInsert
	
}//end createMultiUserInsert	