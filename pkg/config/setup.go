package config

import (
	"context"
	"fmt"
	"log"

	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var MyClient *mongo.Client

func init(){
	ConnectDb()
}


func ConnectDb(){
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(GetMongoURI()))
	if err != nil {
		log.Fatal("Error while connecting database", err)
		panic(err)
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected to MongoDB!")
	MyClient = client
	
}

func GetCollection(collection string) *mongo.Collection{
	if MyClient == nil {
		log.Fatal("MongoDb client is not initalized. Make sure to call ConnectDB() first")
	}
	coll := MyClient.Database("cards").Collection(collection)
	if coll ==nil {
		log.Fatal("Collection not found")
	}
	return coll 
}