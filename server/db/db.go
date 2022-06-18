package db

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Connection URI
const uri = "mongodb://localhost:27017"

func Connect() *mongo.Database {
	credential := options.Credential{
		Username: "root",
		Password: "example",
	}
	clientOpts := options.Client().ApplyURI(uri).SetAuth(credential)
	client, err := mongo.Connect(context.TODO(), clientOpts)
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected and pinged.")
	coll := client.Database("daft-stats")
	return coll
}
