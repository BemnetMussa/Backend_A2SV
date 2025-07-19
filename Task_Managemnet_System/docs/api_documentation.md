
---

# 📄 Task Management REST API Documentation

## 🧾 Overview

This is a RESTful API for managing tasks, built with **Go**, the **Gin** framework, and **MongoDB** for persistent storage. It includes full CRUD functionality and **role-based access control**.

* 🔐 **Admins** can create, update, and delete tasks.

* 👤 **Users** can only view tasks.

---

## 🌐 Base URL

```
http://localhost:8080
```

---

## 🔑 Authentication

All protected endpoints require a valid **JWT token** in the `Authorization` header.

### Header Format:

```
Authorization: Bearer <your_token>
```

Tokens include `email` and `role` claims. Admin-only routes will reject non-admin tokens.

---

## 📚 Endpoints Summary

|Method|Endpoint|Description|Access|
|---|---|---|---|
|GET|`/tasks`|Get all tasks|User/Admin|
|GET|`/tasks/:id`|Get task by ID|User/Admin|
|POST|`/tasks`|Create a new task|Admin only 🔐|
|PUT|`/tasks/:id`|Update an existing task|Admin only 🔐|
|DELETE|`/tasks/:id`|Delete a task by ID|Admin only 🔐|

---

## ➕ POST `/tasks`

### Description

Create a new task (Admin only 🔐)

### Request Body

```json
{
  "title": "Finish project",
  "description": "Implement all features",
  "completed": false
}
```

### Success Response

```json
{
  "message": "Task created successfully",
  "task": {
    "id": "615f7e2bc9d7a6f8c6dfc123",
    "title": "Finish project",
    "description": "Implement all features",
    "completed": false
  }
}
```

### Errors

* `400 Bad Request`: Invalid input

* `403 Forbidden`: User is not an admin

---

## 📥 GET `/tasks`

### Description

Fetch all tasks.

### Success Response

```json
[
  {
    "id": "615f7e2bc9d7a6f8c6dfc123",
    "title": "Finish project",
    "description": "Implement all features",
    "completed": false
  }
]
```

---

## 🔍 GET `/tasks/:id`

### Description

Get a specific task by its ID.

### Success Response

```json
{
  "id": "615f7e2bc9d7a6f8c6dfc123",
  "title": "Finish project",
  "description": "Implement all features",
  "completed": false
}
```

### Error

```json
{
  "error": "Task not found"
}
```

---

## ✏️ PUT `/tasks/:id`

### Description

Update a task by ID (Admin only 🔐)

### Request Body

```json
{
  "title": "Finalize draft",
  "description": "Fix validation logic",
  "completed": true
}
```

### Success Response

```json
{
  "task": {
    "id": "615f7e2bc9d7a6f8c6dfc123",
    "title": "Finalize draft",
    "description": "Fix validation logic",
    "completed": true
  }
}
```

### Errors

* `404 Not Found`: Task not found

* `403 Forbidden`: Not an admin

---

## ❌ DELETE `/tasks/:id`

### Description

Delete a task by ID (Admin only 🔐)

### Success Response

```json
{
  "message": "Task deleted successfully"
}
```

### Errors

* `404 Not Found`: Task not found

* `403 Forbidden`: Not an admin

---

## 👥 Roles

There are two types of users:

* **Admin**: Full access to create, update, delete, and view tasks

* **User**: Read-only access (can view tasks only)

Roles are defined in the JWT token under the `role` claim.

---

## 🛢️ Database

The API uses **MongoDB** to store tasks permanently. Each task is stored as a document with a unique `ObjectID`.

---

## 🧪 Postman Testing Tips

* Set `Content-Type` to `application/json` for `POST` and `PUT` requests.

* Include the JWT token in the `Authorization` header (`Bearer <token>`).

* Test protected routes with both admin and user roles.

* Make sure MongoDB is running locally or remotely and properly connected.

---
