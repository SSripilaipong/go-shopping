package ddbrepo

import "github.com/aws/aws-sdk-go/service/dynamodb"

type DynamodbClient interface {
	PutItem(input *dynamodb.PutItemInput) (*dynamodb.PutItemOutput, error)
}
