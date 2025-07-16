package router

import (
	"github.com/BemnetMussa/Backend_A2SV/tree/main/Task_Managemnet_System/controllers"
	"github.com/gin-gonic/gin"
)

/* GET /tasks: Get a list of all tasks.
GET /tasks/:id: Get the details of a specific task.
PUT /tasks/:id: Update a specific task. This endpoint should accept a JSON body with the new details of the task.
DELETE /tasks/:id: Delete a specific task.
POST /tasks: Create a new task. This endpoint should accept a JSON body with the task's title, description, due date, and status.

*/

func SetupRoute(r *gin.Engine) {
	r.GET("/tasks", controllers.GetTasksController)
	r.GET("/tasks/:id", controllers.GetTaskDetailController)
	r.PUT("/tasks/:id", controllers.UpdateTaskController)
	r.DELETE("/tasks/:id", controllers.DeleteTaskController)
	r.POST("/tasks", controllers.CreateTaskController)
}
