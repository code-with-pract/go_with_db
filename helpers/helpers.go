package helpers

import (
	"context"
	"fmt"
	"log"

	"github.com/snjeev-kushwaha/mongoDbApi/config"
	"github.com/snjeev-kushwaha/mongoDbApi/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Insert 1 record

func InsertOneMovie(movie model.Netflix) {
	inserted, err := config.Collection.InsertOne(context.Background(), movie)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted into 1 movie in db with id:", inserted.InsertedID)
}

// Update 1 record
func UpdateOneMovie(movieId string) {
	id, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"watched": true}}

	result, err := config.Collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("modified count: ", result.ModifiedCount)
}

// Delete one record
func DeleteOneMovie(movieId string) {
	id, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id": id}
	deleteCount, err := config.Collection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Movie go delete with delete Count:", deleteCount)
}

// delete All records from MongoDb
func DeleteAllMovies() int64 {
	deleteResult, err := config.Collection.DeleteMany(context.Background(), bson.D{{}}, nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Number of movies deleted :", deleteResult.DeletedCount)
	return deleteResult.DeletedCount
}

// get All Movies from database
func GetAllMovie() []primitive.M {
	cur, err := config.Collection.Find(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}

	var movies []primitive.M
	for cur.Next(context.Background()) {
		var movie bson.M
		err := cur.Decode(&movie)
		if err != nil {
			log.Fatal(err)
		}
		movies = append(movies, movie)
	}

	defer cur.Close(context.Background())
	return movies
}
