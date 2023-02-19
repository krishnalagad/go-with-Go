package payload

import (
	"context"
	"fmt"
	"log"

	"github.com/krishnalagad/docman-api/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb://localhost:27017"
const dbName = "docman"
const collName = "docs"

var collection *mongo.Collection

// establish connection with MongoDB server
func init() {

	// client options
	clientOption := options.Client().ApplyURI(connectionString)

	// connect to mongodb
	client, err := mongo.Connect(context.TODO(), clientOption)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB server")

	collection = (*mongo.Collection)(client.Database(dbName).Collection(collName))

	fmt.Println("Collection instance is ready.")
}

// helper function to add one record.
func InsertOneDocument(doc model.Document) interface{} {
	data, err := collection.InsertOne(context.Background(), doc)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted 1 movie with id: ", data.InsertedID)
	return data
}

// helper function to get one record.
func GetOneDocument(docid string) model.Document {
	id, _ := primitive.ObjectIDFromHex(docid) // converting string id into ObjectId
	filter := bson.M{"_id": id}               // setting id to get perticular record.

	var document model.Document
	err := collection.FindOne(context.Background(), filter).Decode(&document)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			// This error means your query did not match any documents.
			return document
		}
		panic(err)
	}
	return document
}

// helper function to update one record.
func UpdateOneDocument(doc model.Document, docid string) model.Document {
	id, _ := primitive.ObjectIDFromHex(docid)
	// document := GetOneDocument(docid)
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"name": doc.DocumentName, "type": doc.DocumentType, "islegal": doc.Legal, "Owner": doc.Owner}}

	result, err := collection.UpdateOne(context.Background(), filter, update)

	// result, err := collection.UpdateOne(
	// 	context.Background(),
	// 	bson.M{"_id": id},
	// 	bson.M{"$set", bson.M{"name": doc.}},
	// )
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Modified count: ", result.ModifiedCount)

	return GetOneDocument(docid)
}
