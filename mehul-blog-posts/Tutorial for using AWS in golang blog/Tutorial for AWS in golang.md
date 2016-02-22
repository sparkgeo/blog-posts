# Tutorial for using AWS in golang 

This article discusses about using AWS in golang using the goamz library. 

The official sdk can be found [here](https://github.com/aws/aws-sdk-go/wiki); bear in mind that it is NOT used for this tutorial since it was in beta at the time of writing.

This is a fairly light-weight, beginner level tutorial to get you up and running with AWS in golang. Though using a third party library, we believe the basics could be translated to the official one. 

We will attempt to provide a starting point for using [Amazon Simple Queue Service (SQS)](https://aws.amazon.com/sqs/) and [Amazon DynamoDB](https://aws.amazon.com/dynamodb/?sc_channel=PS&sc_campaign=acquisition_CA&sc_publisher=google&sc_medium=dynamodb_hv_b&sc_content=dynamodb_e&sc_detail=aws%20dynamodb&sc_category=dynamodb&sc_segment=73324893176&sc_matchtype=e&sc_country=CA&s_kwcid=AL!4422!3!73324893176!e!!g!!aws%20dynamodb&ef_id=Vc563AAABcy2IVRA:20160108221108:s). Retry,silent failure, error handling mechanisms have not been discussed in this tutorial, but from our experience can be easily added. 


The following components are necessary for this tutorial:

1. [AWS account](https://aws.amazon.com/free/?sc_channel=PS&sc_campaign=acquisition_CA&sc_publisher=google&sc_medium=cloud_computing_b&sc_content=aws_core_e&sc_detail=a.w.s&sc_category=cloud_computing&sc_segment=84661170416&sc_matchtype=e&sc_country=CA&s_kwcid=AL!4422!3!84661170416!e!!g!!a.w.s&ef_id=Vc563AAABcy2IVRA:20160108221217:s)
2. [golang](https://golang.org/).
3. [goamz](https://github.com/goamz/goamz)  
4. Terminal

The optional components are:

1. [Sublime Text Editor](http://www.sublimetext.com/)
2. [Package Control](https://packagecontrol.io/installation)
3. [GoSublime](https://packagecontrol.io/packages/GoSublime)

## SQS

### Queue Setup
Let's begin with the queue setup;

1. [Creating a new queue using AWS console](http://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSGettingStartedGuide/CreatingQueue.html)

2. Import the library;

	```go
	import (
		"github.com/goamz/goamz/sqs"
	)
	```
	
3. Storing required AWS credentials;
	
	```go
	var access_key 	= // Your AWS access key 
	var secret_key 	= // Your AWS secret key
	var queue_url 	= // URL found under details on AWS console
	```

4. The region name is located on the top right of the AWS console;
	
	```go
	// Region in which queue exists
	var region_name = us.west.2 // For Oregon
	```
	More information can be found [here](http://docs.aws.amazon.com/general/latest/gr/rande.html).

5.	Your queue name;
	
	```go
	var queue_name = //queue_name
	```

6. Establishing the connection to your AWS account in a go program;
	
	```go
	func establish_connection(access_key string, secret_key string, region_name string) *sqs.SQS {
	
		// Creating a new sqs client.
		conn, cerr := sqs.NewFrom(access_key, secret_key, region_name)	
		return conn
	}
	``` 
	The function returns a pointer to a queue object. 
 
7. 	Putting it all together into ```queueInitialize.go```;

	```go
	package main
	
	import (
		"fmt"
		"github.com/goamz/goamz/sqs"
	)
	
	func queue_init() sqs.Queue {
	
		var access_key 	= // Your AWS access key 
		var secret_key 	= // Your AWS secret key
		var queue_url 	= // URL found under details on AWS console
		
		// Region in which queue exists
		var region_name = us.west.2 // For Oregon
		
		var queue_name = //Your queue name
		
		conn := establish_connection(access_key, secret_key, region_name)
	
		q, _ := conn.GetQueue(queue_name)
	
		
		fmt.Println("queue does exist")
		q := conn.QueueFromArn(queue_url)
	
		return *q
	}
	
	func establish_connection(access_key string, secret_key string, region_name string) *sqs.SQS {
		
			// Creating a new sqs client.
			conn, cerr := sqs.NewFrom(access_key, secret_key, region_name)	
			return conn
	}
		
	```
	
### Queue Interface
After the queue is setup and ready, we need to use it. This section discusses the implementation of the basic queue operations such as sending, receiving and deleting messages. To know more about SQS please read, the [developer guide](http://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/Welcome.html).

1. Import the library;

	```go
	import (
		"github.com/goamz/goamz/sqs"
	)
	```

2. Sending a string message to the queue;
	
	```go
	//send msg
	func send_msg(msg string, q *sqs.Queue) sqs.SendMessageResponse {
		res, _ := q.SendMessage(msg)
	
		return *res
	}
	``` 

3. Receiving a single message;
	
	```go
	//recieve msg
	func get_msg(q *sqs.Queue) sqs.ReceiveMessageResponse {
		recieveMessageResponse, err := q.ReceiveMessage(1)
	
		return *recieveMessageResponse
	}
	```
	The response will be used at a later point in time.

4. Deleting a message;

	```go
	//delete msg
	func delete_msg(receiveMessageResposnse *sqs.ReceiveMessageResponse, q *sqs.Queue) sqs.DeleteMessageResponse {
	
	
	deleteMessageResposne, err := q.DeleteMessageUsingReceiptHandle(receiveMessageResposnse.Messages[0].ReceiptHandle)
		
		return *deleteMessageResposne
	}
	```	
	The function uses the ReceiveMessage response returned after receiving a message from the queue. The ReceiveMessageResponse has an attribute called a ReceiptHandle for each message which can be used to delete that message.

5. Putting it all together in ```queueInterface.go```;
	
	```go
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
	```
	

### Testing
1. The main function;

	```go
	func main() {
		q := queue_init()
		fmt.Println(q)
	
		send_msg_response := send_msg("hello", &q)
		fmt.Println(send_msg_response)
	
		receive_msg_response := get_msg(&q)
		fmt.Println(receive_msg_response.Messages[0].Body)
	
		delete_msg_response := delete_msg(&receive_msg_response, &q)
		fmt.Println(delete_msg_response)
	}
	```

2. Traverse into the folder where the program files are located and run this command in a terminal:

	```go
	go run queueInitialize.go queueInterface.go
	```

3. The output will be of the form;

	```go
	{9aa9999a9999 https://sqs.us-west-2.amazonaws.com/999999999999/my_queue_name}
	{9a99999aaa9a9a99a9999a999999a999  aa99aaa9-a9a9-99aa-a99a-99a99aa999a9 {99999aa9-a999-99aa-9a99-99aaa9999aa9 9}}
	hello
	{{99999aaa-a9aa-9999-99aa-99a9a9aaa99a 9}}
	```	

## DynamoDB
To know more about DynamoDB please read the [developer guide](http://docs.aws.amazon.com/amazondynamodb/latest/developerguide/Introduction.html).

### Database Setup
1. [Creating a table using AWS console](http://docs.aws.amazon.com/amazondynamodb/latest/developerguide/ConsoleDynamoDB.html).

2. Import the library;

	```go
	import (
		"github.com/goamz/goamz/aws"
		"github.com/goamz/goamz/dynamodb"
	)
	```

3. Initializing connection;

	```go
	func database_init() *dynamodb.Server {
	
		auth := aws_auth()
	
		region := aws.USWest2
	
		ddbs := NewFrom(auth, region)
	
		return ddbs
	}
	
	func NewFrom(auth aws.Auth, region aws.Region) *dynamodb.Server {
		return &dynamodb.Server{auth, region}
	}
	
	func aws_auth() aws.Auth {
		auth, err := aws.EnvAuth()
		return auth
	}
	```
For SQS, the connection was made manually within the program where the credentials are stored. In the above example, the credentials are stored as environment variables and the function ```aws.EnvAuth()``` is used to establish a connection.
For storing and accessing environment variables in golang please read the documentation on the [os package](https://golang.org/pkg/os/) along with this [example](https://gobyexample.com/environment-variables).

4. Putting it all together in ```databaseInitialize.go```;

	```go
	import (
			"github.com/goamz/goamz/aws"
			"github.com/goamz/goamz/dynamodb"
		)
	
	func database_init() *dynamodb.Server {
		
			auth := aws_auth()
		
			region := aws.USWest2
		
			ddbs := NewFrom(auth, region)
		
			return ddbs
		}
		
		func NewFrom(auth aws.Auth, region aws.Region) *dynamodb.Server {
			return &dynamodb.Server{auth, region}
		}
		
		func aws_auth() aws.Auth {
			auth, err := aws.EnvAuth()
			return auth
		}
	
	```

### Database Interface
1. Import the library;

	```go
	import (
		"github.com/goamz/goamz/dynamodb"
	)
	```
	
2. Before we get into the implementation of the database operations; let's take a look at the helper functions which are also useful for other tasks beyond the interfacing.
	
	2.1. Table Descriptor

	```go
	// Get table description
	func get_table_description(table *dynamodb.Table) *dynamodb.TableDescriptionT {
	table_description, td_err := table.DescribeTable()
		
	return table_description
	}
	```  
	
	2.2. DynamoDB Table

	```go
	// Get a table object to perform table.Operation(param1, ... param n)
	func get_table(table_name string, ddbs *dynamodb.Server) *dynamodb.Table {
		table_descriptor, td_err := ddbs.DescribeTable(table_name)
	
		primary_key := build_primary_key(table_descriptor)
	
		table := ddbs.NewTable(table_name, primary_key)
	
		return table
	}
	```
	
	2.3. DynamoDB Attributes
	For other non key attributes in a tuple.
	
	```go
	// Create Dynamo DB Attribute list
	func attribute_list_creator(attribute string) []dynamodb.Attribute {
		var attribute_list = make([]dynamodb.Attribute, 1)
	
		// cannot have empty attribute list
		attribute := &dynamodb.Attribute{
			Type:      "S",
			Name:      "attribute_name",
			Value:     attribute,
			SetValues: make([]string, 0),
			Exists:    ""}
	
		attribute_list[0] = *attribute
	
		return attribute_list
	}
	```
	
	2.4. DynamoDB Key
	
	```go
	// Create Dynamo DB Key
	func key_creator(primary_hash_key string) *dynamodb.Key {
		key := &dynamodb.Key{
			HashKey: msid}
	
		return key
	}
	
	``` 

3. Writing to the database;

	```go
	// Write to database.
	func database_writer(primary_hash_key string, primary_range_key string, attribute_list []dynamodb.Attribute, table dynamodb.Table) bool {
		bool, _ := table.PutItem(primary_hash_key, primary_range_key, attribute_list) // Overwrites
			
		return bool
	}
	```
	The ```get_table``` function returns the table object used for performing operations.
	
4. Reading from the database;

	```go
	// Read from Database (GetItem(key(primary_hash_key,primary_range_key)))
	// Can be used for checking if record exists or not by the length of the attribute_list
	// Consistent always returns the last updated value
	func database_reader(table *dynamodb.Table, key *dynamodb.Key) (map[string]*dynamodb.Attribute, int) {
	
		attribute_map, _ := table.GetItemConsistent(key, true)
	
		return attribute_map, len(attribute_map)
	}
	```
The ```key_creator``` function is used to generate a ```key *dynamodb.Key``` for querying on the table to retrieve the record(s).

5. Deleting from the database has been left as an exercise for the reader.


### Testing 
Left as an exercise for the reader.
	