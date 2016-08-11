// Xapnik gratefully uses this code Copyright 2013 Ardan Studios. All rights reserved.
// Use of TWFetchWorkManager source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package workmanager implements the WorkManager singleton. This manager
// controls the starting, shutdown and processing of work.

package xpnk_instagram_fetch_workmanager

//This workmanager manages Posts-Manager.go

import (
	"xpnk_instagram/posts-manager"
	"xpnk_instagram/xpnk-instagram-manager/helper"
	"sync/atomic"
	"time"
)

const (
	timerPeriod time.Duration = 600 * time.Second // Interval to wake up on.
)

// IGFetchWorkManager is responsible for starting and shutting down the program.
type IGFetchWorkManager struct {
	Shutdown        int32
	ShutdownChannel chan string
}

var IG_fetch_wm IGFetchWorkManager // Reference to the singleton.

// IGFetchStartup brings the manager to a running state.
func IGFetchStartup() error {
	var err error
	defer helper.CatchPanic(&err, "main", "IGFetchStartup")

	helper.WriteStdout("main", "IGFetchStartup", "Started")

	// Create the work manager to get the program going
	IG_fetch_wm = IGFetchWorkManager{
		Shutdown:        0,
		ShutdownChannel: make(chan string),
	}

	// Start the work timer routine.
	// When IGFetchWorkManager returns the program terminates.
	go IG_fetch_wm.GoRoutineworkTimer()

	helper.WriteStdout("main", "IGFetchStartup", "Completed")
	return err
}

// IGFetchShutdown brings down the manager gracefully.
func IGFetchShutdown() error {
	var err error
	defer helper.CatchPanic(&err, "main", "IGFetchShutdown")

	helper.WriteStdout("main", "IGFetchShutdown", "Started")

	// Shutdown the program
	helper.WriteStdout("main", "IGFetchShutdown", "Info : Shutting Down")
	atomic.CompareAndSwapInt32(&IG_fetch_wm.Shutdown, 0, 1)

	helper.WriteStdout("main", "IGFetchShutdown", "Info : Shutting Down Work Timer")
	IG_fetch_wm.ShutdownChannel <- "Down"
	<-IG_fetch_wm.ShutdownChannel

	close(IG_fetch_wm.ShutdownChannel)

	helper.WriteStdout("main", "IGFetchShutdown", "Completed")
	return err
}

// GoRoutineworkTimer perform the work on the defined interval.
func (IGFetchWorkManager *IGFetchWorkManager) GoRoutineworkTimer() {
	helper.WriteStdout("WorkTimer", "IGFetchWorkManager.GoRoutineworkTimer", "Started")

	wait := timerPeriod

	for {
		helper.WriteStdoutf("WorkTimer", "IGFetchWorkManager.GoRoutineworkTimer", "Info : Wait To Run : Seconds[%.0f]", wait.Seconds())

		select {
		case <-IGFetchWorkManager.ShutdownChannel:
			helper.WriteStdoutf("WorkTimer", "IGFetchWorkManager.GoRoutineworkTimer", "Shutting Down")
			IGFetchWorkManager.ShutdownChannel <- "Down"
			return

		case <-time.After(wait):
			helper.WriteStdoutf("WorkTimer", "IGFetchWorkManager.GoRoutineworkTimer", "Woke Up")
			break
		}

		// Mark the starting time
		startTime := time.Now()

		// Perform the work
		IGFetchWorkManager.PerformTheWork()

		// Mark the ending time
		endTime := time.Now()

		// Calculate the amount of time to wait to start IGFetchWorkManager again.
		duration := endTime.Sub(startTime)
		wait = timerPeriod - duration
	}
}

// PerformTheWork with sleep times.
func (IGFetchWorkManager *IGFetchWorkManager) PerformTheWork() {
	defer helper.CatchPanic(nil, "IGFetchWorkManager", "IGFetchWorkManager.PerformTheWork")
	helper.WriteStdout("WorkTimer", "IGFetchWorkManager.GoRoutineworkTimer", "Started")

	//run posts-manager.go
	posts_manager.Get_posts()
	
	helper.WriteStdout("WorkTimer", "IGFetchWorkManager.GoRoutineworkTimer", "Completed")
}
