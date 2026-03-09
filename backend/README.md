# Backend Service (Go)

The backend is built with a focus on high throughput and clean code. It is written in **Go** using the **Fiber v3** framework.

## 🏗️ Architecture: Controller-Service-Repository

The codebase follows the industry-standard clean architecture:

1.  **Controller Layer**: Handles incoming HTTP requests, validates inputs (using `go-playground/validator`), and maps responses.
2.  **Service Layer**: Contains the core business logic, orchestrating calls between repositories and other external services.
3.  **Repository Layer**: Encapsulates data access and storage logic (PostgreSQL via GORM).

### Key Components:

- **Auth System**: Custom JWT-based authentication middleware.
- **Distributed IDs (Hyperflake)**: We generate 64-bit distributed IDs using **Hyperflake** with a custom epoch of **June 30, 2025** to keep initial IDs small and efficient.
- **Analytics Streams**: Uses **Redis Streams** to publish link-click events asynchronously, preventing database write latency from slowing down the redirect process.
- **Configuration**: Uses `caarlos0/env` for strict environment variable parsing with sensible defaults.

## 🚀 Running in Development

Ensure you have **Air** installed for live-reloading:
`go install github.com/air-verse/air@latest`

### Running the API (Port 5000)
```bash
make dev-api
```

### Running the Redirect Service (Port 5001)
```bash
make dev-redirect
```

---

## 🛠️ Tech Stack & Libraries

- **Framework**: [Fiber v3](https://docs.gofiber.io/)
- **ORM**: [GORM](https://gorm.io/)
- **Database**: PostgreSQL (pgx driver)
- **Caching/Queue**: Redis (go-redis v9)
- **ID Generation**: [Hyperflake](https://github.com/chirag3003/hyperflake)
- **Logging**: [Zerolog](https://github.com/rs/zerolog)

## 🧪 Testing

The backend includes a comprehensive suite of unit tests for services and repositories:

```bash
# Run all tests
make test

# Generate coverage report
make test-coverage
```

## 📜 API Documentation

Detailed documentation for all v1 endpoints is located in `backend/docs/endpoints/INDEX.md`.
Use the provided `backend/docs/postman.json` to import all endpoints into Postman with pre-configured JWT scripts.
