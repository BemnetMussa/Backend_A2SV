package router

import (
	"github.com/BemnetMussa/Backend_A2SV/tree/main/Task_Managemnet_System/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRoute(r *gin.Engine) {
	r.GET("/tasks", controllers.GetTasksController)
	r.GET("/tasks/:id", controllers.GetTaskDetailController)
}
