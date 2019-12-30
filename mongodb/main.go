package main

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Client returns a mongo.Database
func Client(uri string, database string) *mongo.Database {
	clientOptions := options.Client().ApplyURI(uri)

	conn, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	client := conn.Database(database)

	return client
}
