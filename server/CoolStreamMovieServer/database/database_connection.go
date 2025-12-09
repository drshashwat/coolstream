package database

import (
	"log"
	"log/slog"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func DBInstance() *mongo.Client {
	err := godotenv.Load(".env")
	if err != nil {
		slog.Warn("unable to find .env file", err)
	}

	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		log.Fatal("MONGODB_URI is not set")
	}
	slog.Info("MongoDB URI:", uri)

	clientOptions := options.Client().ApplyURI(uri)

	client, err := mongo.Connect(clientOptions)
	if err != nil {
		return nil
	}
	return client
}

var Client *mongo.Client = DBInstance()

func OpenCollection(collectioName string) *mongo.Collection {
	err := godotenv.Load(".env")
	if err != nil {
		slog.Warn("unable to find .env file")
	}

	databaseName := os.Getenv("DATABASE_NAME")
	if databaseName == "" {
		slog.Error("DATABASE_NAME is not set")
		return nil
	}
	slog.Info("DATABASE_NAME: ", databaseName)

	collection := Client.Database(databaseName).Collection(collectioName)
	if collection == nil {
		return nil
	}
	return collection
}
