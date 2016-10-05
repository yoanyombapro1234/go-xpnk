# go-xpnk
Golang project - backend to Angular/Ionic social media app

This is the backend that powers my KAngular and Xpnk projects. 

This project:

* Works with a MySQL database which stores (i) groups of Twitter users, (ii) information about each Twitter user from each group, (iii) Tweets from the last 24 hours only by each member of each group.
* Queries Twitter every 60 seconds to get the lastest tweets from each member of all groups in the DB.
* Queries Instagram every 10 minutes to get the lastest grams from each member of all groups in the DB.
* Queries Disqus every xx minutes to fetch lastest comments by each member of all groups in the DB.
* Stores all posts data in the DB.
* Creates JSON files of posts for each group, using a file-naming convention, that is used by the front end (Angular and Ionic) apps.
* Provides a REST API endpoint from which the front end retrieves the JSON files which it uses to assemble the display of the tweets.
* Provides REST API endpoints for authenticating with Instagram, Disqus (Slack and Twitter coming soon).
* Provides REST API for creating a new Xapnik group from a Slack team.

Eventually, the Go code will send push notifications to the front end apps only when there is new content in the applicable JSON files (rather than having the front end requesting the JSON file from the server every 60 seconds). The front end code will then respond with a request to the REST API end point.

Very much a work in progress and not an alpha, beta or release candidate :) Expect a LOT of `fmt.Printf` output and comments to facilitate ongoing development.
