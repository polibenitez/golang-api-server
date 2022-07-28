package db

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/fatih/color"
	"github.com/joho/godotenv"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

var _ = godotenv.Load(".env")

// Dbconnect -> connects mongo
func Dbconnect() *mongo.Client {
	fmt.Println(os.Getenv("MONGO_URL"))
	clientOptions := options.Client().ApplyURI(os.Getenv("MONGO_URL"))
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal("⛒ Connection Failed to Mongo Database")
		log.Fatal(err)
	}
	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal("⛒ Connection Failed to Mongo Database")
		log.Fatal(err)
	}
	color.Green("⛁ Connected to Database")
	return client
}
