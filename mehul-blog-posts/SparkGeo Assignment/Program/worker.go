package main

import (
	"fmt"
)

//Next is to create something which will handle the requests, a Worker. 
type Worker struct {
	id          int 							// Worker id
	work        chan RequestStructure 			// The channel where the work is received.
	workerQueue chan chan RequestStructure 		// The channel where the workers reside. 
}

// Make function to create a new worker, just like make function for other types such as channels.
func makeWorker(id int, workerQueue chan chan RequestStructure) Worker {

// Create, and return the worker.
	worker := Worker{
		id:          id,
		work:        make(chan RequestStructure), 	// Each worker has it's own queue where the work is received.
		workerQueue: workerQueue, 					// Common, where all the workers are.
	}

	return worker
}

// This function initializes each worker.
func (w Worker) initWorker() { /*func_name() allows it to be called on a struct using the . operator*/
// Since we want the workers to be running and ready to receive work requests simultaneously, so it has to be a go routine 
//	go func() {

// This is to replicate the while(true) infinite loop structure.
		for {

// Put the worker in the worker queue.
			w.workerQueue <- w.work

// Switch case for channels
			select {

// Receive a work request.
			case work := <- w.work:

// Who is working? // Whose work is it?
				fmt.Printf("From Worker %d : Work Request received, at your service %s", w.id, work.userName)
				return
			}
		}
//	}()
}
