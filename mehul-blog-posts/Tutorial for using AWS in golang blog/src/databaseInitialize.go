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
