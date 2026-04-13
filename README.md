# Go Gin Template

A production-ready Go REST API template built with **Clean Architecture** and **Hexagonal Architecture (Ports & Adapters)** principles.

## 🎯 Overview

This template provides a solid foundation for building scalable microservices with:
- ✅ **Clean Architecture** - Clear separation of concerns
- ✅ **Hexagonal Architecture** - Ports & Adapters pattern
- ✅ **Domain-Driven Design** - Organized by feature modules
- ✅ **REST API** - HTTP endpoints with Gin framework
- ✅ **gRPC Client Support** - Ready for microservices communication
- ✅ **Database & Caching** - MySQL with GORM, Redis
- ✅ **Authentication** - JWT & API Key support
- ✅ **API Documentation** - Swagger/OpenAPI

## 🏗️ Architecture

### **Layers**

```
┌─────────────────────────────────────────────────────────┐
│                    Presentation Layer                   │
│              (internal/api - HTTP/REST)                 │
│  ┌──────────┐  ┌──────────┐  ┌────────────────────┐     │
│  │ Handlers │  │  Router  │  │    Middleware      │     │
│  │          │  │          │  │ (Auth, Logging)    │     │
│  └──────────┘  └──────────┘  └────────────────────┘     │
└─────────────────────────────────────────────────────────┘
                          ↓
┌─────────────────────────────────────────────────────────┐
│                     Domain Layer                        │
│              (internal/core/modules)                    │
│  ┌──────────────────────────────────────────────────┐   │
│  │  Ports (Input/Output) - Interfaces               │   │
│  ├──────────────────────────────────────────────────┤   │
│  │  Services - Business Logic                       │   │
│  ├──────────────────────────────────────────────────┤   │
│  │  Domain - Entities, Models, DTOs                 │   │
│  └──────────────────────────────────────────────────┘   │
└─────────────────────────────────────────────────────────┘
                          ↓
┌─────────────────────────────────────────────────────────┐
│                  Infrastructure Layer                   │
│  ┌──────────────┐  ┌──────────────┐  ┌─────────────┐    │
│  │ Repositories │  │   Database   │  │ gRPC Clients│    │
│  │   (MySQL)    │  │ Connections  │  │  (External) │    │
│  └──────────────┘  └──────────────┘  └─────────────┘    │
└─────────────────────────────────────────────────────────┘
```

### **Request Flow**

```
HTTP Request
    ↓
Router → Middleware (Auth, Logging) → Handler
    ↓
Input Port (Use Case Interface)
    ↓
Service (Business Logic)
    ↓
Output Port (Repository Interface) + gRPC Clients
    ↓
Repository Implementation → Database
```

## 📁 Project Structure

```
go-gin-template/
├── cmd/
│   ├── app/                      # HTTP server entry point
│   │   └── main.go
│   └── application/              # Dependency injection & wiring
│       ├── handler.go
│       ├── repositories.go
│       ├── router.go
│       └── services.go
│
├── internal/
│   ├── config/                   # Configuration management
│   │   ├── app.go
│   │   └── gin.go
│   │
│   ├── api/                      # Presentation layer
│   │   ├── handler/              # HTTP request handlers
│   │   ├── router/               # Route definitions
│   │   └── middleware/           # HTTP middleware
│   │       ├── auth/             # JWT & API Key auth
│   │       ├── logging/          # Request logging
│   │       └── response/         # Response helpers
│   │
│   ├── core/                     # Business logic (Domain)
│   │   └── modules/              # Feature modules
│   │       └── users/            # Example: User module
│   │           ├── domain/       # Domain layer
│   │           │   ├── entities/ # Database entities (GORM)
│   │           │   ├── models/   # Business models
│   │           │   └── dto/      # Data transfer objects
│   │           ├── services/     # Business logic
│   │           ├── repository/   # Repository interfaces
│   │           └── ports/        # Hexagonal architecture
│   │               ├── input/    # Use case interfaces
│   │               └── output/   # Repository interfaces
│   │
│   ├── repository/               # Repository implementations
│   │   ├── mysql/                # MySQL repositories
│   │   └── redis/                # Redis cache
│   │
│   ├── database/                 # Database connections
│   │   ├── mysql/
│   │   └── redis/
│   │
│   └── grpc-clients/             # gRPC client layer (optional)
│       ├── proto/                # Proto definitions
│       ├── {service}/            # Service clients
│       └── interceptors/         # Client middleware
│
├── pkg/                          # Shared utilities
│   ├── logger/                   # Logging utilities
│   ├── utils/                    # General utilities
│   │   ├── datetime/
│   │   ├── encryption/
│   │   ├── format/
│   │   └── filestore/
│   ├── constants/                # Application constants
│   └── swagger/                  # Swagger documentation
│
├── config.yaml                   # Configuration file
├── Makefile                      # Build automation
├── go.mod
└── go.sum
```

