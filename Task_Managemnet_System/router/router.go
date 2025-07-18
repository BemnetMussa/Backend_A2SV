package router

import (
	"github.com/BemnetMussa/Backend_A2SV/tree/main/Task_Managemnet_System/controllers"
	"github.com/BemnetMussa/Backend_A2SV/tree/main/Task_Managemnet_System/middleware"
	"github.com/gin-gonic/gin"
)

/* GET /tasks: Get a list of all tasks.
GET /tasks/:id: Get the details of a specific task.
PUT /tasks/:id: Update a specific task. This endpoint should accept a JSON body with the new details of the task.
DELETE /tasks/:id: Delete a specific task.
POST /tasks: Create a new task. This endpoint should accept a JSON body with the task's title, description, due date, and status.


router.POST("/register", func(c *gin.Context) {
  var user User
  if err := c.ShouldBindJSON(&user); err != nil {
    c.JSON(400, gin.H{"error": "Invalid request payload"})
    return
  }

  // TODO: Implement user registration logic
  c.JSON(200, gin.H{"message": "User registered successfully"})
})
*/

func SetupRoute(r *gin.Engine) {
	
	r.POST("/register", controllers.RegisterUserController)
	r.POST("/login", controllers.LoginController)

	auth := r.Group("/", middleware.AuthMiddleware())
	{
		auth.GET("/tasks", controllers.GetTasksController)
		auth.GET("/tasks/:id", controllers.GetTaskDetailController)

		admin := auth.Group("/admin", middleware.AdminOnlyMiddleware())
		{
			admin.POST("/tasks", controllers.CreateTaskController)
			admin.PUT("/tasks/:id", controllers.UpdateTaskController)
			admin.DELETE("/tasks/:id", controllers.DeleteTaskController)
			admin.POST("/promote/:email", controllers.PromoteUserController)
		}
	}
}

