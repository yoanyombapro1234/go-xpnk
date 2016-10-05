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

//this file starts and stops xpnk-instagram-fetch-workmanager.go which manages posts-manager.go

package main

import (
	"bufio"
	"os"
	"xpnk-shared/manager-helper/helper"
	"xpnk_disqus/xpnk-disqus-fetch-manager"
)

// main is the starting point of the program
func main() {
	helper.WriteStdout("main", "main", "Starting Program")

	xpnk_disqus_fetch_workmanager.DisqusFetchStartup()

	// Hit enter to terminate the program gracefully
	reader := bufio.NewReader(os.Stdin)
	reader.ReadString('\n')

	xpnk_disqus_fetch_workmanager.DisqusFetchShutdown()

	helper.WriteStdout("main", "main", "Program Complete")
}
