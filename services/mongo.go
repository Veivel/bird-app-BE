package services

import (
	"bird-app/models"
	"context"
	"fmt"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Database

func ConnectDB() {
	ctx, cancel := context.WithTimeout(context.Background(), 25*time.Second)
	defer cancel()

	options := options.Client().ApplyURI(os.Getenv("MONGO_CONNECTION_STRING")) // get options from connection string
	client, err := mongo.Connect(ctx, options)
	if err != nil {
		fmt.Println("Could not connect to Mongo DB:", err)
	}

	DB = client.Database(os.Getenv("MONGO_DB_NAME"))

	/// TESTING STUFF ///
	cursor, err := DB.Collection("users").Find(ctx, bson.D{})
	// cursor, err := DB.Collection("posts").Indexes().List(ctx)
	if err != nil {
		fmt.Println("ERR:", err)
	} else {
		user := models.User{}
		for cursor.TryNext(ctx) {
			cursor.Next(ctx)
			// fmt.Println(cursor.Current.Elements())
			cursor.Decode(&user)
			fmt.Println(user)

		}
	}
}