## 🚀 Features

### **Architecture Patterns**
- ✅ **Clean Architecture** - Independent of frameworks, UI, database
- ✅ **Hexagonal Architecture** - Ports & Adapters for flexibility
- ✅ **Domain-Driven Design** - Feature-based module organization
- ✅ **Dependency Injection** - Loose coupling, easy testing

### **API & Documentation**
- ✅ **RESTful API** - Standard HTTP endpoints
- ✅ **Swagger/OpenAPI** - Auto-generated API documentation
- ✅ **Request Validation** - Input validation middleware
- ✅ **Error Handling** - Standardized error responses

### **Authentication & Security**
- ✅ **JWT Authentication** - Secure token-based auth
- ✅ **API Key Support** - Alternative authentication method
- ✅ **Middleware Chain** - Auth, logging, recovery, CORS

### **Database & Caching**
- ✅ **MySQL with GORM** - ORM for database operations
- ✅ **Redis Caching** - High-performance caching layer
- ✅ **Connection Pooling** - Optimized database connections
- ✅ **Repository Pattern** - Abstracted data access

### **Microservices Ready**
- ✅ **gRPC Client Support** - Call external services
- ✅ **Service Discovery** - Ready for service mesh
- ✅ **Circuit Breaker** - Resilient service calls
- ✅ **Retry Logic** - Automatic retry with backoff

## 🛠️ Tech Stack

- **Language**: Go 1.24+
- **Framework**: Gin (HTTP router)
- **Database**: MySQL 8.0+ with GORM
- **Cache**: Redis 6.0+
- **Authentication**: JWT (kataras/jwt)
- **Documentation**: Swagger/OpenAPI
- **Logging**: Structured logging

## 📋 Prerequisites

- Go 1.24 or higher
- MySQL 8.0 or higher
- Redis 6.0 or higher
- Docker (optional, for containerization)

## 🚀 Quick Start

### 1. Clone the Repository
```bash
git clone <repository-url>
cd go-gin-template
```

### 2. Install Dependencies
```bash
go mod download
```

### 3. Configure Application
Edit `config.yaml` with your settings:
```yaml
server:
  http:
    port: 8080

database:
  mysql:
    host: localhost
    port: 3306
    username: root
    password: password
    database: mydb
  redis:
    host: localhost
    port: 6379

jwt:
  secret: your-secret-key
  refresh_secret: your-refresh-secret
```

### 4. Run Database Migrations
```bash
# Create your database first
mysql -u root -p -e "CREATE DATABASE mydb;"

# Run the application (auto-migration enabled)
go run ./cmd/app
```

### 5. Start the Server
```bash
go run ./cmd/app
```

The server will start on `http://localhost:8080`

### 6. Access API Documentation
Open your browser and navigate to:
```
http://localhost:8080/swagger/index.html
```

## 📚 API Endpoints

### **Health Check**
```
GET /health
```

### **User Management**
```
POST   /api/v1/users/register    # Register new user
POST   /api/v1/users/login       # User login
GET    /api/v1/users/profile     # Get user profile (requires auth)
PUT    /api/v1/users/profile     # Update profile (requires auth)
```

## 🧪 Testing

```bash
# Run all tests
go test ./...

# Run tests with coverage
go test -v -coverprofile=coverage.out ./...
go tool cover -html=coverage.out
```

## 🐳 Docker

### Build Docker Image
```bash
docker build -t go-gin-template:latest .
```

### Run with Docker Compose
```bash
docker-compose up -d
```

## 📖 Documentation

For detailed architecture documentation, see:
- **[gRPC Client Architecture](./grpc-client-design.md)** - Hybrid REST + gRPC client pattern
- **[gRPC Server Architecture](./grpc-server-design.md)** - Pure gRPC server pattern

## 🔧 Development

### Generate Swagger Documentation
```bash
swag init -g cmd/app/main.go
```

### Format Code
```bash
go fmt ./...
```

### Run Linter
```bash
golangci-lint run
```

## 📝 Adding a New Feature Module

1. Create module structure:
```bash
mkdir -p internal/core/modules/orders/{domain/{entities,models,dto},services,repository,ports/{input,output}}
```

2. Define domain models in `domain/`
3. Create service interfaces in `ports/input/`
4. Implement business logic in `services/`
5. Define repository interface in `ports/output/`
6. Implement repository in `internal/repository/mysql/`
7. Create HTTP handlers in `internal/api/handler/`
8. Register routes in `internal/api/router/`

## 🤝 Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## 📄 License

Apache 2.0

## 👥 Authors

Your Team

---

**Built with ❤️ using Clean Architecture principles**
