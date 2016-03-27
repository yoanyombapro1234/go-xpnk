package xpnk_createInstaJSON

/***
make json file of each group's Instagrams, use file name convention  'group-name_instas.json'
***/

//(1) get all the user id's for a single group
//(2) get all the instagram user id's for each of those user ids
//(3) get all the instagrams associated with each of those instagram_id's
//(4) json-encode all the instagrams and write to a file using a naming convention

import (
	"fmt"
   	_ "github.com/go-sql-driver/mysql"
    "strings"
    "bytes"
    "os"
    "encoding/json"
)

//stores only the group_name of the group_ID
type Groupname struct {
    GroupName		string		`db:"group_name"`
}

//stores the Instagram user names for each GroupMemberID
type GroupInstagrammer struct {
    InstagramName	string				`db:"insta_name"`
    ProfileImage 	string				`db:"profile_image" json:"profile_image"`
}

//stores only the user_ID's of each member of a group
type GroupMemberID struct {
    UserID	string					`db:"user_ID"`
}

//stores the object for an Instagram post
type XpnkInstagram struct {
	InstagramUser	string	`db:"insta_user" 		json:"user.username"`
	InstagramName	string	`db:"insta_name"		json:"user.full_name"`
	InstagramUserID	string	`db:"insta_userid" 		json:"user.id"`
	InstagramID		string	`db:"instagram_url" 	json:"id"`
	InstagramOembed	string	`db:"instagram_oembed" 	json:"tweet_oembed"`
	InstagramDate	string	`db:"instagram_date"	json:"created_time"`
	ProfileImageURL	string	`db:"profile_image" 	json:"profile_image"`
}

func createInstaJSON() {
	
	dbmap := initDb()
	defer dbmap.Db.Close()
	
	//get ids for all groups	
	group_ids := Get_Groups()

	var this_group int

	//this is where we iterate over each group to get all the Instagram posts per group
	//everything else happens inside this loop
	for i := 0; i < len(group_ids); i++ {
		this_group = group_ids[i].GroupID
		
		var group_name Groupname
	
		/******
		* converting the group name into a string to use in the filename
		******/
		//get group name to use for the filename	
		err := dbmap.SelectOne(&group_name, "SELECT `group_name` FROM `GROUPS` WHERE `Group_ID`=?", this_group)
	
		if err != nil {fmt.Printf("There was an error ", err)}

		fmt.Printf("\n==========\nGROUP NAME:%+v\n",group_name)
	
		//extract just the group name string from group_name	
		this_name := group_name.GroupName
	
		//convert the group name into a hyphenated string for use in json filename
		this_name = strings.Replace(this_name, " ", "-", -1)	
	
		//convert all characters to lowercase
		this_name = strings.ToLower(this_name)
	
		fmt.Printf("\n==========\nGROUP NAME IS NOW:%+v\n",this_name)

		/******
		* get all the xpnk user_ID's associated with a particular Group_ID from 		
		USER_GROUPS
		******/
		var group_members []GroupMemberID

		_, err = dbmap.Select(&group_members, "SELECT `user_ID` FROM `USER_GROUPS` WHERE `Group_ID`=?", this_group)

		//fmt.Printf("\n==========\nMember ID's:%+v\n",group_members)

		checkErr(err, "Select failed")

		/******
		* get the insta_userid from USERS for each user_ID in group_members
		******/		
		var group_instagrammers []GroupInstagrammer

		for i := 0; i < len(group_members); i++ {
	
			instagrammer := group_members[i].UserID
	
			_, err := dbmap.Select(&group_instagrammers, "SELECT `insta_userid`, `profile_image` FROM `USERS` WHERE `user_ID`=?", instagrammer)
	
			checkErr(err, "Select failed")
		}
	
		/******
		* write the Instagram user names to a file using a naming convention
		******/
		this_users, err := os.Create(this_name+"_users.json")
		
		//convert group_instagrammers struct to json
		users_str, err := json.Marshal(group_instagrammers)
		if err != nil {
			fmt.Println("Error encoding JSON")
			return
		}
	
			this_users.WriteString(string(users_str))
	
		/*******
		* get all the groups' Instagram posts from the db
		*******/    
		var group_instagrams []XpnkInstagram

		for i := 0; i < len(group_instagrams); i++ {
	
			instagrammer := group_instagrams[i].InstagramUser
	
			_, err := dbmap.Select(&group_instagrams, "SELECT * FROM `instagram_posts` WHERE `insta_user`=?", instagrammer)
	
			checkErr(err, "Select failed")
	
		}
		fmt.Printf("\n==========\nGROUP INSTAGRAMS:%+v\n==========\n",group_instagrams)

	
		/******
		* write the contents of group_instagrams to a .json file
		******/
		//create the file according to our naming convention
		this_file, err := os.Create("/Users/mizkirsten/Desktop/Node/XAPNIK/data/"+this_name+"_instagrams.json")
			fmt.Printf("\n==========\nCREATED:%+v\n==========\n",this_file)

		//convert group_instagrams to json
		str, err := JSONMarshal(group_instagrams, true)
		if err != nil {
			fmt.Println("Error encoding JSON")
			return
		}
	
			this_file.WriteString(string(str))
	}//end group_ids loop		
			
}//end main  

/******
* keep html tags in tact in our json file
******/
func JSONMarshal(v interface{}, safeEncoding bool) ([]byte, error) {
    b, err := json.Marshal(v)

    if safeEncoding {
        b = bytes.Replace(b, []byte("\\u003c"), []byte("<"), -1)
        b = bytes.Replace(b, []byte("\\u003e"), []byte(">"), -1)
        b = bytes.Replace(b, []byte("\\u0026"), []byte("&"), -1)
        //b = bytes.Replace(b, []byte("\\\""), []byte("\""), -1)
    }
    return b, err
} 