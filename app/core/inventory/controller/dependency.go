package controller

import (
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"go-shopping/app/core/inventory/adapter"
	"go-shopping/app/core/inventory/controller/http"
	"go-shopping/app/core/inventory/domain"
	"go-shopping/app/core/inventory/usecase"
	"go-shopping/lambler/repository/dynamodbrepository"
)

func NewDependency(dynamodbClient *dynamodb.DynamoDB, config *Config) http.Dependency {
	return http.Dependency{
		CreateNewItemUsecase: usecase.CreateNewItem(
			dynamodbrepository.NewSingleKeyTable[*domain.Item](dynamodbClient, config.ItemTableName).Repository(),
			adapter.GenerateItemId,
		),
	}
}
