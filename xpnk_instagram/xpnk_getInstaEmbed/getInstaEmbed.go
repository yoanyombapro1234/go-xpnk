package xpnk_getInstaEmbed

import (
  "fmt"
  "net/http"
  "encoding/json"
  "io/ioutil"
)

type instaEmbed struct {
	Html	string	`json:"html"`
}
	
func getInstaEmbed(instaUrl string) {

	instaEmbedEndPt := "https://api.instagram.com/oembed/?url="
	InstaEmbedCall := instaEmbedEndPt+instaUrl
	fmt.Printf("Embed Call URL  %v\n", InstaEmbedCall)
		
	resp, err := http.Get(InstaEmbedCall)
	if err != nil {
		panic(err.Error())
	}	
	defer resp.Body.Close()
		
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err.Error())
	}	
	
	thisEmbed, err := getOembedHtml([]byte(body))
	fmt.Printf("\nHTML: %v\n", thisEmbed)
}

func getOembedHtml(body []byte) (*instaEmbed, error) {
	var h = new(instaEmbed)
	err := json.Unmarshal(body, &h)
	if(err != nil) {
		fmt.Println("whoops: ", err)
	}
	return h, err
}
