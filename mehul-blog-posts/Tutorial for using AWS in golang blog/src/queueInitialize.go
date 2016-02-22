package main

import (
    "fmt"
    "github.com/goamz/goamz/sqs"
)

func queue_init() sqs.Queue {

    var access_key  = // Your AWS access key
    var secret_key  = // Your AWS secret key
    var queue_url   = // URL found under details on AWS console

    // Region in which queue exists
    var region_name = us.west.2 // For Oregon

    var queue_name = //Your queue name

    conn := establish_connection(access_key, secret_key, region_name)

    q, _ := conn.GetQueue(queue_name)


    fmt.Println("queue does exist")
    q := conn.QueueFromArn(queue_url)

    return *q
}

   func establish_connection(access_key string, secret_key   string, region_name string) *sqs.SQS {

        // Creating a new sqs client.
        conn, cerr := sqs.NewFrom(access_key, secret_key, region_name)
        return conn
}
