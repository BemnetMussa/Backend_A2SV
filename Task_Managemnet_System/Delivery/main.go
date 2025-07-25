package main

import (
	"context"
	"fmt"
	"log"

	"github.com/BemnetMussa/Backend_A2SV/Task_Managemnet_System/Delivery/routers"
	"github.com/BemnetMussa/Backend_A2SV/Task_Managemnet_System/infrastructure"
	"github.com/BemnetMussa/Backend_A2SV/Task_Managemnet_System/repositories"
	"github.com/BemnetMussa/Backend_A2SV/Task_Managemnet_System/Delivery/controllers"
	"github.com/BemnetMussa/Backend_A2SV/Task_Managemnet_System/usecases"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	jwtService := infrastructure.NewJWTService("jwt_secrete_code_temporary", "my.app.test") // Use your actual secret key
	passwordService := infrastructure.NewPasswordService()
	// Connect to MongoDB
	taskCollection, userCollection := connectToMongoDB()

	// Initialize Repositories
	taskRepo := repositories.NewMongoTaskRepository(taskCollection)
	userRepo := repositories.NewMongoUserRepository(userCollection)

	// Initialize Usecases
	taskUsecase := usecases.NewTaskUsecase(taskRepo)
	userUsecase := usecases.NewUserUsecase(userRepo, jwtService, passwordService)

	// controllers
	userController := controllers.NewUserController(userUsecase)
	taskController := controllers.NewTaskController(taskUsecase)

	// Setup Gin router and inject usecases
	r := gin.Default()
	
	routers.SetupRoute(r, userController, taskController)

	// Start server
	if err := r.Run(); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}

func connectToMongoDB() (*mongo.Collection, *mongo.Collection) {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatalf("MongoDB connection failed: %v", err)
	}

	if err := client.Ping(context.TODO(), nil); err != nil {
		log.Fatalf("MongoDB ping failed: %v", err)
	}

	taskCollection := client.Database("taskdb").Collection("tasks")
	userCollection := client.Database("taskdb").Collection("users")

	fmt.Println("Connected to MongoDB")
	return taskCollection, userCollection
}
