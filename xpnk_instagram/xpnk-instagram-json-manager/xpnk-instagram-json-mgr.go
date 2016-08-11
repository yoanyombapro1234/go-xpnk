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

//this file starts & stops instagram_json_workmanager which mgs createInstaJSON.go

package main

import (
	"bufio"
	"os"
	"xpnk_instagram/json-manager"
	"xpnk_instagram/xpnk-instagram-manager/helper"
	"sync/atomic"
	"time"

)

const (
	timerPeriod time.Duration = 600 * time.Second // Interval to wake up on.
)

// IGjsonWorkManager is responsible for starting and shutting down the program.
type IGjsonWorkManager struct {
	Shutdown        int32
	ShutdownChannel chan string
}

var IG_json_wm IGjsonWorkManager // Reference to the singleton.

// main is the starting point of the program
func main() {
	helper.WriteStdout("main", "main", "Starting Program")

	IGJSONStartup()

	// Hit enter to terminate the program gracefully
	reader := bufio.NewReader(os.Stdin)
	reader.ReadString('\n')

	IGJSONShutdown()

	helper.WriteStdout("main", "main", "Program Complete")
}  
// end main

// IGJSONStartup brings the manager to a running state.

func IGJSONStartup() error {
	var err error
	defer helper.CatchPanic(&err, "main", "IGJSONStartup")

	helper.WriteStdout("main", "IGJSONStartup", "Started")

	// Create the work manager to get the program going
	IG_json_wm = IGjsonWorkManager{
		Shutdown:        0,
		ShutdownChannel: make(chan string),
	}

	// Start the work timer routine.
	// When IGjsonWorkManager returns the program terminates.
	go IG_json_wm.GoRoutineworkTimer()

	helper.WriteStdout("main", "IGJSONStartup", "Completed")
	return err
}

// IGJSONShutdown brings down the manager gracefully.

func IGJSONShutdown() error {
	var err error
	defer helper.CatchPanic(&err, "main", "IGJSONShutdown")

	helper.WriteStdout("main", "IGJSONShutdown", "Started")

	// Shutdown the program
	helper.WriteStdout("main", "IGJSONShutdown", "Info : Shutting Down")
	atomic.CompareAndSwapInt32(&IG_json_wm.Shutdown, 0, 1)

	helper.WriteStdout("main", "IGJSONShutdown", "Info : Shutting Down Work Timer")
	IG_json_wm.ShutdownChannel <- "Down"
	<-IG_json_wm.ShutdownChannel

	close(IG_json_wm.ShutdownChannel)

	helper.WriteStdout("main", "IGJSONShutdown", "Completed")
	return err
}

// GoRoutineworkTimer perform the work on the defined interval.
func (IGjsonWorkManager *IGjsonWorkManager) GoRoutineworkTimer() {
	helper.WriteStdout("WorkTimer", "IGjsonWorkManager.GoRoutineworkTimer", "Started")

	wait := timerPeriod

	for {
		helper.WriteStdoutf("WorkTimer", "IGjsonWorkManager.GoRoutineworkTimer", "Info : Wait To Run : Seconds[%.0f]", wait.Seconds())

		select {
		case <-IGjsonWorkManager.ShutdownChannel:
			helper.WriteStdoutf("WorkTimer", "IGjsonWorkManager.GoRoutineworkTimer", "Shutting Down")
			IGjsonWorkManager.ShutdownChannel <- "Down"
			return

		case <-time.After(wait):
			helper.WriteStdoutf("WorkTimer", "IGjsonWorkManager.GoRoutineworkTimer", "Woke Up")
			break
		}

		// Mark the starting time
		startTime := time.Now()

		// Perform the work
		IGjsonWorkManager.PerformTheWork()

		// Mark the ending time
		endTime := time.Now()

		// Calculate the amount of time to wait to start IGjsonWorkManager again.
		duration := endTime.Sub(startTime)
		wait = timerPeriod - duration
	}
}

// PerformTheWork with sleep times.
func (IGjsonWorkManager *IGjsonWorkManager) PerformTheWork() {
	defer helper.CatchPanic(nil, "IGjsonWorkManager", "IGjsonWorkManager.PerformTheWork")
	helper.WriteStdout("WorkTimer", "IGjsonWorkManager.GoRoutineworkTimer", "Started")

	//run json-manager.go
	instagram_json_workmanager.Create_group_ig_json()
	
	helper.WriteStdout("WorkTimer", "WorkManager.GoRoutineworkTimer", "Completed")
}