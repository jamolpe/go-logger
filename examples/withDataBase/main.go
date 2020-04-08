package main

import (
	"context"
	"fmt"
	logger "gologger"
	"gologger/pkg/models"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type repository struct {
	client        *mongo.Client
	database      *mongo.Database
	logCollection *mongo.Collection
}

func configureAndConnect() *mongo.Client {
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

func new(client *mongo.Client) models.Repository {
	database := client.Database("Logger")
	logCollection := database.Collection("logs")
	return &repository{logCollection: logCollection}
}

func (repo *repository) SaveLog(log models.LogModel) error {
	_, err := repo.logCollection.InsertOne(context.TODO(), log)
	return err
}

func main() {
	client := configureAndConnect()
	repo := new(client)
	config := models.Configuration{DisplayLogs: true, SaveLogs: true,
		LogLevels: models.DisplayConfiguration{
			DisplayDebug:    true,
			DisplayWarnings: true,
			DisplayError:    true,
			DisplayInfo:     true,
		},
		Repository: repo,
	}
	logger := logger.New(config)
	logger.DEBUG("I'm debug")
	logger.ERROR("I'm error")
	logger.INFO("I'm info")
	logger.WARNING("I'm waring")
}
