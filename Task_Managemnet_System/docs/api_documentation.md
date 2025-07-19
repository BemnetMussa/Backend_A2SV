# Task Management REST API Documentation

## Overview

This is a RESTful API for managing tasks, built using the Go programming language, Gin framework, and MongoDB for persistent storage. The API supports full CRUD operations with role-based access control:

* *   **Admins** can create, update, and delete tasks.
*     
* *   **Users** can view tasks.
*     

* * *

## Base URL

arduino

CopyEdit

`http://localhost:8080`

* * *

## Authentication

All protected endpoints require a valid JWT token in the `Authorization` header using the Bearer scheme.

### Example:

makefile

CopyEdit

`Authorization: Bearer <your_jwt_token>`

Tokens include the user's `email` and `role` as claims.

* * *

## Endpoints Summary

Method

Endpoint

Description

Access

GET

`/tasks`

Get all tasks

User/Admin

GET

`/tasks/:id`

Get task by ID

User/Admin

POST

`/tasks`

Create a new task

Admin

PUT

`/tasks/:id`

Update an existing task

Admin

DELETE

`/tasks/:id`

Delete a task by ID

Admin

* * *

## POST `/tasks`

Create a new task. Requires admin access.

### Request Body (JSON)

json

CopyEdit

`{   "title": "Finish project",   "description": "Implement all features",   "completed": false }`

### Success Response

* *   **Code:** `200 OK`
*     

json

CopyEdit

`{   "message": "Task created successfully",   "task": {     "id": "615f7e2bc9d7a6f8c6dfc123",     "title": "Finish project",     "description": "Implement all features",     "completed": false   } }`

### Error Response

* *   **Code:** `400 Bad Request`
*     

json

CopyEdit

`{   "error": "Invalid task data" }`

* *   **Code:** `403 Forbidden` (if non-admin)
*     

json

CopyEdit

`{   "error": "Admin access required" }`

* * *

## GET `/tasks`

Retrieve all tasks.

### Success Response

* *   **Code:** `200 OK`
*     

json

CopyEdit

`[   {     "id": "615f7e2bc9d7a6f8c6dfc123",     "title": "Finish project",     "description": "Implement all features",     "completed": false   } ]`

* * *

## GET `/tasks/:id`

Retrieve task details by ID.

### Success Response

* *   **Code:** `200 OK`
*     

json

CopyEdit

`{   "id": "615f7e2bc9d7a6f8c6dfc123",   "title": "Finish project",   "description": "Implement all features",   "completed": false }`

### Error Response

* *   **Code:** `404 Not Found`
*     

json

CopyEdit

`{   "error": "Task not found" }`

* * *

## PUT `/tasks/:id`

Update a task by ID. Admin-only.

### Request Body (JSON)

json

CopyEdit

`{   "title": "Finalize draft",   "description": "Fix validation logic",   "completed": true }`

### Success Response

* *   **Code:** `200 OK`
*     

json

CopyEdit

`{   "task": {     "id": "615f7e2bc9d7a6f8c6dfc123",     "title": "Finalize draft",     "description": "Fix validation logic",     "completed": true   } }`

### Error Response

* *   **Code:** `404 Not Found`
*     

json

CopyEdit

`{   "error": "Task not found" }`

* *   **Code:** `403 Forbidden`
*     

json

CopyEdit

`{   "error": "Admin access required" }`

* * *

## DELETE `/tasks/:id`

Delete a task by ID. Admin-only.

### Success Response

* *   **Code:** `200 OK`
*     

json

CopyEdit

`{   "message": "Task deleted successfully" }`

### Error Response

* *   **Code:** `404 Not Found`
*     

json

CopyEdit

`{   "error": "Task not found" }`

* *   **Code:** `403 Forbidden`
*     

json

CopyEdit

`{   "error": "Admin access required" }`

* * *

## User Roles

The system has two roles:

* *   **Admin**: Full access to all endpoints, including task creation, updating, and deletion.
*     
* *   **User**: Read-only access to tasks.
*     

Role-based access is enforced using JWT token claims.

* * *

## MongoDB Integration

All task data is persisted in a MongoDB collection. Each task is stored with a unique ObjectID.

* * *

## Postman Testing Tips

* *   Set `Content-Type: application/json` for `POST` and `PUT` requests.
*     
* *   Include a valid JWT token in the `Authorization` header.
*     
* *   Ensure your MongoDB server is running locally or remotely and properly connected.
*     
* *   Admin operations require a token with `role: "admin"`.
*

Copy Markdown

