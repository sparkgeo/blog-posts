package main

import (
	"fmt"
	"net/http"
)


func main() {

// Start the dispatcher.
	fmt.Println("From Main : Starting the dispatcher")
	initDispatcher(10)

// Register our catcherFunction as an HTTP handler function.
	fmt.Println("From Main : Registering the catcher")
	http.HandleFunc("/work", catcherFunction)

// ListenAndServe HTTP server!
	fmt.Println("From Main : HTTP server listening on", "127.0.0.1:1026")

// For Testing purposes.
	err := http.ListenAndServe("127.0.0.1:1026", nil)
	
	if err != nil {
		fmt.Println(err.Error())
	}
}