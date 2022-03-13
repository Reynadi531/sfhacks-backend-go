package database

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type LoggerBSON struct {
	Id        primitive.ObjectID `bson:"_id,omitempty"`
	UserId    string             `bson:"userid"`
	Heartrate int16              `bson:"heartrate"`
	Emotion   string             `bson:"emotion"`
	Result    string             `bson:"result"`
	Time      time.Time          `bson:"time"`
}

func LoggerCollection() *mongo.Collection {
	return Database.Collection("logger")
}

func InsertLog(data LoggerBSON) (interface{}, error) {
	result, err := LoggerCollection().InsertOne(context.TODO(), data)
	if err != nil {
		return nil, err
	}

	return result.InsertedID, nil
}

func GetLogsByUserId(userid string, limit int64) ([]LoggerBSON, error) {
	var logs []LoggerBSON

	findOptions := options.Find()
	findOptions.SetLimit(limit)
	cursor, err := LoggerCollection().Find(context.TODO(), bson.D{{"userid", userid}}, findOptions)

	if err != nil {
		return logs, err
	}

	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		var log LoggerBSON
		if err := cursor.Decode(&log); err != nil {
			panic(err)
		}

		logs = append(logs, log)
	}

	return logs, nil
}
