package main

import (
	"context"
	"fmt"
	logger "gologger"
	"gologger/pkg/models"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConfigureAndConnect() *mongo.Client {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		panic("err")
	}
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		panic("error")
	}

	fmt.Println("Connected to MongoDB!")
	return client
}

func main() {
	config := models.Configuration{DisplayLogs: true, SaveLogs: false,
		LogLevels: models.DisplayConfiguration{
			DisplayDebug:    true,
			DisplayWarnings: true,
			DisplayError:    true,
			DisplayInfo:     true,
		},
	}
	logger := logger.New(config)
	logger.DEBUG("I'm debug")
	logger.ERROR("I'm error")
	logger.INFO("I'm info")
	logger.WARNING("I'm waring")
}
