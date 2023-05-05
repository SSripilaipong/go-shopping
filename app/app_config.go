package app

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	appLambda "go-shopping/app/lambda"
)

type Config struct {
	Inventory *InventoryConfig
}

func newConfigFromEnvironment() *Config {
	return &Config{
		Inventory: newInventoryConfigFromEnvironment(),
	}
}

func newDependency(config *Config) appLambda.Dependency {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	dynamodbClient := dynamodb.New(sess)

	return appLambda.Dependency{
		Inventory: newInventoryDependency(dynamodbClient, config.Inventory),
	}
}
