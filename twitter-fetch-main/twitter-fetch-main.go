// Reuse of this code by Xapnik is courtesy of
// Copyright 2013 Ardan Studios. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
This program provides an sample to learn how to implement a timer
routine and graceful shutdown pattern.

Ardan Studios
12973 SW 112 ST, Suite 153
Miami, FL 33186
bill@ardanstudios.com

http://www.goinggo.net/2013/09/timer-routines-and-graceful-shutdowns.html
*/

//this file starts and stops twitter_fetch_manager.go which manages get_tweets.go

package main

import (
	"bufio"
	"os"

	"xpnk-twitter-manager/helper"
	"xpnk_twitter_fetch_mgr"
)

// main is the starting point of the program
func main() {
	helper.WriteStdout("main", "main", "Starting Program")

	xpnk_twitter_fetch_mgr.TWFetchStartup()

	// Hit enter to terminate the program gracefully
	reader := bufio.NewReader(os.Stdin)
	reader.ReadString('\n')

	xpnk_twitter_fetch_mgr.TWFetchShutdown()

	helper.WriteStdout("main", "main", "Program Complete")
}
