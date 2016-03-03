// Xapnik gratefully uses this code Copyright 2013 Ardan Studios. All rights reserved.
// Use of TWFetchWorkManager source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package workmanager implements the WorkManager singleton. This manager
// controls the starting, shutdown and processing of work.


//This workmanager manages get_tweets.go


package xpnk_twitter_fetch_mgr

import (
	"xpnk_twitter"
	"xpnk-twitter-manager/helper"
	"sync/atomic"
	"time"
)

const (
	timerPeriod time.Duration = 60 * time.Second // Interval to wake up on.
)

// TWFetchWorkManager is responsible for starting and shutting down the program.
type TWFetchWorkManager struct {
	Shutdown        int32
	ShutdownChannel chan string
}

var Tw_fetch_wm TWFetchWorkManager // Reference to the singleton.

// TWFetchStartup brings the manager to a running state.
func TWFetchStartup() error {
	var err error
	defer helper.CatchPanic(&err, "main", "TWFetchStartup")

	helper.WriteStdout("main", "TWFetchStartup", "Started")

	// Create the work manager to get the program going
	Tw_fetch_wm = TWFetchWorkManager{
		Shutdown:        0,
		ShutdownChannel: make(chan string),
	}

	// Start the work timer routine.
	// When TWFetchWorkManager returns the program terminates.
	go Tw_fetch_wm.GoRoutineworkTimer()

	helper.WriteStdout("main", "TWFetchStartup", "Completed")
	return err
}

// TWFetchShutdown brings down the manager gracefully.
func TWFetchShutdown() error {
	var err error
	defer helper.CatchPanic(&err, "main", "TWFetchShutdown")

	helper.WriteStdout("main", "TWFetchShutdown", "Started")

	// Shutdown the program
	helper.WriteStdout("main", "TWFetchShutdown", "Info : Shutting Down")
	atomic.CompareAndSwapInt32(&Tw_fetch_wm.Shutdown, 0, 1)

	helper.WriteStdout("main", "TWFetchShutdown", "Info : Shutting Down Work Timer")
	Tw_fetch_wm.ShutdownChannel <- "Down"
	<-Tw_fetch_wm.ShutdownChannel

	close(Tw_fetch_wm.ShutdownChannel)

	helper.WriteStdout("main", "TWFetchShutdown", "Completed")
	return err
}

// GoRoutineworkTimer perform the work on the defined interval.
func (TWFetchWorkManager *TWFetchWorkManager) GoRoutineworkTimer() {
	helper.WriteStdout("WorkTimer", "TWFetchWorkManager.GoRoutineworkTimer", "Started")

	wait := timerPeriod

	for {
		helper.WriteStdoutf("WorkTimer", "TWFetchWorkManager.GoRoutineworkTimer", "Info : Wait To Run : Seconds[%.0f]", wait.Seconds())

		select {
		case <-TWFetchWorkManager.ShutdownChannel:
			helper.WriteStdoutf("WorkTimer", "TWFetchWorkManager.GoRoutineworkTimer", "Shutting Down")
			TWFetchWorkManager.ShutdownChannel <- "Down"
			return

		case <-time.After(wait):
			helper.WriteStdoutf("WorkTimer", "TWFetchWorkManager.GoRoutineworkTimer", "Woke Up")
			break
		}

		// Mark the starting time
		startTime := time.Now()

		// Perform the work
		TWFetchWorkManager.PerformTheWork()

		// Mark the ending time
		endTime := time.Now()

		// Calculate the amount of time to wait to start TWFetchWorkManager again.
		duration := endTime.Sub(startTime)
		wait = timerPeriod - duration
	}
}

// PerformTheWork with sleep times.
func (TWFetchWorkManager *TWFetchWorkManager) PerformTheWork() {
	defer helper.CatchPanic(nil, "TWFetchWorkManager", "TWFetchWorkManager.PerformTheWork")
	helper.WriteStdout("WorkTimer", "TWFetchWorkManager.GoRoutineworkTimer", "Started")

	//run get_tweets.go
	xpnk_twitter.Get_tweets()
	
	helper.WriteStdout("WorkTimer", "TWFetchWorkManager.GoRoutineworkTimer", "Completed")
}
