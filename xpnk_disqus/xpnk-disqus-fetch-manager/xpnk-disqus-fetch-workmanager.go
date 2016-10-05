// Xapnik gratefully uses this code Copyright 2013 Ardan Studios. All rights reserved.
// Use of TWFetchWorkManager source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package workmanager implements the WorkManager singleton. This manager
// controls the starting, shutdown and processing of work.

package xpnk_disqus_fetch_workmanager

//This workmanager manages posts-manager.go

import (
	"xpnk_disqus/posts-manager"
	"xpnk-shared/manager-helper/helper"
	"sync/atomic"
	"time"
)

const (
	timerPeriod time.Duration = 6000 * time.Second // Interval to wake up on.
)

// DisqusFetchWorkManager is responsible for starting and shutting down the program.
type DisqusFetchWorkManager struct {
	Shutdown        int32
	ShutdownChannel chan string
}

var Disqus_fetch_wm DisqusFetchWorkManager // Reference to the singleton.

// DisqusFetchStartup brings the manager to a running state.
func DisqusFetchStartup() error {
	var err error
	defer helper.CatchPanic(&err, "main", "DisqusFetchStartup")

	helper.WriteStdout("main", "DisqusFetchStartup", "Started")

	// Create the work manager to get the program going
	Disqus_fetch_wm = DisqusFetchWorkManager{
		Shutdown:        0,
		ShutdownChannel: make(chan string),
	}

	// Start the work timer routine.
	// When DisqusFetchWorkManager returns the program terminates.
	go Disqus_fetch_wm.GoRoutineworkTimer()

	helper.WriteStdout("main", "DisqusFetchStartup", "Completed")
	return err
}

// DisqusFetchShutdown brings down the manager gracefully.
func DisqusFetchShutdown() error {
	var err error
	defer helper.CatchPanic(&err, "main", "DisqusFetchShutdown")

	helper.WriteStdout("main", "DisqusFetchShutdown", "Started")

	// Shutdown the program
	helper.WriteStdout("main", "DisqusFetchShutdown", "Info : Shutting Down")
	atomic.CompareAndSwapInt32(&Disqus_fetch_wm.Shutdown, 0, 1)

	helper.WriteStdout("main", "DisqusFetchShutdown", "Info : Shutting Down Work Timer")
	Disqus_fetch_wm.ShutdownChannel <- "Down"
	<-Disqus_fetch_wm.ShutdownChannel

	close(Disqus_fetch_wm.ShutdownChannel)

	helper.WriteStdout("main", "DisqusFetchShutdown", "Completed")
	return err
}

// GoRoutineworkTimer perform the work on the defined interval.
func (DisqusFetchWorkManager *DisqusFetchWorkManager) GoRoutineworkTimer() {
	helper.WriteStdout("WorkTimer", "DisqusFetchWorkManager.GoRoutineworkTimer", "Started")

	wait := timerPeriod

	for {
		helper.WriteStdoutf("WorkTimer", "DisqusFetchWorkManager.GoRoutineworkTimer", "Info : Wait To Run : Seconds[%.0f]", wait.Seconds())

		select {
		case <-DisqusFetchWorkManager.ShutdownChannel:
			helper.WriteStdoutf("WorkTimer", "DisqusFetchWorkManager.GoRoutineworkTimer", "Shutting Down")
			DisqusFetchWorkManager.ShutdownChannel <- "Down"
			return

		case <-time.After(wait):
			helper.WriteStdoutf("WorkTimer", "DisqusFetchWorkManager.GoRoutineworkTimer", "Woke Up")
			break
		}

		// Mark the starting time
		startTime := time.Now()

		// Perform the work
		DisqusFetchWorkManager.PerformTheWork()

		// Mark the ending time
		endTime := time.Now()

		// Calculate the amount of time to wait to start DisqusFetchWorkManager again.
		duration := endTime.Sub(startTime)
		wait = timerPeriod - duration
	}
}

// PerformTheWork with sleep times.
func (DisqusFetchWorkManager *DisqusFetchWorkManager) PerformTheWork() {
	defer helper.CatchPanic(nil, "DisqusFetchWorkManager", "DisqusFetchWorkManager.PerformTheWork")
	helper.WriteStdout("WorkTimer", "DisqusFetchWorkManager.GoRoutineworkTimer", "Started")

	//run posts-manager.go
	posts_manager.Get_posts()
	
	helper.WriteStdout("WorkTimer", "DisqusFetchWorkManager.GoRoutineworkTimer", "Completed")
}
