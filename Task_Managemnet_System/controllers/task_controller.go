package controllers

import (
	"net/http"
	"strconv"

	"github.com/BemnetMussa/Backend_A2SV/tree/main/Task_Managemnet_System/data"
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
	detailTask, exists  := data.GetTaskDetail(taskId)

	if !exists {
		c.JSON(404, gin.H{"error": "Task not found!"})
		return
	}
	c.JSON(200, gin.H{"details": detailTask})
}