# Rubix Store

A modern e-commerce API built with Go and Gin framework. Authentication, user profiles, and extensible architecture for cart, orders, and inventory management.

## Quick Start

**Requirements:** Go 1.25.5+, PostgreSQL

```bash
git clone https://github.com/yourusername/rubix-store.git
cd rubix-store
go mod download

# Setup .env file
echo 'DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=your_password
DB_NAME=rubix_store
PORT=8080' > .env

createdb rubix_store
go run cmd/api/main.go
```

Server runs on `http://localhost:8080`

## Tech Stack

Go 1.25.5 • Gin Framework • PostgreSQL • GORM • JWT Authentication

## Project Structure

```
internal/
├── config/       # Database configuration
├── handlers/     # HTTP handlers
├── middleware/   # Auth middleware
├── models/       # Data models
├── repository/   # Database access
├── service/      # Business logic
└── utils/        # JWT utilities
```

## API Routes

All routes start with `/api`. Protected routes require `Authorization: Bearer <token>` header.

| Method | Endpoint | Description | Auth |
|--------|----------|-------------|------|
| POST | `/auth/register` | Register new user | ❌ |
| POST | `/auth/login` | Login & get JWT token | ❌ |
| GET | `/profile` | Get user profile | ✅ |
| PUT | `/profile` | Update user profile | ✅ |

<details>
<summary><b>Request/Response Examples</b></summary>

### Register
```bash
curl -X POST http://localhost:8080/api/auth/register \
  -H "Content-Type: application/json" \
  -d '{
    "email": "user@example.com",
    "password": "password123",
    "first_name": "John",
    "last_name": "Doe"
  }'
```
Response: `201 Created`
```json
{
  "message": "User registered successfully",
  "user": {
    "id": "uuid",
    "email": "user@example.com",
    "first_name": "John",
    "last_name": "Doe"
  }
}
```

### Login
```bash
curl -X POST http://localhost:8080/api/auth/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "user@example.com",
    "password": "password123"
  }'
```
Response: `200 OK`
```json
{
  "message": "Login successful",
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
  "user": {"id": "uuid", "email": "user@example.com"}
}
```

### Get Profile
```bash
curl -X GET http://localhost:8080/api/profile \
  -H "Authorization: Bearer <your_token>"
```
Response: `200 OK`

### Update Profile
```bash
curl -X PUT http://localhost:8080/api/profile \
  -H "Authorization: Bearer <your_token>" \
  -H "Content-Type: application/json" \
  -d '{
    "first_name": "Jane",
    "last_name": "Smith",
    "phone": "+1234567890"
  }'
```

</details>

## Planned Features

- Products (CRUD operations)
- Shopping Cart management
- Order processing
- Admin dashboard
- Inventory management

## Development

```bash
# Run in development mode
go run cmd/api/main.go

# Build for production
go build -o rubix-store cmd/api/main.go

# Run tests
go test ./...
```

## Contributing

1. Fork and create a branch: `git checkout -b feature/your-feature`
2. Make changes and commit: `git commit -m "Add feature"`
3. Push: `git push origin feature/your-feature`
4. Open a Pull Request

## License

MIT License
