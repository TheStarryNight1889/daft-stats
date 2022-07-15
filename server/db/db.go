package db

import (
	"context"
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Connect() *mongo.Database {

	//  if env DB_PORT, DB_HOST are set, use them
	//  otherwise use default values
	port := os.Getenv("DB_PORT")
	if port == "" {
		port = "27017"
	}
	host := os.Getenv("DB_HOST")
	if host == "" {
		host = "mongodb://localhost"
	}
	// build connection URI
	uri := fmt.Sprintf("%s:%s", host, port)

	// if env DB_USER, DB_PASS are set, use them
	// otherwise use default values
	user := os.Getenv("DB_USER")
	if user == "" {
		user = "root"
	}
	pass := os.Getenv("DB_PASS")
	if pass == "" {
		pass = "example"
	}
	credential := options.Credential{
		Username: user,
		Password: pass,
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
