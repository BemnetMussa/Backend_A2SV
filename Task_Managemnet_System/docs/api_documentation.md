
# 📄 Task Management REST API Documentation

## 👨‍💻 Base URL

[http://localhost:8080](http://localhost:8080)
## 📌 Endpoints Summary

| Method | Endpoint         | Description               |
|--------|------------------|---------------------------|
| GET    | `/tasks`         | Get all tasks             |
| GET    | `/tasks/:id`     | Get task by ID            |
| POST   | `/tasks`         | Create a new task         |
| PUT    | `/tasks/:id`     | Update an existing task   |
| DELETE | `/tasks/:id`     | Delete a task by ID       |

---

## 📥 POST `/tasks`

### Description
Create a new task.

### Request Body (JSON)

```json
{
  "title": "Finish project",
  "description": "Implement all features",
  "completed": false
}
````

### Success Response

* **Code:** `200 OK`

```json
{
  "message": "Task created successfully",
  "task": {
    "id": 1,
    "title": "Finish project",
    "description": "Implement all features",
    "completed": false
  }
}
```

### Error Response

* **Code:** `400 Bad Request`

```json
{
  "error": "Invalid task data"
}
```

---

## 📤 GET `/tasks`

### Description

## Fetch all tasks.

### Success Response

* **Code:** `200 OK`

```json
[
  {
    "id": 1,
    "title": "Finish project",
    "description": "Implement all features",
    "completed": false
  }
]
```

---

## 🔍 GET `/tasks/:id`

### Description

## Get details of a specific task by ID.

### Success Response

* **Code:** `200 OK`

```json
{
  "id": 1,
  "title": "Finish project",
  "description": "Implement all features",
  "completed": false
}
```

### Error Response

* **Code:** `404 Not Found`

```json
{
  "error": "Task not found!"
}
```

---

## ✏️ PUT `/tasks/:id`

### Description

Update an existing task.

### Request Body (JSON)

```json
{
  "title": "Finish final draft",
  "description": "Add error handling",
  "completed": true
}
```

### Success Response

* **Code:** `200 OK`

```json
{
  "task": {
    "id": 1,
    "title": "Finish final draft",
    "description": "Add error handling",
    "completed": true
  }
}
```

### Error Response

* **Code:** `404 Not Found`

```json
{
  "error": "Task not found!"
}
```

---

## ❌ DELETE `/tasks/:id`

### Description

Delete a task by ID.

### Success Response

* **Code:** `200 OK`

```json
{
  "message": "Task deleted successfully"
}
```

### Error Response

* **Code:** `404 Not Found`

```json
{
  "error": "Task not found"
}
```

---

## ✅ Testing Tips (Postman)

* Set `Content-Type` to `application/json` for `POST` and `PUT` requests.
* Use raw JSON body for request payloads.
* Be sure to create a task first before testing `GET`, `PUT`, or `DELETE` on `/tasks/:id`.

---

## 📁 Notes

* This API uses **in-memory** storage, meaning all data will reset when the server restarts.
* Proper HTTP status codes and error messages are returned for all requests.

```

