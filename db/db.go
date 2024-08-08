package db

import (
	"context"
	"fmt"
	"log"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var db *mongo.Client

func InitMongodbClient() {
	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		log.Fatal("Set your 'MONGODB_URI' environment variable. " +
			"See: " +
			"www.mongodb.com/docs/drivers/go/current/usage-examples/#environment-variable")
	}
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))

	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Successfully connected to database")
	db = client
}

func GetDB() *mongo.Database {
	return db.Database("notes")
}

func CloseClientConnection() {
	if err := db.Disconnect(context.TODO()); err != nil {
		panic(err)
	}
}
