package data

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	database = db("mongodb://localhost:27017", "test")
)

func db(uri string, name string) *mongo.Database {
	opts := options.Client().ApplyURI(uri)

	c, err := mongo.Connect(context.Background(), opts)
	if err != nil {
		log.Fatal(err)
	}

	return c.Database(name)
}

type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Name      string
	CreatedAt int64
}

// FindOneAndUpdate wrapper around FindOneAndUpdate mongo driver.
func FindOneAndUpdate(ctx context.Context, filter, update interface{}) (*User, error) {
	coll := database.Collection("users")

	upsert := true
	after := options.After

	opts := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
		Upsert:         &upsert,
	}

	var res User
	err := coll.FindOneAndUpdate(ctx, filter, update, &opts).Decode(&res)
	if err != nil {
		return nil, fmt.Errorf("FindOneAndUpdate %w", err)
	}
	return &res, nil
}

// FindOne wrapper around FindOne mongo driver.
func FindOne(ctx context.Context, filter interface{}) (*User, error) {
	coll := database.Collection("users")

	var res User
	err := coll.FindOne(ctx, filter).Decode(&res)
	if err != nil {
		return nil, fmt.Errorf("FindOne %w", err)
	}
	return &res, nil
}

func main() {
	filter := bson.M{"email": "email@test.com"}
	update := bson.M{
		"$set": bson.M{
			"createdAt": time.Now().Add(15 * time.Minute).Unix(),
			"name":      "secret",
		}}

	res, _ := FindOneAndUpdate(context.Background(), filter, update)
	fmt.Println(res.Name, res.CreatedAt)
}
