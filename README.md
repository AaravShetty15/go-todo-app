# Go Todo API

A RESTful Todo backend API built in Go using a layered architecture. The application supports CRUD operations, pagination, input validation, middleware, and external API integrations.

---

## Features

- CRUD operations for managing todos
- Layered architecture (Handlers → Services → Repository)
- SQLite database
- Pagination support for listing todos
- Input validation
- Standardized JSON API responses
- Logging middleware
- Basic Authentication middleware for protected endpoints
- External API integrations for task suggestions and weather data
- Environment-based configuration using `.env`

---

## Project Structure

```
go-todo-app
│
├── cmd/              # Application entry point
├── config/           # Environment configuration
├── handlers/         # HTTP request handlers
├── middleware/       # Logging and authentication middleware
├── models/           # Data models
├── repository/       # Database access layer
├── routes/           # API route definitions
├── services/         # Business logic
├── external/         # External API integrations
├── utils/            # Helper functions (standard API responses)
```

---

## API Endpoints

### Todos

| Method | Endpoint | Description |
|------|------|------|
| POST | `/todos` | Create a new todo (Auth required) |
| GET | `/todos` | Get all todos (supports pagination) |
| GET | `/todos/{id}` | Get a specific todo |
| PUT | `/todos/{id}` | Update a todo (Auth required) |
| DELETE | `/todos/{id}` | Delete a todo (Auth required) |

### External APIs

| Method | Endpoint | Description |
|------|------|------|
| GET | `/suggest-task` | Get a random task suggestion |
| GET | `/weather` | Get weather information |

---

## Example Request

### Create Todo

```
POST /todos
```

Body:

```json
{
  "title": "Learn Go",
  "description": "Build REST APIs",
  "completed": false
}
```

Response:

```json
{
  "success": true,
  "message": "Todo created successfully",
  "data": {
    "id": 1,
    "title": "Learn Go",
    "description": "Build REST APIs",
    "completed": false,
    "created_at": "2026-03-09T12:00:00Z",
    "updated_at": "2026-03-09T12:00:00Z"
  }
}
```

---

### Get Todos (Pagination)

```
GET /todos?page=1&limit=2
```

Response:

```json
{
  "success": true,
  "data": [
    {
      "id": 1,
      "title": "Learn Go",
      "description": "Build REST APIs",
      "completed": false
    },
    {
      "id": 2,
      "title": "Practice Go",
      "description": "Write more code",
      "completed": true
    }
  ]
}
```

---

### Update Todo

```
PUT /todos/1
```

Body:

```json
{
  "title": "Learn Go Advanced",
  "description": "Practice building APIs",
  "completed": true
}
```

Response:

```json
{
  "success": true,
  "message": "Todo updated successfully",
  "data": {
    "id": 1,
    "title": "Learn Go Advanced",
    "description": "Practice building APIs",
    "completed": true
  }
}
```

---

### Delete Todo

```
DELETE /todos/1
```

Response:

```json
{
  "success": true,
  "message": "Todo deleted successfully"
}
```

---

### Get Task Suggestion

```
GET /suggest-task
```

Response:

```json
{
  "success": true,
  "data": {
    "activity": "Learn basic origami",
    "type": "education"
  }
}
```

---

### Get Weather

```
GET /weather
```

Response:

```json
{
  "success": true,
  "data": {
    "temperature": 27.5,
    "windspeed": 10.2
  }
}
```

---

## Author

Aarav Shetty