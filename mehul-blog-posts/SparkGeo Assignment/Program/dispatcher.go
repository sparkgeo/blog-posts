
package main

import (
	"fmt"
)
// This function initializes the workers.
func initDispatcher(workerCount int) { // number of workers

// Initializing the common channel where the workers will wait for the work. 	
workerQueue := make(chan chan RequestStructure, workerCount)

// Depending on the count, creating the required number of workers.
	for i := 0; i < workerCount; i++ {
		fmt.Println("From Dispatcher : Starting Worker", i+1)
		worker := makeWorker(i+1, workerQueue)
		go worker.initWorker()
	}

// Outer Anonymous Go routine.
	go func() {
		for {
			select {
// Pull off work request from the queue
			case work := <-catcherQueue:
				fmt.Println("From Dispatcher : Received work requeust")
// Inner Anonymous Go routine.
				go func() {
					worker := <-workerQueue
					fmt.Println("From Dispatcher : Dispatching work request")
// Send the work request to worker.
					worker <- work
				}()
			}
		}
	}()
}
