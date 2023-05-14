package create

import (
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/stretchr/testify/assert"
	gfun "go-shopping/go/gfunc"
	"go-shopping/lambler/repository/dynamodbrepository"
	"go-shopping/tests/lambler/repository/dynamodbrepository/mock"
	"testing"
)

func Test_should_put_marshalled_item(t *testing.T) {
	type Entity struct {
		Id string
		A  string
		B  int
	}
	client := &mock.DynamodbClient{}
	repo := dynamodbrepository.NewSingleKeyTable[*Entity](client, "my-table").Repository()
	_ = repo.Create(func() *Entity {
		return &Entity{
			Id: "1234",
			A:  "xxx",
			B:  9999,
		}
	})

	assert.Equal(t, &dynamodb.PutItemInput{
		Item: map[string]*dynamodb.AttributeValue{
			"Id": {S: gfun.Pointer("1234")},
			"A":  {S: gfun.Pointer("xxx")},
			"B":  {N: gfun.Pointer("9999")},
		},
		TableName: gfun.Pointer("my-table"),
	}, client.PutItem_input)
}
