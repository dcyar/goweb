package config

import (
	"context"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
)

const Database string = "goweb"

func DatabaseConnect() *mongo.Client {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	mongoUri := os.Getenv("MONGODB_URI")

	if mongoUri == "" {
		log.Fatal("You must set your 'MONGODB_URI' environmental variable")
	}

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(mongoUri))

	if err != nil {
		panic(err)
	}

	return client
}

func DatabaseClose(client *mongo.Client) {
	if err := client.Disconnect(context.TODO()); err != nil {
		panic(err)
	}
}
