# Go Backend Template

A production-ready, highly structured Go backend template using [Fiber v3](https://docs.gofiber.io/) and [MongoDB](https://www.mongodb.com/). This template follows industry best practices for security, architecture, and maintainability, making it an ideal starting point for scalable applications.

## Features

- **Framework:** [Fiber v3 (beta)](https://docs.gofiber.io/) — Fast, Express-inspired web framework for Go.
- **Architecture:** Clean Layered Architecture (Controller → Service → Repository).
- **Database:** MongoDB (using official [mongo-go-driver v2](https://pkg.go.dev/go.mongodb.org/mongo-driver/v2)).
- **Authentication:** JWT (JSON Web Token) with secure password hashing (bcrypt).
- **File Uploads:** AWS S3 integration with multi-file support.
- **Image Optimization:** Automatic image optimization and conversion to WebP using [libvips](https://www.libvips.org/).
- **Validation:** Robust input validation using [go-playground/validator](https://github.com/go-playground/validator).
- **Logging:** Structured logging with [zerolog](https://github.com/rs/zerolog).
- **Error Handling:** Centralized error handling with custom `AppError` types and HTTP status mapping.
- **Testing:** Comprehensive test suite (45+ tests) using `go test` and auto-generated mocks with [gomock](https://github.com/uber-go/mock).
- **Containerization:** Multi-stage `Dockerfile` and `docker-compose.yml` for production-ready deployment.
- **CI/CD:** GitHub Actions workflow for linting, testing, and build verification.

## Folder Structure

```text
.
├── main.go                  # Entrypoint: dependency injection, server start, graceful shutdown
├── Makefile                 # Automation for build, run, test, lint, and mocks
├── Dockerfile               # Multi-stage production Docker build
├── compose.yml              # Docker Compose for production/staging environments
├── deploy.sh                # Deployment script with health checks
├── .air.toml                # Configuration for live-reloading (Air)
├── .env.example             # Template for environment variables
├── .golangci.yml            # Linter configuration (golangci-lint)
│
├── config/                  # Configuration management (env parsing, defaults)
├── controller/              # HTTP layer: parses requests, delegates to services
├── db/                      # Database connection and index management
├── dto/                     # Data Transfer Objects (Request/Response shapes)
├── helpers/                 # App-specific utilities (AWS setup, image processing)
├── middleware/              # HTTP middlewares (auth, logging, error handling, etc.)
├── models/                  # Database document models
├── pkg/                     # Reusable, self-contained internal packages
│   ├── apperror/            # Structured application errors
│   ├── auth/                # JWT and password hashing services
│   ├── logger/              # Global structured logger setup
│   └── validate/            # Input validation wrapper
├── repository/              # Data Access Layer: MongoDB and S3 operations
│   └── mock/                # Auto-generated mocks for testing
├── routes/                  # API route definitions and grouping
└── service/                 # Business Logic Layer: core application rules
```

## Prerequisites

- **Go 1.23+**
- **MongoDB** (running locally or via Docker)
- **libvips-dev** (required for image optimization)
- **Optional Tools:**
  - [Air](https://github.com/air-verse/air) (for live reloading)
  - [mockgen](https://github.com/uber-go/mock) (for generating test mocks)
  - [golangci-lint](https://golangci-lint.run/) (for linting)

## Getting Started

1. **Clone the repository:**
   ```bash
   git clone https://github.com/chirag3003/go-backend-template.git
   cd go-backend-template
   ```

2. **Setup environment variables:**
   ```bash
   cp .env.example .env
   # Edit .env with your specific configuration (S3 keys, Mongo URI, etc.)
   ```

3. **Install dependencies:**
   ```bash
   go mod tidy
   ```

4. **Run the application:**
   ```bash
   make run
   # OR for live-reloading:
   make dev
   ```
   The server will start at `http://localhost:5000`.

## Environment Variables

| Variable | Description | Default |
|---|---|---|
| `PORT` | Port for the server to listen on | `5000` |
| `MONGO_URI` | MongoDB connection string | `mongodb://localhost:27017` |
| `MONGO_DB` | MongoDB database name | `go-template` |
| `JWT_SECRET` | Secret key for signing JWTs | (Required) |
| `JWT_EXPIRATION`| Token validity duration (e.g., 24h, 1h) | `24h` |
| `S3_ACCESS_KEY` | AWS IAM Access Key | (Required for S3) |
| `S3_SECRET_KEY` | AWS IAM Secret Key | (Required for S3) |
| `S3_REGION` | AWS S3 Region | `ap-south-1` |
| `S3_BUCKET` | AWS S3 Bucket Name | (Required for S3) |
| `S3_ENDPOINT` | S3 API Endpoint | (Required for S3) |
| `LOG_LEVEL` | Logging level (debug, info, warn, error) | `info` |

## Make Targets

| Target | Description |
|---|---|
| `make build` | Build the server binary to `./bin/server` |
| `make run` | Build and run the application locally |
| `make dev` | Run the application with live-reloading (Air) |
| `make test` | Run all tests with race detection |
| `make test-coverage`| Run tests and generate HTML coverage report |
| `make lint` | Run the linter (golangci-lint) |
| `make mocks` | Regenerate test mocks from repository interfaces |
| `make docker-run` | Start the app and MongoDB via Docker Compose |
| `make tidy` | Tidy Go modules |
| `make clean` | Remove build artifacts and temporary files |

## API Endpoints

### Public Endpoints
- `GET  /health` — Check server health status
- `GET  /ready` — Check database and AWS connectivity
- `POST /api/v1/auth/register` — Create a new user account
- `POST /api/v1/auth/login` — Authenticate and receive a JWT token

### Protected Endpoints (Requires `Authorization: Bearer <token>`)
- `GET  /api/v1/user/me` — Retrieve current authenticated user's profile
- `POST /api/v1/media/upload` — Upload and optimize multiple images to S3

## Testing

This project uses `go test` for unit and integration testing. Mocking is handled by `go.uber.org/mock`.

- **Run all tests:** `make test`
- **Regenerate mocks:** `make mocks`
- **Check coverage:** `make test-coverage`

## Docker

Build and run the entire stack (app + MongoDB) using Docker Compose:
```bash
make docker-run
```
To stop the services:
```bash
make docker-stop
```

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## Author

**Chirag Bhalotia**
- GitHub: [@chirag3003](https://github.com/chirag3003)

## License

This project is licensed under the MIT License - see the LICENSE file for details.
