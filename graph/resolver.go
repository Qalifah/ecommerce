package graph

//go:generate go run github.com/99designs/gqlgen

import(
	"context"
	"log"
	"time"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct{
	DB *mongo.Database
}


func ConnectDB() *mongo.Client {
	ctx, cancel := context.WithTimeout(context.Background(), 15 * time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://127.0.0.1:27017/"))
	checkError(err)
	err = client.Ping(ctx, readpref.Primary())
	checkError(err)
	fmt.Println("Successfully connected and pinged.")
	return client
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}