package ddbevent

import (
	"fmt"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

func NewImageData[T any](record *Record, target *T) error {
	err := dynamodbattribute.UnmarshalMap(record.newImage, target)
	if err != nil {
		panic(fmt.Errorf("unhandled error"))
	}
	fmt.Printf("%#v\n", target)
	return nil
}
