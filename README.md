# REST API Project with Go

A production-ready REST API built with Go, Fiber, GORM, and MySQL. Demonstrates clean architecture, CRUD operations, and modern Go development practices.

## 🚀 Features

- **Complete CRUD Operations** for User management
- **Clean Architecture** with Repository and Service patterns
- **Input Validation** using go-playground/validator
- **Password Hashing** with bcrypt
- **Database Migration** support
- **JSON API Responses**
- **Error Handling** with proper HTTP status codes
- **Factory Pattern** for dependency injection
- **GORM ORM** integration with MySQL

## 📋 API Endpoints

### User Management

| Method | Endpoint | Description |
|--------|----------|-------------|
| `GET` | `/api/v1/users` | Get all users |
| `GET` | `/api/v1/users/:id` | Get user by ID |
| `POST` | `/api/v1/users` | Create new user |
| `PUT` | `/api/v1/users/:id` | Update user |
| `DELETE` | `/api/v1/users/:id` | Delete user |

### System Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| `GET` | `/` | Welcome message |
| `GET` | `/health` | Health check |
| `GET` | `/hello` | Hello world |

## 🛠️ Technology Stack

- **Language**: Go 1.23+
- **Web Framework**: [Fiber v2](https://gofiber.io/)
- **ORM**: [GORM v2](https://gorm.io/)
- **Database**: MySQL 8.0+
- **Validation**: [go-playground/validator](https://github.com/go-playground/validator)
- **Password Hashing**: bcrypt
- **Migration**: [golang-migrate](https://github.com/golang-migrate/migrate)

## 📁 Project Structure

```
REST-API-Project-with-GO/
├── cmd/
│   └── rest-api-project-with-go/
│       └── main.go                 # Application entry point
├── internal/
│   ├── config/
│   │   └── config.go              # Configuration management
│   ├── database/
│   │   └── connection.go          # Database connection
│   ├── handlers/
│   │   └── user_handler.go        # HTTP handlers
│   ├── models/
│   │   └── user.go                # Data models
│   ├── repositories/
│   │   └── user_repository.go     # Data access layer
│   ├── routes/
│   │   └── routes.go              # Route definitions
│   └── services/
│       └── user_service.go        # Business logic layer
├── migrations/
│   ├── 000001_create_users_table.up.sql
│   └── 000001_create_users_table.down.sql
├── go.mod
├── go.sum
└── README.md
```

## 🏗️ Architecture

**Clean Architecture** with clear layer separation:

- **Handler Layer**: HTTP request/response handling
- **Service Layer**: Business logic and validation  
- **Repository Layer**: Data access abstraction
- **Model Layer**: Domain entities and DTOs

```
Handler → Service → Repository → Database
```

**Key Design Decisions:**
- Factory pattern for dependency injection
- Interface-based abstractions for testability
- Separation of concerns for maintainability

## ⚙️ Setup & Installation

### Prerequisites

- Go 1.23 or higher
- MySQL 8.0 or higher
- [golang-migrate](https://github.com/golang-migrate/migrate) CLI tool

### 1. Clone the Repository

```bash
git clone https://github.com/Shamim-cse-pust/REST-API-Project-with-GO.git
cd REST-API-Project-with-GO
```

### 2. Install Dependencies

```bash
go mod tidy
```

### 3. Database Setup

Create a MySQL database:

```sql
CREATE DATABASE rest_api_db;
```

### 4. Configuration

Update the database configuration in `internal/config/config.go` or set environment variables:

```go
// Database configuration
DB_HOST=localhost
DB_PORT=3306
DB_USER=your_username
DB_PASSWORD=your_password
DB_NAME=rest_api_db
```

### 5. Run Migrations

```bash
migrate -path migrations -database "mysql://username:password@tcp(localhost:3306)/rest_api_db" up
```

### 6. Run the Application

```bash
go run cmd/rest-api-project-with-go/main.go
```

The server will start on `http://localhost:8082`

## 📡 API Usage Examples

### Create User

```bash
curl -X POST http://localhost:8082/api/v1/users \
  -H "Content-Type: application/json" \
  -d '{
    "name": "John Doe",
    "email": "john@example.com",
    "password": "securepassword123"
  }'
```

### Get All Users

```bash
curl -X GET http://localhost:8082/api/v1/users
```

### Get User by ID

```bash
curl -X GET http://localhost:8082/api/v1/users/1
```

### Update User

```bash
curl -X PUT http://localhost:8082/api/v1/users/1 \
  -H "Content-Type: application/json" \
  -d '{
    "name": "John Smith",
    "email": "johnsmith@example.com"
  }'
```

### Delete User

```bash
curl -X DELETE http://localhost:8082/api/v1/users/1
```

## 📝 Request/Response Examples

### Create User Request

```json
{
  "name": "John Doe",
  "email": "john@example.com",
  "password": "securepassword123"
}
```

### User Response

```json
{
  "message": "User created successfully",
  "data": {
    "id": 1,
    "name": "John Doe",
    "email": "john@example.com",
    "created_at": "2025-11-02T00:00:00Z",
    "updated_at": "2025-11-02T00:00:00Z"
  }
}
```

### Error Response

```json
{
  "error": "Validation failed",
  "message": "Key: 'CreateUserRequest.Email' Error:Field validation for 'Email' failed on the 'required' tag"
}
```

## 👨‍💻 Author

**Md. Shamim Miah**  
*Software Engineer II | Brain Station 23 PLC*

### 📫 Connect With Me
- **Email**: [shamim.pust.cse@gmail.com](mailto:shamim.pust.cse@gmail.com)
- **Phone**: +8801784766676
- **GitHub**: [@Shamim-cse-pust](https://github.com/Shamim-cse-pust)
- **LinkedIn**: [md-shamim-miah](https://www.linkedin.com/in/md-shamim-miah-b01833217/)
- **Portfolio**: [shamim-cse-pust.github.io/myself-shamim](https://shamim-cse-pust.github.io/myself-shamim/)
- **Resume**: [Download PDF](https://shamim-cse-pust.github.io/myself-shamim/resume.pdf)

### 🚀 About Me
Driven by curiosity and a passion for growth, I thrive on solving complex problems and embracing new challenges. My goal is to craft impactful software solutions, collaborate with talented professionals, and continuously evolve—contributing to both organizational success and my own journey of learning and creativity.


⭐ **Star this repository if you found it helpful!**

🔄 **Keep this repository updated in your watch list for new features and improvements.**
