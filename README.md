# Rubix Store

A modern e-commerce API built with Go and Gin framework. Rubix Store provides a robust backend for managing products, orders, inventory, and user authentication.

## Features

- **User Authentication**: Secure registration and login with JWT token-based authentication
- **User Profiles**: Create and update user profile information
- **API Versioning**: Clean API structure with versioned endpoints (`/api/v1`)
- **Protected Routes**: Role-based access control with middleware
- **Database Integration**: PostgreSQL support via GORM ORM
- **Error Handling**: Comprehensive error handling and validation
- **Environment Configuration**: Configurable via `.env` file

## Tech Stack

- **Language**: Go 1.25.5
- **Framework**: Gin Web Framework
- **Database**: PostgreSQL with GORM ORM
- **Authentication**: JWT (JSON Web Tokens)
- **Additional Libraries**:
  - github.com/golang-jwt/jwt/v5 - JWT token generation and validation
  - github.com/joho/godotenv - Environment variable management

## Project Structure

```
rubix-store/
├── cmd/
│   └── api/
│       └── main.go              # Application entry point
├── internal/
│   ├── config/
│   │   └── database.go          # Database configuration and initialization
│   ├── handlers/
│   │   └── auth.go              # HTTP request handlers
│   ├── middleware/
│   │   └── auth.go              # Authentication middleware
│   ├── models/
│   │   └── user.go              # Data models
│   ├── repository/
│   │   └── user_repository.go   # Database access layer
│   ├── service/
│   │   └── auth_service.go      # Business logic
│   └── utils/
│       └── jwt.go               # JWT utilities
├── migrations/                   # Database migrations
├── service/                      # Microservices (future expansion)
│   ├── cart/
│   ├── inventory/
│   ├── notification/
│   ├── order/
│   └── product/
├── go.mod                        # Go module definition
└── .env                          # Environment variables (not committed)
```

## Prerequisites

- Go 1.25.5 or higher
- PostgreSQL database
- Git

## Getting Started

### 1. Clone the Repository

```bash
git clone https://github.com/yourusername/rubix-store.git
cd rubix-store
```

### 2. Install Dependencies

```bash
go mod download
go mod tidy
```

### 3. Environment Configuration

Create a `.env` file in the root directory:

```env
# Database Configuration
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=your_password
DB_NAME=rubix_store

# Server Configuration
PORT=8080

# JWT Configuration
JWT_SECRET=your_secret_key_here
JWT_EXPIRY=24h
```

### 4. Database Setup

Ensure PostgreSQL is running and create the database:

```bash
createdb rubix_store
```

Run migrations (if any):

```bash
# Migration commands here
```

### 5. Run the Application

```bash
go run cmd/api/main.go
```

The server will start on `http://localhost:8080`

## API Documentation

### Base URL

```
http://localhost:8080/api
```

### Authentication Endpoints

#### Register User

Create a new user account.

**Request:**
```http
POST /api/auth/register
Content-Type: application/json

{
  "email": "user@example.com",
  "password": "securepassword123",
  "first_name": "John",
  "last_name": "Doe"
}
```

**Response:**
```json
{
  "message": "User registered successfully",
  "user": {
    "id": "uuid",
    "email": "user@example.com",
    "first_name": "John",
    "last_name": "Doe",
    "created_at": "2026-01-16T00:00:00Z"
  }
}
```

**Status Codes:**
- `201 Created`: User successfully registered
- `400 Bad Request`: Invalid input or validation error
- `409 Conflict`: Email already exists

---

#### Login

Authenticate and receive JWT token.

**Request:**
```http
POST /api/auth/login
Content-Type: application/json

{
  "email": "user@example.com",
  "password": "securepassword123"
}
```

**Response:**
```json
{
  "message": "Login successful",
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
  "user": {
    "id": "uuid",
    "email": "user@example.com",
    "first_name": "John",
    "last_name": "Doe"
  }
}
```

**Status Codes:**
- `200 OK`: Login successful
- `400 Bad Request`: Invalid credentials or validation error
- `401 Unauthorized`: Invalid email or password

---

### Protected Routes

All protected routes require the JWT token in the Authorization header:

```
Authorization: Bearer <your_jwt_token>
```

#### Get User Profile

Retrieve the authenticated user's profile information.

