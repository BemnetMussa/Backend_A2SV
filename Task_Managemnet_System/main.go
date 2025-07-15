package main

/*

Implement a REST API with the following endpoints:
GET /tasks: Get a list of all tasks.
GET /tasks/:id: Get the details of a specific task.
PUT /tasks/:id: Update a specific task. This endpoint should accept a JSON body with the new details of the task.
DELETE /tasks/:id: Delete a specific task.
POST /tasks: Create a new task. This endpoint should accept a JSON body with the task's title, description, due date, and status.

*/

import (
	"github.com/BemnetMussa/Backend_A2SV/tree/main/Task_Managemnet_System/router"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	
	router.SetupRoute(r)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}