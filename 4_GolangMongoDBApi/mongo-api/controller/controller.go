package controller

import (
	"context"
	"fmt"
	"log"

	"github.com/krishnalagad/mongo-api/model"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	// "go.mongodb.org/mongo-driver/x/mongo/driver/mongocrypt/options"
)

const connectionString = "mongodb://localhost:27017"
const dbName = "netflix"
const collName = "watchlist"

// MOST IMPORTANT
var collection *mongo.Collection

// connect with mongoDB
func init() {

	// client options
	clientOption := options.Client().ApplyURI(connectionString)

	// connect to mongoDB
	client, err := mongo.Connect(context.TODO(), clientOption)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("MongoDB connections successfull")

	// collection = (*mongo.Collection)(client.Database(dbName).Collection(collName))
	collection = client.Database(dbName).Collection(collName)

	// collection instance
	fmt.Println("Collection instance is ready")

}

// MONGODB Helpers - file

// insert one record
func insertOneMovie(movie model.Netflix){
	inserted, err := collection.InsertOne(context.Background(), movie)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted 1 movie in db with id: ", inserted.InsertedID)
}
