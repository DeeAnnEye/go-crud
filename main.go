package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func main() {
	
	client,err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://daniela:12345@cluster0.u4ybs.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"))
	if err != nil { log.Fatal(err) }
    ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil { log.Fatal(err) }

	defer client.Disconnect(ctx)
	err = client.Ping(ctx, readpref.Primary())
	if err != nil { log.Fatal(err) }

	// databases,err := client.ListDatabaseNames(ctx,bson.M{})
	// if err != nil { log.Fatal(err) }
	// fmt.Println(databases)

	testDatabase := client.Database("test")
	taskCollection := testDatabase.Collection("task")

	taskResult,err := taskCollection.InsertMany(ctx, []interface{}{
		bson.D{
			{"taskName","Shopping"},
			{"taskDay","17th May"},		
		},
		bson.D{
			{"taskName","Groceries"},
			{"taskDay","20th May"},		
		},
	})
	if err != nil { log.Fatal(err) }
	fmt.Println(taskResult.InsertedIDs...)
}
