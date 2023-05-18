package ddbevent

import (
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type dynamoDBEvent struct {
	Records []dynamoDBEventRecord `json:"Records"`
}

type dynamoDBEventRecord struct {
	Change dynamoDBStreamRecord `json:"dynamodb"`
}

type dynamoDBStreamRecord struct {
	NewImage map[string]*dynamodb.AttributeValue `json:"NewImage,omitempty"`
}
