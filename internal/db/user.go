package db

import (
	"context"
	"log"

	"go-api/internal/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func AddUser(collection *mongo.Collection, user model.User) (*mongo.InsertOneResult, error) {
	res, err := collection.InsertOne(context.Background(), user)
	return res, err
}

// GetUser ...
func GetUser(collection *mongo.Collection, id int) (bson.M, error) {
	var result bson.M

	err := collection.FindOne(context.TODO(), bson.D{{"_id", id}}).Decode(&result)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			log.Printf("No users found: %v", result)
			return result, err
		}
		log.Fatal(err)
	}

	return result, nil
}

// AllUsers ...
func AllUsers(collection *mongo.Collection) ([]bson.M, error) {

	var users []bson.M

	cursor, err := collection.Find(context.Background(), bson.M{})

	if err != nil {
		if err == mongo.ErrNoDocuments {
			log.Printf("No users found: %v", users)
			return users, nil
		}
		log.Fatal(err)
	}

	if err = cursor.All(context.Background(), &users); err != nil {
		log.Fatal(err)
	}

	return users, nil
}
