package xpnk_getDisqusThreadDetails

/**************************************************************************************
Takes a Disqus thread ID and retrieves details - particularly 'link' for creating
a link to the exact comment to store in the db
**************************************************************************************/

import (
  "fmt"
  "net/url"
  "xpnk_disqus/golang-disqus/disqus"
)

type DisqusThreadObject struct {
	Link			string			`json:"link"`
	Title			string			`json:"title"`
}

func GetDisqusThreadDetails(disqus_ThreadID string) DisqusThreadObject{

	disqusClient := "Cen6kb83THtogsxE5I5cXh2VTgZyNKhH9th5G2kXsiA1UGYkG5NseX9zh4RO9ERx"
	disqusThread := disqus_ThreadID

	api := disqus.New(disqusClient)
	
	fmt.Println("\n=====================\nSuccessfully created disqus.Api with app credentials")
	fmt.Println("\n=====================\nAPI looks like:  %+v", api)
	
    fmt.Println("\n=====================\ndisqusThread: %v", disqusThread)
	
	params := url.Values{}
	
	if disqusThread != "" {
		params.Set("thread", disqusThread)
	} else {
		panic("Disqus thread ID must be provided.")
	}
	
	fmt.Println("Params: %v", params)
	
	fmt.Println("Getting thread details for thread ID: %v", disqusThread)
			
	disqusThreadResponse, err := api.GetThreadDetails(params)
	
	if err != nil {
		fmt.Printf("At line 40 of getDisqusThreadDetails:  %+v",err)
		fmt.Printf ("At line 40 of getDisqusThreadDetails:  ")

	}	
	
	fmt.Printf("Response Object:  %+v\n", disqusThreadResponse)
	fmt.Printf("\n====================\nThread link: %v\n", disqusThreadResponse.Contents.Link)
	fmt.Printf("\n====================\nThread title: %v\n", disqusThreadResponse.Contents.Title)
	
	var thisDisqusThread DisqusThreadObject
	thisDisqusThread.Link = disqusThreadResponse.Contents.Link
	fmt.Printf("\n====================\nComment ID: %v\n", thisDisqusThread.Link)
	thisDisqusThread.Title = disqusThreadResponse.Contents.Title
	fmt.Printf("\n====================\nComment ID: %v\n", thisDisqusThread.Title)	
		
	return thisDisqusThread
}	