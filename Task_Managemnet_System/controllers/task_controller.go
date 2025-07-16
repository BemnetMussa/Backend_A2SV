package controllers

import (
	"net/http"
	"strconv"

	"github.com/BemnetMussa/Backend_A2SV/tree/main/Task_Managemnet_System/data"
	"github.com/BemnetMussa/Backend_A2SV/tree/main/Task_Managemnet_System/models"
	"github.com/gin-gonic/gin"
)

func GetTasksController(c *gin.Context) {
	tasks := data.GetAllTasks()
	c.JSON(http.StatusOK, tasks)
}

func GetTaskDetailController(c *gin.Context) {
	id := c.Param("id")
	taskId, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid id"})
	}
	detailTask, exists := data.GetTaskDetail(taskId)

	if !exists {
		c.JSON(404, gin.H{"error": "Task not found!"})
		return
	}
	c.JSON(200, gin.H{"details": detailTask})
}

func UpdateTaskController(c *gin.Context) {
	id := c.Param("id")
	taskId, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid ID"})
		return
	}

	var updatedData models.Task
	if err := c.ShouldBindJSON(&updatedData); err != nil {
		c.JSON(400, gin.H{"error": "Invalid task data"})
		return
	}

	updatedTask, exists := data.UpdateTask(taskId, updatedData)
	if !exists {
		c.JSON(404, gin.H{"error": "Task not found"})
		return
	}

	c.JSON(200, gin.H{"task": updatedTask})
}

func DeleteTaskController(c *gin.Context) {
	id := c.Param("id")
	taskId, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid ID"})
		return
	}
	if !data.RemoveTask(taskId) {
		c.JSON(404, gin.H{"error": "Task not found!"})
		return
	}

	c.JSON(200, gin.H{"message": "Task Removed Successfully"})
}

func CreateTaskController(c *gin.Context) {
	var newData models.Task

	if err := c.ShouldBindJSON(&newData); err != nil {
		c.JSON(400, gin.H{"error": "Invalid task data"})
		return
	}

	createdTask := data.CreateTask(newData)
	c.JSON(200, gin.H{
		"message": "Task created successfully",
		"task":    createdTask,
	})
}
