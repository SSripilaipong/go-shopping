package app

import (
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"go-shopping/app/core/inventory/adapter"
	"go-shopping/app/core/inventory/controller/http"
	"go-shopping/app/core/inventory/domain"
	"go-shopping/app/core/inventory/usecase"
	"go-shopping/lambler/repository/dynamodbrepository"
	"os"
)

type InventoryConfig struct {
	ItemTableName string
}

func newInventoryConfigFromEnvironment() *InventoryConfig {
	return &InventoryConfig{
		ItemTableName: os.Getenv("INVENTORY_ITEM_TABLE_NAME"),
	}
}

func newInventoryDependency(dynamodbClient *dynamodb.DynamoDB, config *InventoryConfig) http.Dependency {
	return http.Dependency{
		CreateNewItemUsecase: usecase.CreateNewItem(
			dynamodbrepository.NewSingleKeyTable[*domain.Item](dynamodbClient, config.ItemTableName).Repository(),
			adapter.GenerateItemId,
		),
	}
}