**Request:**
```http
GET /api/profile
Authorization: Bearer <your_jwt_token>
```

**Response:**
```json
{
  "id": "uuid",
  "email": "user@example.com",
  "first_name": "John",
  "last_name": "Doe",
  "phone": "+1234567890",
  "created_at": "2026-01-16T00:00:00Z",
  "updated_at": "2026-01-16T00:00:00Z"
}
```

**Status Codes:**
- `200 OK`: Profile retrieved successfully
- `401 Unauthorized`: Missing or invalid token
- `404 Not Found`: User not found

---

#### Update User Profile

Update the authenticated user's profile information.

**Request:**
```http
PUT /api/profile
Authorization: Bearer <your_jwt_token>
Content-Type: application/json

{
  "first_name": "Jane",
  "last_name": "Smith",
  "phone": "+9876543210"
}
```

**Response:**
```json
{
  "message": "Profile updated successfully",
  "user": {
    "id": "uuid",
    "email": "user@example.com",
    "first_name": "Jane",
    "last_name": "Smith",
    "phone": "+9876543210",
    "updated_at": "2026-01-16T10:30:00Z"
  }
}
```

**Status Codes:**
- `200 OK`: Profile updated successfully
- `400 Bad Request`: Invalid input
- `401 Unauthorized`: Missing or invalid token
- `404 Not Found`: User not found

---

## Future Endpoints (Planned)

The following endpoints are planned for future implementation:

### Products
- `GET /api/products` - List all products
- `GET /api/products/:id` - Get product details
- `POST /api/admin/products` - Create product (admin only)
- `PUT /api/admin/products/:id` - Update product (admin only)
- `DELETE /api/admin/products/:id` - Delete product (admin only)

### Cart
- `GET /api/cart` - Get user's cart
- `POST /api/cart/items` - Add item to cart
- `PUT /api/cart/items/:id` - Update cart item
- `DELETE /api/cart/items/:id` - Remove item from cart

### Orders
- `POST /api/orders` - Create order
- `GET /api/orders` - Get user's orders
- `GET /api/orders/:id` - Get order details
- `PUT /api/orders/:id/cancel` - Cancel order

### Admin
- `GET /api/admin/users` - List all users (admin only)
- `GET /api/admin/orders` - List all orders (admin only)

## Error Handling

All endpoints return consistent error responses:

```json
{
  "error": "Error message describing what went wrong"
}
```

Common HTTP status codes:
- `200 OK` - Successful request
- `201 Created` - Resource successfully created
- `400 Bad Request` - Invalid request data
- `401 Unauthorized` - Authentication required or failed
- `404 Not Found` - Resource not found
- `409 Conflict` - Resource already exists
- `500 Internal Server Error` - Server error

## Testing with cURL

### Register a new user:
```bash
curl -X POST http://localhost:8080/api/auth/register \
  -H "Content-Type: application/json" \
  -d '{
    "email": "test@example.com",
    "password": "password123",
    "first_name": "Test",
    "last_name": "User"
  }'
```

### Login:
```bash
curl -X POST http://localhost:8080/api/auth/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "test@example.com",
    "password": "password123"
  }'
```

### Get profile (replace TOKEN with actual JWT):
```bash
curl -X GET http://localhost:8080/api/profile \
  -H "Authorization: Bearer TOKEN"
```

### Update profile:
```bash
curl -X PUT http://localhost:8080/api/profile \
  -H "Authorization: Bearer TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "first_name": "Updated",
    "last_name": "Name",
    "phone": "+1234567890"
  }'
```

## Development

### Running in Development Mode

```bash
go run cmd/api/main.go
```

### Building for Production

```bash
go build -o rubix-store cmd/api/main.go
./rubix-store
```

### Running Tests

```bash
go test ./...
```

## Contributing

1. Create a new branch for your feature
   ```bash
   git checkout -b feature/your-feature-name
   ```

2. Make your changes and commit them
   ```bash
   git add .
   git commit -m "Add your feature"
   ```

3. Push to the branch
   ```bash
   git push origin feature/your-feature-name
   ```

4. Open a Pull Request

## License

This project is licensed under the MIT License - see the LICENSE file for details.

## Support

For issues, questions, or suggestions, please open an issue on GitHub.

---

**Last Updated**: January 16, 2026
