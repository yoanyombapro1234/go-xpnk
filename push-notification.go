package main

import (
"fmt"
"github.com/parnurzeal/gorequest"
)

func main () {

	request := gorequest.New().SetBasicAuth("8f1d650549cb63c60274ea359e1ac409f1439fa5c4c94b25", "")
	
	resp, body, errs := request.Post("https://push.ionic.io/api/v1/push").
	Set("Content-Type","application/json").
	Set("X-Ionic-Application-Id","39fb755e").	
	Send(`{"tokens":["DEV-41c11b25-af7d-4cb1-84f1-6b7e2edfdc68", "DEV-4300e98b-e1a5-40d3-a85c-b5f1831d1ee7"], "notification":{"alert":"I come from planet Ion."}}`).
	End()	
	
	defer resp.Body.Close()
	fmt.Printf("\n==========\nBODY: \n%+v\n",body)
	fmt.Printf("\n==========\nERRS: \n%+v\n",errs)
}
