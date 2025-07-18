package controllers

import (
	"net/http"

	"github.com/BemnetMussa/Backend_A2SV/tree/main/Task_Managemnet_System/data"
	"github.com/BemnetMussa/Backend_A2SV/tree/main/Task_Managemnet_System/models"
	"github.com/gin-gonic/gin"
)

func GetTasksController(c *gin.Context) {
	tasks, err := data.GetAllTasks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch tasks"})
		return
	}
	c.JSON(http.StatusOK, tasks)
}

func GetTaskDetailController(c *gin.Context) {
	id := c.Param("id")
	task, err := data.GetTaskDetail(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}
	c.JSON(http.StatusOK, task)
}

func UpdateTaskController(c *gin.Context) {
	id := c.Param("id")

	var updatedData models.Task
	if err := c.ShouldBindJSON(&updatedData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task data"})
		return
	}

	if err := data.UpdateTask(id, updatedData); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update task"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Task updated successfully"})
}

func DeleteTaskController(c *gin.Context) {
	id := c.Param("id")
	if err := data.RemoveTask(id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Task deleted successfully"})
}

func CreateTaskController(c *gin.Context) {
	var newData models.Task

	if err := c.ShouldBindJSON(&newData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task data"})
		return
	}

	createdTask, err := data.CreateTask(newData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create task"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Task created successfully",
		"id":      createdTask.ID.Hex(), // convert ObjectID to string
	})
}
