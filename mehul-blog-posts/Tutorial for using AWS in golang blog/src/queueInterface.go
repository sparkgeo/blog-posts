import (
        "github.com/goamz/goamz/sqs"
    )

    //send msg
    func send_msg(msg string, q *sqs.Queue) sqs.SendMessageResponse {
        res, _ := q.SendMessage(msg)

        return *res
    }

    //recieve msg
    func get_msg(q *sqs.Queue) sqs.ReceiveMessageResponse {
        recieveMessageResponse, err := q.ReceiveMessage(1)

        return *recieveMessageResponse
    }

    //delete msg
    func delete_msg(receiveMessageResposnse *sqs.ReceiveMessageResponse, q *sqs.Queue) sqs.DeleteMessageResponse {


    deleteMessageResposne, err := q.DeleteMessageUsingReceiptHandle(receiveMessageResposnse.Messages[0].ReceiptHandle)

        return *deleteMessageResposne
    }
