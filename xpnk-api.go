package main

import (
    "github.com/gin-gonic/gin"
    "fmt"
    "log"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
    "github.com/gopkg.in/gorp.v1"
)
 
type NewMobileReg struct {
	 Received		string					`json:"received"`
	 Userid			string					`json:"user_id"`
	 Name			string					`json:"name"`
	 Appid			string					`json:"app_id"`
	 Push			map[string][]string		`json:"_push"`
	 Message		string					`json:"message"`
}

type MobileRegInsert struct {
	 Received		string					`db:"received"`
	 Userid			string					`db:"mobile_user_id"`
	 Name			string					`db:"mobile_name"`
	 Appid			string					`db:"-"`
	 Androidtokens	string					`db:"android_tokens"`
	 IOSTokens		string					`db:"ios_tokens"`			
	 Message		string					`db:"message"`
}
 
func main() {
         
	r := gin.Default()
	
	r.GET("/ping", func(c *gin.Context) {
        c.String(200, "Hello")
      })  
	
	v1 := r.Group("api/v1")
		{
			v1.POST("/users", PostUser)
		}
		
	r.Run(":9090")
}	

func PostUser(c *gin.Context) {
	
	var user NewMobileReg
	c.Bind(&user)
		
	if user.Userid != "" {
	
		 content := &NewMobileReg{
		 	 Received:		user.Received,
			 Userid:		user.Userid,
			 Name:			user.Name,
			 Appid:			user.Appid,
			 Message:		user.Message,
			 Push:			user.Push,
		 }
		 
		 c.JSON(201, content)
		 doinsert(user)
	} else {
		c.JSON(422, gin.H{"error": "No user_id was sent."})
	}	 	
}

/***************************
*DATABASE INSERT FUNCTION
***************************/

func doinsert(newmobilereg NewMobileReg) string{

	insert := create_insert(newmobilereg)

	dbmap := initDb()
	defer dbmap.Db.Close()
	
//map the MobileRegInsert struct to the 'install_ids' db table
	dbmap.AddTableWithName(MobileRegInsert{}, "install_ids")
		
//Insert the NewMobileReg!	
			
	//db insert function 
	err := dbmap.Insert(&insert)
	if err != nil {fmt.Printf("There was an error ", err)		
	}
	
	return "inserted" 
}	

func create_insert(n NewMobileReg) (i MobileRegInsert) {

	//pull the list of Android or IOS tokens out of the _pull map and put into a string that can be inserted into a field in the db

	this_ios := n.Push["ios_tokens"]
	this_android := n.Push["android_tokens"]
	
	var ThisIOSTokens string
	for _, value := range this_ios {
	 	ThisIOSTokens += value + ","
	 }
	 	
	var ThisAndroidTokens string 
	for _, value := range this_android {
	 	ThisAndroidTokens += value + ","
	}

	i.Received = n.Received
	i.Userid = n.Userid
	i.Name = n.Name
	i.Androidtokens = ThisAndroidTokens
	i.IOSTokens = ThisIOSTokens
	i.Message = n.Message
		
	return i
}
//end db insert routine	
//Maps are not inherently safe for concurrency - will have to use sync.RWMutex	


/***************************
* db connection config
***************************/

func checkErr(err error, msg string) {
    if err != nil {
        log.Fatalln(msg, err)
	}
}	
	
func initDb() *gorp.DbMap {
db, err := sql.Open("mysql",
	"root:root@tcp(localhost:8889)/password")
checkErr(err, "sql.Open failed")

dbmap := &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{}}

return dbmap
}