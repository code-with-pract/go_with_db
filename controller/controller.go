package controller

import (
	"context"
	"fmt"
	"log"

	"github.com/snjeev-kushwaha/mongoDbApi/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb+srv://admin:admin@sanjeevcluster.h2whbxu.mongodb.net/go-lang?retryWrites=true&W=majority"

const dbName = "netflix"
const colName = "watchlist"

// Most important
var collection *mongo.Collection

// connect with mongoDB
func init() {
	// client option
	clientOption := options.Client().ApplyURI(connectionString)

	// connect to mongoDB
	client, err := mongo.Connect(context.TODO(), clientOption)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("MongoDB connection success")
	collection = client.Database(dbName).Collection(colName)

	// collection reference
	fmt.Println("Collection instance is ready")
}

// MongoDb helpers - file

// Insert 1 record

func insertOneMovie(movie model.Netflix) {
	inserted, err := collection.InsertOne(context.Background(), movie)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted into 1 movie in db with id:", inserted.InsertedID)
}

// Update 1 record
func updateOneMovie(movieId string) {
	id, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"watched": true}}

	result, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("modified count: ", result.ModifiedCount)
}

// Delete one record
func deleteOneMovie(movieId string){
	
}