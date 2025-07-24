package repositories

import (
	"context"
	"errors"
	"time"

	"github.com/BemnetMussa/Backend_A2SV/Task_Managemnet_System/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type TaskRepository interface {
	Create(ctx context.Context, task domain.Task) (domain.Task, error)
	GetAll(ctx context.Context) ([]domain.Task, error)
	GetByID(ctx context.Context, id string) (domain.Task, error)
	Update(ctx context.Context, id string, task domain.Task) error
	Delete(ctx context.Context, id string) error
}

type MongoTaskRepository struct {
	Collection *mongo.Collection
}

func NewMongoTaskRepository(c *mongo.Collection) *MongoTaskRepository {
	return &MongoTaskRepository{Collection: c}
}

func (r *MongoTaskRepository) Create(ctx context.Context, task domain.Task) (domain.Task, error) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	res, err := r.Collection.InsertOne(ctx, task)
	if err != nil {
		return task, err
	}
	task.ID = res.InsertedID.(primitive.ObjectID)
	return task, nil
}

func (r *MongoTaskRepository) GetAll(ctx context.Context) ([]domain.Task, error) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	cursor, err := r.Collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var tasks []domain.Task
	for cursor.Next(ctx) {
		var task domain.Task
		if err := cursor.Decode(&task); err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return tasks, nil
}

func (r *MongoTaskRepository) GetByID(ctx context.Context, id string) (domain.Task, error) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return domain.Task{}, err
	}

	var task domain.Task
	err = r.Collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&task)
	return task, err
}

func (r *MongoTaskRepository) Update(ctx context.Context, id string, task domain.Task) error {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	update := bson.M{
		"$set": bson.M{
			"title":       task.Title,
			"description": task.Description,
			"completed":   task.Completed,
		},
	}

	_, err = r.Collection.UpdateOne(ctx, bson.M{"_id": objID}, update)
	return err
}

func (r *MongoTaskRepository) Delete(ctx context.Context, id string) error {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	result, err := r.Collection.DeleteOne(ctx, bson.M{"_id": objID})
	if err != nil {
		return err
	}
	if result.DeletedCount == 0 {
		return errors.New("task not found")
	}
	return nil
}
