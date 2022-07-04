package services

import (
	"context"
	"daft-stats/graph/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func InsertProperty(db *mongo.Database, property model.Property) error {
	_, err := db.Collection("properties").InsertOne(context.TODO(), property)
	if err != nil {
		return err
	}
	return nil
}

func UpdateProperty(db *mongo.Database, property model.Property) error {
	collection := db.Collection("properties")
	_, err := collection.UpdateOne(context.TODO(), bson.D{{"_id", property.ID}}, property)
	if err != nil {
		return err
	}
	return nil
}

func FindProperties(db *mongo.Database) ([]*model.Property, error) {
	var properties []*model.Property
	collection := db.Collection("properties")
	cursor, err := collection.Find(context.TODO(), bson.D{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())
	for cursor.Next(context.TODO()) {
		var property model.Property
		err := cursor.Decode(&property)
		if err != nil {
			return nil, err
		}
		properties = append(properties, &property)
	}
	return properties, nil
}

func FindProperty(db *mongo.Database, id string) (*model.Property, error) {
	var property model.Property
	collection := db.Collection("properties")
	err := collection.FindOne(context.TODO(), bson.D{{"_id", id}}).Decode(&property)
	if err != nil {
		return nil, err
	}
	return &property, nil
}
