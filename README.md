# go-xpnk
Golang project - backend to Angular/Ionic social media app

This is the backend that powers my KAngular and Xpnk projects. 

This project:

* Works with a MySQL database which stores (i) groups of Twitter users, (ii) information about each Twitter user from each group, (iii) Tweets from the last 24 hours only by each member of each group.
* Queries Twitter every 60 seconds to get the lastest tweets from each member of all groups in the DB.
* Stores the tweets from the Twitter response in the DB.
* Creates a JSON file of the tweets for each group, using a file-naming convention, that is used by the front end (Angular and Ionic) apps.
* Provides a REST API endpoint from which the front end retrieves the JSON files which it uses to assemble the display of the tweets.
* Currently in the process of adding the functionality for Instagram

Eventually, the Go code will send push notifications to the front end apps only when there is new content in the applicable JSON files (rather than having the front end requesting the JSON file from the server every 60 seconds). The front end code will then respond with a request to the REST API end point.

Very much a work in progress and not an alpha, beta or release candidate :) Expect a LOT of `fmt.Printf` output and comments to facilitate ongoing development.
