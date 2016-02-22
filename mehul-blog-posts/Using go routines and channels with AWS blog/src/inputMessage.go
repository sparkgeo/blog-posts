package main

import (
    "encoding/json"
    "net/http"
)

type InputMessage struct {
    User_Name           string `json:"userName"`
    Request_Number      string `json:"requestNumber"`
    Request_Timestamp   string `json:"requestTimestamp"`
}

//Integration test only
func messageCreation(rq *http.Request) (string, bool) {

    // For creating a work request we need a userName and requestNumber.
    userName                :=  rq.FormValue("userName")

    // Since the return type is string we need to convert it to an int.
    requestNumberString     :=  rq.FormValue("requestNumber")

    // Similarly for the time stamp of the request.
    requestTimeStampString  :=  rq.FormValue("requestTimeStamp")

    input_message := &InputMessage{
        User_Name: userName,
        Request_Number: requestNumberString,
        Request_Timestamp: requestTimeStampString}

    // Convert to JSON string
    json_msg, msg_creation := create_json_msg(input_message)
    msg := string(json_msg)

    return msg, msg_creation
}

func create_json_msg(input_message *InputMessage) ([]byte, bool) {
    json_msg, err := json.Marshal(input_message)

    if err != nil { // Wait and Retry to be added
        //  fmt.Println(err.Error())
        //create_json_msg(input_message)
        return json_msg, false
    }
    return json_msg, true
}
