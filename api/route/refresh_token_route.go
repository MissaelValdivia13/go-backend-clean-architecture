# Go Clean Architecture Backend

A production-ready RESTful API boilerplate developed in **Go (Golang)** implementing **Clean Architecture** and **SOLID** principles. This project showcases decoupled layers, robust middleware integration, automated testing, and containerized deployment.

## 🏗️ Architectural Layers

This project strictly adheres to separation of concerns by splitting the codebase into five distinct layers, ensuring high testability and maintainability:

* **Domain:** Contains the core business entities, enterprise logic, and repository/use-case interface definitions. This layer has zero external dependencies.
* **Usecase (Services):** Implements the specific business use cases and coordinates the data flow between controllers and repositories.
* **Repository:** Handles data persistence logic and direct communication with the database.
* **Controller:** Manages HTTP request parsing, payload validation, and structures the API responses.
* **Router:** Sets up the HTTP endpoints and attaches global or route-specific middlewares.

---

## 🛠️ Tech Stack & Key Packages

* **Core:** [Go (Golang)](https://go.dev/) - High-performance, concurrent language.
* **Web Framework:** [Gin Gonic](https://github.com/gin-gonic/gin) - Fast, lightweight HTTP web framework for performance-critical APIs.
* **Database:** [MongoDB Official Go Driver](https://github.com/mongodb/mongo-go-driver) - NoSQL data persistence layer.
* **Security & Auth:** [JWT-Go](https://github.com/golang-jwt/jwt) - Secure access and refresh token authentication flow; [Bcrypt](https://pkg.go.dev/golang.org/x/crypto/bcrypt) for secure password hashing.
* **Configuration:** [Viper](https://github.com/spf13/viper) - Environment variable and configuration management supporting multiple formats.
* **Testing & Mocks:** [Testify](https://github.com/stretchr/testify) assertions and [Mockery](https://github.com/vektra/mockery) for automated mock generation.

---

## 🔒 API Request Flow

### Public Endpoints
`Client Request` ➔ `Router` ➔ `Controller` ➔ `Usecase` ➔ `Repository` ➔ `MongoDB`

### Private Endpoints (Secured via JWT)
`Client Request` ➔ `Router` ➔ **`JWT Auth Middleware`** ➔ `Controller` ➔ `Usecase` ➔ `Repository` ➔ `MongoDB`

---

## 🚀 Getting Started

### Prerequisites
* Go (1.21+ recommended)
* MongoDB instance (local or Atlas) OR Docker installed

### Setup Environment
Clone the repository and create your configuration file:
```bash
cp .env.example .env```

###Run Locally (Without Docker)
# Install dependencies
go mod download

# Run the application
go run cmd/main.go

###Run with Docker Compose
docker-compose up -d

###Running Unit Tests
go test ./... -v

###Regenerating Mocks
# Generate core domain mocks
mockery --dir=domain --output=domain/mocks --outpkg=mocks --all

# Generate database client mocks
mockery --dir=mongo --output=mongo/mocks --outpkg=mocks --all


Project Structure
.
├── cmd/                  # Application entry point (main.go)
├── api/
│   ├── controller/       # HTTP request handlers & validation
│   ├── middleware/       # JWT authentication & security filters
│   └── route/            # Route grouping & endpoint mappings
├── bootstrap/            # App initialization, config loading, and DB connection
├── domain/               # Core entities, interface blueprints, and models
├── internal/             # Private application utilities (token helpers)
├── mongo/                # Native MongoDB client setup
├── repository/           # Data access implementations
├── usecase/              # Core business logic implementations
├── Dockerfile
└── docker-compose.yaml

API Endpoints Reference
1. Authentication & Profile
POST /signup - Register a new user. Returns access and refresh tokens.

POST /login - Authenticate credentials. Returns access and refresh tokens.

POST /refresh - Renew an expired access token using a valid refresh token.

GET /profile (Protected) - Retrieve logged-in user details.

2. Task Management (Protected)
POST /task - Create a new task item.

GET /task - Fetch all tasks associated with the authenticated user.