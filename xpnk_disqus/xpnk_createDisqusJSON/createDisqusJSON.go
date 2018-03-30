package xpnk_createDisqusJSON

/**************************************************************************************
Takes a group ID and writes all Disqus posts for that group to a json file using a file-naming convention.

(1) get all the user id's for a single group
(2) get all the Disqus user names for each of those user ids
(3) get all the Disqus posts associated with each of those disqus_user's
(4) put all the user's Disqus posts into an object keyed by user's Xpnk_ID
(5) json-encode all the Disqus posts and write to a file using a naming convention
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

//stores the Disqus user name and XpnkID for each GroupMemberID
type GroupDisquser struct {
	XpnkID			string
    DisqusUserName	string	`db:"disqus_username"`
}

//stores the Xapnik user_ID of each member of a group
type GroupMemberID struct {
    UserID	string			`db:"user_ID"`
}

//stores the object for an Disqus post
type XpnkDisqus struct {
	DisqusPID		string	`db:"disqus_pid" 		json:"disqus_pid"`
	DisqusUser		string	`db:"disqus_user" 		json:"disqus_user"`
	DisqusName		string	`db:"disqus_name"		json:"disqus_name"`
	DisqusUserID	string	`db:"disqus_userid" 	json:"disqus_userid"`
	DisqusPermalink	string	`db:"disqus_permalink" 	json:"disqus_permalink"`
	DisqusTitle		string	`db:"disqus_title"		json:"disqus_title"`
	DisqusOembed	string	`db:"disqus_embed" 		json:"disqus_embed"`
	DisqusDate		string	`db:"disqus_date"		json:"disqus_date"`
	ProfileImageURL	string	`db:"disqus_avatar" 	json:"disqus_avatar"`
	DisqusForum		string	`db:"disqus_forum"		json:"disqus_forum"`
	DisqusFavicon	string	`db:"disqus_favicon"	json:"disqus_favicon"`
	DisqusMedia		string	`db:"disqus_media"		json:"disqus_media"`
}

type UserDisqus struct {
	XpnkID			string
	DisqusPosts		[]XpnkDisqus
}

func CreateDisqusJSON(group_id int) string {
	
	dbmap := initDb()
	defer dbmap.Db.Close()
	
	var group_name Groupname
	
	/******
	* converting the group name into a string to use in the filename
	******/
	//get group name to use for the filename	
	err := dbmap.SelectOne(&group_name, "SELECT `group_name` FROM `groups` WHERE `Group_ID`=?", group_id)

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
	* get the disqus_username from USERS for each user_ID in group_members
	****/
	var group_disqusers []GroupDisquser

	for i := 0; i < len(group_members); i++ {

		disquser_xpnkid := group_members[i].UserID
				
		thisDisqusUsername, err := dbmap.SelectStr("SELECT `disqus_username` FROM `USERS` WHERE `user_ID`=?", disquser_xpnkid)
		
		fmt.Printf("\n==========\nTHIS DISQUSER:%+v\n==========\n",thisDisqusUsername)
		
		var thisDisquser GroupDisquser
		
		thisDisquser.DisqusUserName = thisDisqusUsername
		thisDisquser.XpnkID = disquser_xpnkid
		
		group_disqusers = append(group_disqusers, thisDisquser)

		checkErr(err, "Select failed")
	}
	
	fmt.Printf("\n==========\nDISQUSERS:%+v\n==========\n",group_disqusers)
	
	/******
	* write the Disqus user names to a file using file-naming convention
	******/
	this_users, err := os.Create("/Users/mizkirsten/Desktop/Node/XAPNIK/data/"+this_name+"_disqus_users.json")
	
	//convert group_disqusers struct to json
	users_str, err := json.Marshal(group_disqusers)
	if err != nil {
		fmt.Println("Error encoding JSON")
	}

		this_users.WriteString(string(users_str))
	/*******
	* get all the group Disqus posts from the db
	*******/ 
	var group_disqusposts 	[]UserDisqus
	
	var this_user 			UserDisqus

	var user_disqusposts 	[]XpnkDisqus
	
	for i := 0; i < len(group_disqusers); i++ {
			
		disquser := group_disqusers[i].DisqusUserName
		
		this_user.XpnkID = group_disqusers[i].XpnkID

		fmt.Printf("\n==========\nDISQUSER:%+v\n==========\n",this_user.XpnkID)
		
		_, err := dbmap.Select(&user_disqusposts, "SELECT * FROM `disqus_comments` WHERE `disqus_user`=?", disquser)
		
		this_user.DisqusPosts = user_disqusposts
		
		fmt.Printf("\n==========\nTHIS USERS POSTS:%+v\n==========\n",this_user)
		
		group_disqusposts = append(group_disqusposts, this_user)
		
		user_disqusposts = []XpnkDisqus{} 
		//this has to be emptied or it carries over to the next disquser
				
		checkErr(err, "Select failed at line 161 createDisqusJSON: ")

	}
	fmt.Printf("\n==========\nGROUP DISQUS POSTS:%+v\n==========\n",group_disqusposts)

	//write the contents of group_disqusposts to a .json file
	//name  file according to file-naming convention
	
	this_file, err := os.Create("/Users/mizkirsten/Desktop/Node/XAPNIK/data/"+this_name+"_disqus.json")
		fmt.Printf("\n==========\nCREATED:%+v\n==========\n",this_file)

	//convert group_disqusposts to json
	str, err := JSONMarshal(group_disqusposts, true)
	if err != nil {
		fmt.Println("Error encoding JSON")
	}

	this_file.WriteString(string(str))
	
	return "File created!"
			
}// end createDisqusJSON

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
	"")
checkErr(err, "sql.Open failed")

dbmap := &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{}}

return dbmap
}

func checkErr(err error, msg string) {
  if err != nil {
  log.Fatalln(msg, err)
	}	
} 