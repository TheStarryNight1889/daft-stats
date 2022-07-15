package db

import (
	"context"
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Connect() *mongo.Database {

	mongostring := os.Getenv("MONGO_URI")
	clientOpts := options.Client().ApplyURI(mongostring)
	client, err := mongo.Connect(context.TODO(), clientOpts)
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected and pinged.")
	coll := client.Database("daft-stats")
	return coll
}
