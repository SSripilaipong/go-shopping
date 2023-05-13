package lambler

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

var dynamodbClient *dynamodb.DynamoDB

func GetDynamodbClient() *dynamodb.DynamoDB {
	if dynamodbClient == nil {
		sess := session.Must(session.NewSessionWithOptions(session.Options{
			SharedConfigState: session.SharedConfigEnable,
		}))
		dynamodbClient = dynamodb.New(sess)
	}
	return dynamodbClient
}
