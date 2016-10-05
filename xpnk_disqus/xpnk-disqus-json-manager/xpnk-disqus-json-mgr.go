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

//this file starts & stops json-manager.go which mgs createDisqusJSON.go

package main

import (
	"bufio"
	"os"
	
	"xpnk_disqus/json-manager"
	"xpnk-shared/manager-helper/helper"
	"sync/atomic"
	"time"

)

const (
	timerPeriod time.Duration = 6000 * time.Second // Interval to wake up on.
)

// DisqusjsonWorkManager is responsible for starting and shutting down the program.
type DisqusjsonWorkManager struct {
	Shutdown        int32
	ShutdownChannel chan string
}

var Disqus_json_wm DisqusjsonWorkManager // Reference to the singleton.

// main is the starting point of the program
func main() {
	helper.WriteStdout("main", "main", "Starting Program")

	DisqusJSONStartup()

	// Hit enter to terminate the program gracefully
	reader := bufio.NewReader(os.Stdin)
	reader.ReadString('\n')

	DisqusJSONShutdown()

	helper.WriteStdout("main", "main", "Program Complete")
}  
// end main

// DisqusJSONStartup brings the manager to a running state.

func DisqusJSONStartup() error {
	var err error
	defer helper.CatchPanic(&err, "main", "DisqusJSONStartup")

	helper.WriteStdout("main", "DisqusJSONStartup", "Started")

	// Create the work manager to get the program going
	Disqus_json_wm = DisqusjsonWorkManager{
		Shutdown:        0,
		ShutdownChannel: make(chan string),
	}

	// Start the work timer routine.
	// When DisqusjsonWorkManager returns the program terminates.
	go Disqus_json_wm.GoRoutineworkTimer()

	helper.WriteStdout("main", "DisqusJSONStartup", "Completed")
	return err
}

// DisqusJSONShutdown brings down the manager gracefully.

func DisqusJSONShutdown() error {
	var err error
	defer helper.CatchPanic(&err, "main", "DisqusJSONShutdown")

	helper.WriteStdout("main", "DisqusJSONShutdown", "Started")

	// Shutdown the program
	helper.WriteStdout("main", "DisqusJSONShutdown", "Info : Shutting Down")
	atomic.CompareAndSwapInt32(&Disqus_json_wm.Shutdown, 0, 1)

	helper.WriteStdout("main", "DisqusJSONShutdown", "Info : Shutting Down Work Timer")
	Disqus_json_wm.ShutdownChannel <- "Down"
	<-Disqus_json_wm.ShutdownChannel

	close(Disqus_json_wm.ShutdownChannel)

	helper.WriteStdout("main", "DisqusJSONShutdown", "Completed")
	return err
}

// GoRoutineworkTimer perform the work on the defined interval.
func (DisqusjsonWorkManager *DisqusjsonWorkManager) GoRoutineworkTimer() {
	helper.WriteStdout("WorkTimer", "DisqusjsonWorkManager.GoRoutineworkTimer", "Started")

	wait := timerPeriod

	for {
		helper.WriteStdoutf("WorkTimer", "DisqusjsonWorkManager.GoRoutineworkTimer", "Info : Wait To Run : Seconds[%.0f]", wait.Seconds())

		select {
		case <-DisqusjsonWorkManager.ShutdownChannel:
			helper.WriteStdoutf("WorkTimer", "DisqusjsonWorkManager.GoRoutineworkTimer", "Shutting Down")
			DisqusjsonWorkManager.ShutdownChannel <- "Down"
			return

		case <-time.After(wait):
			helper.WriteStdoutf("WorkTimer", "DisqusjsonWorkManager.GoRoutineworkTimer", "Woke Up")
			break
		}

		// Mark the starting time
		startTime := time.Now()

		// Perform the work
		DisqusjsonWorkManager.PerformTheWork()

		// Mark the ending time
		endTime := time.Now()

		// Calculate the amount of time to wait to start DisqusjsonWorkManager again.
		duration := endTime.Sub(startTime)
		wait = timerPeriod - duration
	}
}

// PerformTheWork with sleep times.
func (DisqusjsonWorkManager *DisqusjsonWorkManager) PerformTheWork() {
	defer helper.CatchPanic(nil, "DisqusjsonWorkManager", "DisqusjsonWorkManager.PerformTheWork")
	helper.WriteStdout("WorkTimer", "DisqusjsonWorkManager.GoRoutineworkTimer", "Started")

	//run json-manager.go
	disqus_json_workmanager.Create_group_disqus_json()
	
	helper.WriteStdout("WorkTimer", "WorkManager.GoRoutineworkTimer", "Completed")
}