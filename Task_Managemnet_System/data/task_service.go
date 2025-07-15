package data

import (
	"github.com/BemnetMussa/Backend_A2SV/tree/main/Task_Managemnet_System/models"
)

// Task DB using maps
var TaskDB = map[int]models.Task{
	1: {ID: 1, Title: "complete api", Detail: "for the task three it must be compelted", Done: true},
}

type TaskResponse struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}


func GetAllTasks() []TaskResponse {
	tasks := []TaskResponse{}
	for _, task := range TaskDB {
		tasks = append(tasks, TaskResponse{
			ID: task.ID,
			Title: task.Title,
			Done: task.Done,
		})
	}	
	return tasks
}

func GetTaskDetail(taskId int) (string, bool) {
	task, exist := TaskDB[taskId]
	if !exist {
		return "", false
	}

	return task.Detail, exist
}