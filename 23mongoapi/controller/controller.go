package controller

import (
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

const connectionString = "mongodb+srv://<db_username>:<db_password>@cluster0.n44sg.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0"
const dbName = "netflix"
const colName = "watchlist"

// IMP
var collection *mongo.Collection

// Connect with mongoDB

func init() {
	// client options
	clientOption := options.Client().ApplyURI(connectionString)

	// Connect to mongoDB
	client, err := mongo.Connect(clientOption)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("MongoDB connection successful")

	collection = client.Database(dbName).Collection(colName)

	// collection instance
	fmt.Println("Collection instance is ready")
}
