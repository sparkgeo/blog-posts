
package main

import (
    // "fmt"
    "net/http"
)

// SQS
var q = queue_init()

func main() {
    http_handler()
}

func http_handler() {
    http.HandleFunc("/work", handle_func)
    listen_and_serve()
}

func handle_func(rs http.ResponseWriter, rq *http.Request) {
    input_message_json, _ := messageCreation(rq)
    send_msg(input_message_json, &q)
    rs.WriteHeader(http.StatusOK)
}

func listen_and_serve() {
    http.ListenAndServe("127.0.0.1:7001", nil)
}
