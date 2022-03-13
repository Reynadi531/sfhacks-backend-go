package database

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserBSON struct {
	Id       string `bson:"_id,omitempty"`
	Name     string `bson:"name"`
	Email    string `bson:"email"`
	Password string `bson:"password"`
	Phone    string `bson:"phone"`
	Age      int8   `bson:"age"`
	Gender   string `bson:"gender"`
}

func UserCollection() *mongo.Collection {
	return Database.Collection("users")
}

func InsertUser(user UserBSON) (interface{}, error) {
	result, err := UserCollection().InsertOne(context.TODO(), user)
	if err != nil {
		return nil, err
	}

	return result.InsertedID, nil
}

func FindUserByEmail(email string) (UserBSON, error) {
	var result UserBSON

	if err := UserCollection().FindOne(context.TODO(), bson.D{{"email", email}}).Decode(&result); err != nil {
		return result, err
	}

	return result, nil
}
