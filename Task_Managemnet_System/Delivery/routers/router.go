package routers

import (
	"github.com/BemnetMussa/Backend_A2SV/Task_Managemnet_System/Delivery/controllers"
	"github.com/BemnetMussa/Backend_A2SV/Task_Managemnet_System/infrastructure"
	"github.com/gin-gonic/gin"
)

func SetupRoute(r *gin.Engine, userController *controllers.UserController, taskController *controllers.TaskController) {
	r.POST("/register", userController.RegisterUser)
	r.POST("/login", userController.Login)

	auth := r.Group("/", infrastructure.AuthMiddleware())
	{
		auth.GET("/tasks", taskController.GetAllTasks)
		auth.GET("/tasks/:id", taskController.GetTaskByID)

		admin := auth.Group("/admin", infrastructure.AdminOnlyMiddleware())
		{
			admin.POST("/tasks", taskController.CreateTask)
			admin.PUT("/tasks/:id", taskController.UpdateTask)
			admin.DELETE("/tasks/:id", taskController.DeleteTask)
			admin.POST("/promote/:email", userController.PromoteUser)
		}
	}
}
