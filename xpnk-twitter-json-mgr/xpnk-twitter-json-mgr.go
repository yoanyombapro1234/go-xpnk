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

//this file starts & stops twitter_json_workmanager which mgs create-group-tweets-json

package main

import (
	"bufio"
	"os"
	
	"xpnk_twitter"
	"xpnk-twitter-manager/helper"
	"sync/atomic"
	"time"

)

const (
	timerPeriod time.Duration = 60 * time.Second // Interval to wake up on.
)

// TWjsonWorkManager is responsible for starting and shutting down the program.
type TWjsonWorkManager struct {
	Shutdown        int32
	ShutdownChannel chan string
}

var Tw_json_wm TWjsonWorkManager // Reference to the singleton.

// main is the starting point of the program
func main() {
	helper.WriteStdout("main", "main", "Starting Program")

	TWJSONStartup()

	// Hit enter to terminate the program gracefully
	reader := bufio.NewReader(os.Stdin)
	reader.ReadString('\n')

	TWJSONShutdown()

	helper.WriteStdout("main", "main", "Program Complete")
}  
// end main

// TWJSONStartup brings the manager to a running state.

func TWJSONStartup() error {
	var err error
	defer helper.CatchPanic(&err, "main", "TWJSONStartup")

	helper.WriteStdout("main", "TWJSONStartup", "Started")

	// Create the work manager to get the program going
	Tw_json_wm = TWjsonWorkManager{
		Shutdown:        0,
		ShutdownChannel: make(chan string),
	}

	// Start the work timer routine.
	// When TWjsonWorkManager returns the program terminates.
	go Tw_json_wm.GoRoutineworkTimer()

	helper.WriteStdout("main", "TWJSONStartup", "Completed")
	return err
}

// TWJSONShutdown brings down the manager gracefully.

func TWJSONShutdown() error {
	var err error
	defer helper.CatchPanic(&err, "main", "TWJSONShutdown")

	helper.WriteStdout("main", "TWJSONShutdown", "Started")

	// Shutdown the program
	helper.WriteStdout("main", "TWJSONShutdown", "Info : Shutting Down")
	atomic.CompareAndSwapInt32(&Tw_json_wm.Shutdown, 0, 1)

	helper.WriteStdout("main", "TWJSONShutdown", "Info : Shutting Down Work Timer")
	Tw_json_wm.ShutdownChannel <- "Down"
	<-Tw_json_wm.ShutdownChannel

	close(Tw_json_wm.ShutdownChannel)

	helper.WriteStdout("main", "TWJSONShutdown", "Completed")
	return err
}

// GoRoutineworkTimer perform the work on the defined interval.
func (TWjsonWorkManager *TWjsonWorkManager) GoRoutineworkTimer() {
	helper.WriteStdout("WorkTimer", "TWjsonWorkManager.GoRoutineworkTimer", "Started")

	wait := timerPeriod

	for {
		helper.WriteStdoutf("WorkTimer", "TWjsonWorkManager.GoRoutineworkTimer", "Info : Wait To Run : Seconds[%.0f]", wait.Seconds())

		select {
		case <-TWjsonWorkManager.ShutdownChannel:
			helper.WriteStdoutf("WorkTimer", "TWjsonWorkManager.GoRoutineworkTimer", "Shutting Down")
			TWjsonWorkManager.ShutdownChannel <- "Down"
			return

		case <-time.After(wait):
			helper.WriteStdoutf("WorkTimer", "TWjsonWorkManager.GoRoutineworkTimer", "Woke Up")
			break
		}

		// Mark the starting time
		startTime := time.Now()

		// Perform the work
		TWjsonWorkManager.PerformTheWork()

		// Mark the ending time
		endTime := time.Now()

		// Calculate the amount of time to wait to start TWjsonWorkManager again.
		duration := endTime.Sub(startTime)
		wait = timerPeriod - duration
	}
}

// PerformTheWork with sleep times.
func (TWjsonWorkManager *TWjsonWorkManager) PerformTheWork() {
	defer helper.CatchPanic(nil, "TWjsonWorkManager", "TWjsonWorkManager.PerformTheWork")
	helper.WriteStdout("WorkTimer", "TWjsonWorkManager.GoRoutineworkTimer", "Started")

	//run create-group-tweets-json.go
	xpnk_twitter.Create_group_tweets_json()
	
	helper.WriteStdout("WorkTimer", "WorkManager.GoRoutineworkTimer", "Completed")
}
