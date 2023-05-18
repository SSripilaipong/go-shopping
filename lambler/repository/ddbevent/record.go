package ddbevent

import (
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type Record struct {
	newImage map[string]*dynamodb.AttributeValue `json:"NewImage,omitempty"`
}
