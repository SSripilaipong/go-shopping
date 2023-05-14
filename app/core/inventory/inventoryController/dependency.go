package inventoryController

import (
	"go-shopping/app/core/inventory/adapter"
	"go-shopping/app/core/inventory/domain"
	"go-shopping/app/core/inventory/inventoryController/inventoryHttp"
	"go-shopping/app/core/inventory/usecase"
	"go-shopping/lambler/repository/dynamodbrepository"
)

type Dependency struct {
	http inventoryHttp.Dependency
}

func NewDependency(createNewItemUsecase usecase.CreateNewItemFunc) Dependency {
	return Dependency{
		http: inventoryHttp.Dependency{
			CreateNewItemUsecase: createNewItemUsecase,
		},
	}
}

func NewRuntimeDependency(config *Config) Dependency {
	return NewDependency(
		usecase.CreateNewItem(
			dynamodbrepository.NewSingleKeyTable[*domain.Item](config.ItemTableName).Repository(),
			adapter.GenerateItemId,
		),
	)
}
