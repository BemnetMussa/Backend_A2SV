package data

import (
	"context"
	"time"

	"github.com/BemnetMussa/Backend_A2SV/tree/main/Task_Managemnet_System/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

)

var TaskCollection *mongo.Collection

// Set the MongoDB collection from main.go
func SetTaskCollection(c *mongo.Collection) {
	TaskCollection = c
}


func CreateTask(newTask models.Task) (models.Task, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	res, err := TaskCollection.InsertOne(ctx, newTask)
	if err != nil {
		return newTask, err
	}

	newTask.ID = res.InsertedID.(primitive.ObjectID)
	return newTask, nil
} 

func GetAllTasks() ([]models.Task, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := TaskCollection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var tasks []models.Task
	if err := cursor.All(ctx, &tasks); err != nil {
		return nil, err
	}

	return tasks, nil
}	

func GetTaskDetail(taskId string) (models.Task, error) {
	var task models.Task

	// Convert string ID to MongoDB ObjectID
	objID, err := primitive.ObjectIDFromHex(taskId)
	if err != nil {
		return task, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = TaskCollection.FindOne(ctx, bson.M{"_id": objID}).Decode(&task)
	if err != nil {
		return task, err
	}

	return task, nil
}


func UpdateTask(id string, updatedData models.Task) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	update := bson.M{
		"$set": bson.M{
			"title":       updatedData.Title,
			"description": updatedData.Description,
			"completed":   updatedData.Completed,
		},
	}
	_, err = TaskCollection.UpdateOne(ctx, bson.M{"_id": objID}, update)
	return err
}


func RemoveTask(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	_, err = TaskCollection.DeleteOne(ctx, bson.M{"_id": objID})
	return err
}

