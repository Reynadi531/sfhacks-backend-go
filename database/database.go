package database

import (
	"context"
	"errors"
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	Connection *mongo.Client
	Database   *mongo.Database
)

func InitDatabase() error {
	var err error

	mongoURI := os.Getenv("MONGO_URI")

	clientOption := options.Client().ApplyURI(mongoURI)

	Connection, err = mongo.NewClient(clientOption)
	if err != nil {
		return err
	}

	err = Connection.Connect(context.Background())
	if err != nil {
		return err
	}

	Database = Connection.Database(os.Getenv("MONGO_DATABASE"))
	fmt.Println(os.Getenv("MONGO_DATABASE"))

	fmt.Println("[LOG] Database succesfully connected")
	return nil
}

func DBDisconnect() error {
	if Connection == nil {
		return errors.New("connection hasnt initialized")
	}

	err := Connection.Disconnect(context.TODO())
	if err != nil {
		return err
	}

	fmt.Println("[LOG] Success disconnect from database")
	return nil
}
