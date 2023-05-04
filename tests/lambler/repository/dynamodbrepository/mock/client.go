package mock

import (
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type DynamodbClient struct {
	PutItem_input *dynamodb.PutItemInput
}

func (c *DynamodbClient) PutItem(input *dynamodb.PutItemInput) (*dynamodb.PutItemOutput, error) {
	c.PutItem_input = input
	return nil, nil
}
