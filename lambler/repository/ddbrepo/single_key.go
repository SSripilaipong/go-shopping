package ddbrepo

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"go-shopping/lambler"
	"go-shopping/lambler/repository"

	"log"
)

type singleKeyRepository[E any] struct {
	client    DynamodbClient
	tableName string
}

func (r *singleKeyRepository[E]) Create(factory func() E) error {
	entityToCreate := factory()
	marshalledEntity, err := dynamodbattribute.MarshalMap(entityToCreate)
	if err != nil {
		log.Fatalf("Got error marshalling new movie item: %s", err)
	}
	_, err = r.client.PutItem(&dynamodb.PutItemInput{
		Item:      marshalledEntity,
		TableName: aws.String(r.tableName),
	})
	if err != nil {
		log.Fatalf("Got error calling PutItem: %s", err)
	}
	return nil
}

type singleKeyTable[E any] struct {
	repository repository.SingleKeyRepository[E]
}

func NewSingleKeyTableWithClient[E any](client DynamodbClient, tableName string) repository.SingleKeyTable[E] {
	return &singleKeyTable[E]{
		repository: &singleKeyRepository[E]{
			tableName: tableName,
			client:    client,
		},
	}
}

func NewSingleKeyTable[E any](tableName string) repository.SingleKeyTable[E] {
	return NewSingleKeyTableWithClient[E](lambler.GetDynamodbClient(), tableName)
}

func (t *singleKeyTable[E]) Repository() repository.SingleKeyRepository[E] {
	return t.repository
}
