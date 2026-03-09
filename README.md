# Go Todo API

A RESTful Todo backend API built in Go using a layered architecture. The application supports CRUD operations, pagination, input validation, middleware, and external API integrations.

---

## Features

- CRUD operations for managing todos
- Layered architecture (Handlers → Services → Repository)
- SQLite database for persistence
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

## Pagination Example

```
GET /todos?page=1&limit=5
```

---

## Environment Configuration

Create a `.env` file in the root directory:

```
PORT=8080
DB_PATH=./todos.db
AUTH_USER=admin
AUTH_PASS=password
```

---

## Installation

### 1. Clone the repository

```
git clone https://github.com/YOUR_USERNAME/go-todo-app.git
cd go-todo-app
```

### 2. Install dependencies

```
go mod tidy
```

### 3. Run the application

```
go run cmd/main.go
```

Server will start on:

```
http://localhost:8080
```

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

---

## Tech Stack

- Go
- Gorilla Mux (Router)
- SQLite
- godotenv

---

## Author

Aarav Shetty