package db

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Connect() *mongo.Database {
	mongostring := "" //os.Getenv("MONGO_URI")
	if mongostring == "" {
		mongostring = "mongodb://root:example@127.0.0.1/"
	}

	clientOpts := options.Client().ApplyURI(mongostring)
	client, err := mongo.Connect(context.TODO(), clientOpts)
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected and pinged.")
	coll := client.Database("daft-stats")
	return coll
}
