package usecases_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/BemnetMussa/Backend_A2SV/Task_Managemnet_System/domain"
	"github.com/BemnetMussa/Backend_A2SV/Task_Managemnet_System/usecases"
	"github.com/BemnetMussa/Backend_A2SV/Task_Managemnet_System/repositories"
)

func TestCreateTask_Success(t *testing.T) {
	mockRepo := new(repositories.MockTaskRepository)
	uc := usecases.NewTaskUsecase(mockRepo)

	input := domain.Task{
		Title:       "Test",
		Description: "Testing",
	}

	expected := domain.Task{
		ID:          primitive.NewObjectID(),
		Title:       "Test",
		Description: "Testing",
	}

	mockRepo.On("Create", mock.Anything, input).Return(expected, nil)

	result, err := uc.CreateTask(input)

	assert.NoError(t, err)
	assert.Equal(t, expected, result)
	mockRepo.AssertExpectations(t)
}

func TestGetAllTasks_Success(t *testing.T) {
	mockRepo := new(repositories.MockTaskRepository)
	uc := usecases.NewTaskUsecase(mockRepo)

	expected := []domain.Task{
		{ID: primitive.NewObjectID(), Title: "Task 1"},
		{ID: primitive.NewObjectID(), Title: "Task 2"},
	}

	mockRepo.On("GetAll", mock.Anything).Return(expected, nil)

	result, err := uc.GetAllTasks()

	assert.NoError(t, err)
	assert.Equal(t, expected, result)
	mockRepo.AssertExpectations(t)
}

func TestGetTaskByID_Success(t *testing.T) {
	mockRepo := new(repositories.MockTaskRepository)
	uc := usecases.NewTaskUsecase(mockRepo)

	taskID := primitive.NewObjectID()
	expected := domain.Task{ID: taskID, Title: "Sample"}

	mockRepo.On("GetByID", mock.Anything, taskID.Hex()).Return(expected, nil)

	result, err := uc.GetTaskByID(taskID.Hex())

	assert.NoError(t, err)
	assert.Equal(t, expected, result)
	mockRepo.AssertExpectations(t)
}


func TestUpdateTask_Success(t *testing.T) {
	mockRepo := new(repositories.MockTaskRepository)
	uc := usecases.NewTaskUsecase(mockRepo)

	taskID := primitive.NewObjectID()
	input := domain.Task{Title: "Updated"}

	mockRepo.On("Update", mock.Anything, taskID.Hex(), input).Return(nil)

	err := uc.UpdateTask(taskID.Hex(), input)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestDeleteTask_Success(t *testing.T) {
	mockRepo := new(repositories.MockTaskRepository)
	uc := usecases.NewTaskUsecase(mockRepo)

	taskID := primitive.NewObjectID()

	mockRepo.On("Delete", mock.Anything, taskID.Hex()).Return(nil)

	err := uc.DeleteTask(taskID.Hex())

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}
