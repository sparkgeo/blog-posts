// Mehul Solanki
// solanki@unbc.ca, ms280690@gmail.com
// SparkGeo Test 

package main

import (
	"fmt"
	"net/http"
	"strconv"
 )


// The catcher queue must be implemented through channels in go, since subroutines(catching and dispatching) can communicate and synchronize over it.
// But go channels need to be typed, so we need to define a structure of messages that the channel will work on.

// Channel Structure.
// We can add other fields later.
type RequestStructure struct{
// In the curl command the following data will be sent,

// The current user running the command.  
	userName 			string

// To send multiple request, a for loop will be used so the value of the variable(i). 
	requestNumber 		int

// The current unix timestamp in unix.
	requestTimeStamp	string			
}

//Creating new channel of type RequestStructure with buffer size 1000.
var catcherQueue = make(chan RequestStructure, 1000)

// The Catcher function, prints out the request for testing purposes. 
// The required fields are extracted from the request and converted to the appropriate format.   
func catcherFunction(rs http.ResponseWriter, rq *http.Request) {

/*
//Just Printing out the request parameter by parameter for testing purposes.
	fmt.Println("HTTPRequest Mehtod",				rq.Method)
	fmt.Println("HTTPRequest Proto",				rq.Proto)
	fmt.Println("HTTPRequest ProtoMajor",			rq.ProtoMajor)
	fmt.Println("HTTPRequest ProtoMinor",			rq.ProtoMinor)
	fmt.Println("HTTPRequest Header", 				rq.Header)
	fmt.Println("HTTPRequest Body", 				rq.Body)
	fmt.Println("HTTPRequest ContentLength", 		rq.ContentLength)
	fmt.Println("HTTPRequest TransferEncoding", 	rq.TransferEncoding)
	fmt.Println("HTTPRequest Close", 				rq.Close)
	fmt.Println("HTTPRequest Host", 				rq.Host)
	fmt.Println("HTTPRequest Form", 				rq.Form)
	fmt.Println("HTTPRequest PostForm", 			rq.PostForm)
	fmt.Println("HTTPRequest MultipartForm", 		rq.MultipartForm)
	fmt.Println("HTTPRequest Trailer", 				rq.Trailer)
	fmt.Println("HTTPRequest RemoteAddr", 			rq.RemoteAddr)
	fmt.Println("HTTPRequest RequestURI", 			rq.RequestURI)
	fmt.Println("HTTPRequest TLS", 					rq.TLS)
	fmt.Println()
*/

// For creating a work request we need a userName and requestNumber.  
	userName 				:= 	rq.FormValue("userName")

// Since the return type is string we need to convert it to an int.	
	requestNumberString 	:= 	rq.FormValue("requestNumber")
	requestNumber, _		:=	strconv.Atoi(requestNumberString)

// Similarly for the time stamp of the request.
	requestTimeStampString	:=	rq.FormValue("requestTimeStamp")
// The string is parsed into int64 since the time.Unix function requires it. 
//	requestTimeStampInt, _	:= 	strconv.ParseInt(requestTimeStampString, 10, 64)
//	requestTimeStamp		:= 	time.Unix(requestTimeStampInt, 0)

// Creating the RequestStructure to put on the channel.
	workRequest := RequestStructure{ userName: userName, requestNumber: requestNumber, requestTimeStamp: requestTimeStampString}
	
// Printing the request.
	fmt.Println("userName", userName)
	fmt.Println("requestNumber", requestNumberString)
	fmt.Println("requestTimeStamp", requestTimeStampString)

// Put it on the queue.
	catcherQueue <- workRequest
	fmt.Println("From Catcher : Work request queued")

// Creating a status.
	rs.WriteHeader(http.StatusCreated)
	fmt.Println("From Catcher : StatusCreated")
	fmt.Println()
	return
}

/*
func main() {

	http.HandleFunc("/work", catcherFunction)
	http.ListenAndServe("127.0.0.1:1025", nil)
}
*/