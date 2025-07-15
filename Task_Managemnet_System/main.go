package main

import (
	"fmt"
)
/* 

Implement a REST API with the following endpoints:
GET /tasks: Get a list of all tasks.
GET /tasks/:id: Get the details of a specific task.
PUT /tasks/:id: Update a specific task. This endpoint should accept a JSON body with the new details of the task.
DELETE /tasks/:id: Delete a specific task.
POST /tasks: Create a new task. This endpoint should accept a JSON body with the task's title, description, due date, and status.

*/

func main() {
	
	fmt.Print("Task manager system!")
	
}