package main

import (
	"context"
	"log"
	"fmt"
	"github.com/BemnetMussa/Backend_A2SV/tree/main/Task_Managemnet_System/router"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"github.com/BemnetMussa/Backend_A2SV/tree/main/Task_Managemnet_System/data"
)

var TaskCollection *mongo.Collection
var UserCollection *mongo.Collection


func main() {
	ConnectToMongoDB()
	r := gin.Default()
	router.SetupRoute(r)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}


func ConnectToMongoDB() {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	TaskCollection = client.Database("taskdb").Collection("tasks")
	UserCollection = client.Database("taskdb").Collection("users")
	data.SetTaskCollection(TaskCollection)
	data.SetUserCollection(UserCollection)
	fmt.Println("Connected to MongoDB!")
}
