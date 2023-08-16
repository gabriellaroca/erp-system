package config

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoClient *mongo.Client

func SetMongoClient() (err error) {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI("mongodb+srv://gabrielb:afqFuKHOpQLHeKZK@cluster0.qdtrf.mongodb.net/?retryWrites=true&w=majority").SetServerAPIOptions(serverAPI)
	MongoClient, err = mongo.Connect(context.TODO(), opts)
	if err != nil {
		panic(err)
	}
	return
}
