package config

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb+srv://admin:admin@sanjeevcluster.h2whbxu.mongodb.net/go-lang?retryWrites=true&W=majority"

const dbName = "netflix"
const colName = "watchlist"

// Most important
var Collection *mongo.Collection

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
	Collection = client.Database(dbName).Collection(colName)

	// collection reference
	fmt.Println("Collection instance is ready")
}
