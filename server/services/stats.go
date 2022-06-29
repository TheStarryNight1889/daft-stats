package services

import (
	"context"
	"daft-stats/graph/model"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func FindStats(db *mongo.Database) ([]*model.Stat, error) {
	var stats []*model.Stat
	collection := db.Collection("stats")
	cursor, err := collection.Find(context.TODO(), bson.D{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())
	for cursor.Next(context.TODO()) {
		var stat model.Stat
		err := cursor.Decode(&stat)
		if err != nil {
			return nil, err
		}
		stats = append(stats, &stat)
	}
	return stats, nil
}
func InsertStats(db *mongo.Database, stat model.Stat) error {
	_, err := db.Collection("stats").InsertOne(context.TODO(), stat)
	if err != nil {
		fmt.Print(err)
	}
	return err
}
