# Rubix Store - Clean Architecture Backend

A professional Go backend service built with **Gin**, **PostgreSQL**, **JWT authentication**, and **Clean Architecture** principles.

## Project Structure

```
service/auth/
├── domain/              # Entities & interfaces (core business logic)
│   ├── user.go
│   └── repository.go   # Repository interface
├── repository/          # Data access layer (database)
│   └── postgres.go
├── usecase/            # Business logic layer
│   ├── register.go
│   └── login.go
├── handler/            # HTTP handlers (controllers)
│   ├── auth.go
│   └── auth_test.go
├── middleware/         # HTTP middleware
│   └── jwt.go
├── routes/             # Route registration
│   └── auth.go
└── migration.sql       # Database migration
```

## Architecture Layers

- **Domain**: Pure business entities and interfaces (no dependencies on frameworks)
- **Usecase**: Business logic (registration, login) - depends only on domain
- **Repository**: Data access abstraction (PostgreSQL implementation)
- **Handler**: HTTP request/response handling - depends on usecase
- **Middleware**: Cross-cutting concerns (JWT validation)
- **Routes**: Route registration and dependency injection

## Features

- ✅ User registration with bcrypt password hashing
- ✅ JWT-based authentication
- ✅ Protected routes middleware
- ✅ Unit tests with sqlmock
- ✅ Professional error handling
- ✅ PostgreSQL database

## Setup

### Prerequisites

- Go 1.25+
- PostgreSQL
- Environment variables:
  - `DATABASE_URL`: PostgreSQL connection string
  - `JWT_SECRET`: JWT signing secret
  - `PORT`: Server port (default 8080)

### Database Migration

Create the `users` table:

```sql
CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    email TEXT NOT NULL UNIQUE,
    password_hash TEXT NOT NULL,
    roles TEXT DEFAULT 'user',
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
```

Run migration:
```powershell
psql -U user -d rubixtore -f service/auth/migration.sql
```

### Install & Run

```powershell
cd d:\workspace\learning\golang\rubixtore
go mod tidy
go run main.go
```

## API Endpoints

### Health Check
```bash
GET /health
```

### Register
```bash
POST /auth/register
Content-Type: application/json

{
  "email": "user@example.com",
  "password": "password123",
  "roles": "user"
}
```

**Response (201)**:
```json
{
  "id": 1,
  "email": "user@example.com"
}
```

### Login
```bash
POST /auth/login
Content-Type: application/json

{
  "email": "user@example.com",
  "password": "password123"
}
```

**Response (200)**:
```json
{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
}
```

### Protected Route (Get User Info)
```bash
GET /me
Authorization: Bearer <token>
```

**Response (200)**:
```json
{
  "user_id": 1,
  "email": "user@example.com"
}
```

## Testing

```powershell
go test ./service/auth/... -v
```

Tests use `sqlmock` for database mocking:
- `TestRegister_Success`: Validates user registration
- `TestLogin_Success`: Validates JWT token generation

## Clean Architecture Benefits

- **Testability**: Each layer can be tested independently
- **Maintainability**: Clear separation of concerns
- **Scalability**: Easy to add new features without affecting existing code
- **Flexibility**: Can swap implementations (e.g., different databases)

