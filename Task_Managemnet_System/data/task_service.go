package data

import (
	"github.com/BemnetMussa/Backend_A2SV/tree/main/Task_Managemnet_System/models"
)

// Task DB using maps
var TaskDB = map[int]models.Task{
	1: {ID: 1, Title: "complete api", Description: "for the task three it must be compelted", Completed: true},
}
var currentID = 2

type TaskResponse struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Completed  bool   `json:"Completed"`
}


func GetAllTasks() []TaskResponse {
	tasks := []TaskResponse{}
	for _, task := range TaskDB {
		tasks = append(tasks, TaskResponse{
			ID: task.ID,
			Title: task.Title,
			Completed: task.Completed,
		})
	}	
	return tasks
}

func GetTaskDetail(taskId int) (string, bool) {
	task, exist := TaskDB[taskId]
	if !exist {
		return "", false
	}
	return task.Description, exist
}

func UpdateTask(id int, newTask models.Task) (models.Task, bool) {
	task, exists := TaskDB[id]
	if !exists {
		return models.Task{}, false
	}

	// Update logic
	task.Title = newTask.Title
	task.Description = newTask.Description
	task.Completed = newTask.Completed

	TaskDB[id] = task
	return task, true
}

func RemoveTask(id int) bool {
	_ , exists := TaskDB[id]
	if !exists {
		return false
	}
	delete(TaskDB, id)
	return true
}

func CreateTask(newTask models.Task) models.Task {
	newTask.ID = currentID
	TaskDB[currentID] = newTask
	currentID++
	return newTask
}
