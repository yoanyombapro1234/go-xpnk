package xpnk_createInstaJSON

/**************************************************************************************
Takes a group ID and writes all Instagram posts for that group to a json file using a file-naming convention.

(1) get all the user id's for a single group
(2) get all the instagram user id's for each of those user ids
(3) get all the instagrams associated with each of those instagram_id's
(4) put all the user's instagrams into an object keyed by user's Xpnk_ID
(5) json-encode all the instagrams and write to a file using a naming convention
**************************************************************************************/

import (
	"fmt"
	"database/sql"
   	_ "github.com/go-sql-driver/mysql"
   	"github.com/gopkg.in/gorp.v1"
    "strings"
    "bytes"
    "os"
    "encoding/json"
    "log"
)

//stores only the group_name of the group_ID
type Groupname struct {
    GroupName		string	`db:"group_name"`
}

//stores the Instagram user name and XpnkID for each GroupMemberID
type GroupInstagrammer struct {
	XpnkID			string
    InstagramUserID	string	`db:"insta_userid"`
}

//stores the Xapnik user_ID's of each member of a group
type GroupMemberID struct {
    UserID	string			`db:"user_ID"`
}

//stores the object for an Instagram post
type XpnkInstagram struct {
	InstagramUser	string	`db:"insta_user" 		json:"instagram_user"`
	InstagramName	string	`db:"insta_name"		json:"instagram_name"`
	InstagramUserID	string	`db:"insta_userid" 		json:"instagram_id"`
	InstagramID		string	`db:"instagram_url" 	json:"instagram_url"`
	InstagramOembed	string	`db:"instagram_oembed" 	json:"instagram_oembed"`
	InstagramDate	string	`db:"instagram_date"	json:"created_time"`
	ProfileImageURL	string	`db:"instagram_avatar" 	json:"profile_image"`
}

type UserInstagrams struct {
	XpnkID			string
	InstagramPosts	[]XpnkInstagram
}

func CreateInstaJSON(group_id int) string {
	
	dbmap := initDb()
	defer dbmap.Db.Close()
	
	var group_name Groupname
	
	/******
	* converting the group name into a string to use in the filename
	******/
	//get group name to use for the filename	
	err := dbmap.SelectOne(&group_name, "SELECT `group_name` FROM `GROUPS` WHERE `Group_ID`=?", group_id)

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
	* get all the xpnk user_ID's associated with the Group_ID from USER_GROUPS
	******/
	var group_members []GroupMemberID

	_, err = dbmap.Select(&group_members, "SELECT `user_ID` FROM `USER_GROUPS` WHERE `Group_ID`=?", group_id)

	fmt.Printf("\n==========\nMember ID's:%+v\n",group_members)

	checkErr(err, "Select failed")
	/****
	* get the insta_userid from USERS for each user_ID in group_members
	****/
	var group_instagrammers []GroupInstagrammer

	for i := 0; i < len(group_members); i++ {

		instagrammer := group_members[i].UserID
				
		thisInstaId, err := dbmap.SelectStr("SELECT `insta_userid` FROM `USERS` WHERE `user_ID`=?", instagrammer)
		
		fmt.Printf("\n==========\nTHIS INSTAGRAMMER:%+v\n==========\n",thisInstaId)
		
		var thisInstagrammer GroupInstagrammer
		
		thisInstagrammer.InstagramUserID = thisInstaId
		thisInstagrammer.XpnkID = instagrammer
		
		group_instagrammers = append(group_instagrammers, thisInstagrammer)

		checkErr(err, "Select failed")
	}
	
	fmt.Printf("\n==========\nINSTAGRAMMERS:%+v\n==========\n",group_instagrammers)
	
	/******
	* write the Instagram user names to a file using file-naming convention
	******/
	this_users, err := os.Create("/home/xapnik/node-v0.12.5/XAPNIK/data/"+this_name+"_insta_users.json")
	
	//convert group_instagrammers struct to json
	users_str, err := json.Marshal(group_instagrammers)
	if err != nil {
		fmt.Println("Error encoding JSON")
	}

		this_users.WriteString(string(users_str))
	/*******
	* get all the group Instagram posts from the db
	*******/ 
	var group_instagrams []UserInstagrams
	
	var this_user  UserInstagrams

	var user_instagrams []XpnkInstagram
	
	for i := 0; i < len(group_instagrammers); i++ {
			
		instagrammer := group_instagrammers[i].InstagramUserID
		
		this_user.XpnkID = group_instagrammers[i].XpnkID

		fmt.Printf("\n==========\nINSTAGRAMMER:%+v\n==========\n",this_user.XpnkID)
		
		_, err := dbmap.Select(&user_instagrams, "SELECT * FROM `instagram_posts` WHERE `insta_userid`=?", instagrammer)
		
		this_user.InstagramPosts = user_instagrams
		
		fmt.Printf("\n==========\nTHIS USERS POSTS:%+v\n==========\n",this_user)
		
		group_instagrams = append(group_instagrams, this_user)
		
		user_instagrams = []XpnkInstagram{} 
		//this has to be emptied or it carries over to the next instagrammer
				
		checkErr(err, "Select failed")

	}
	fmt.Printf("\n==========\nGROUP INSTAGRAMS:%+v\n==========\n",group_instagrams)

	//write the contents of group_instagrams to a .json file
	//name  file according to file-naming convention
	
	this_file, err := os.Create("/Users/mizkirsten/Desktop/Node/XAPNIK/data/"+this_name+"_instagrams.json")
		fmt.Printf("\n==========\nCREATED:%+v\n==========\n",this_file)

	//convert group_instagrams to json
	str, err := JSONMarshal(group_instagrams, true)
	if err != nil {
		fmt.Println("Error encoding JSON")
	}

	this_file.WriteString(string(str))
	
	return "File created!"
			
}// end createInstaJSON

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

/***************************
* db connection config
***************************/	
func initDb() *gorp.DbMap {
db, err := sql.Open("mysql",
	"root:hqao79eJegoZfXLMVpoCeQtZjpVa@tcp(localhost:3306)/xapnik")
checkErr(err, "sql.Open failed")

dbmap := &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{}}

return dbmap
}

func checkErr(err error, msg string) {
  if err != nil {
  log.Fatalln(msg, err)
	}	
} 