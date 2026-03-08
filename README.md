# URL Shortener

A production-ready URL shortener with a Go (Fiber v3) backend and a Next.js frontend.

## Features

- **Backend**: Go (Fiber v3), PostgreSQL, Redis, JWT Authentication, Analytics.
- **Frontend**: Next.js (App Router), Tailwind CSS.
- **Infrastructure**: Docker Compose, PostgreSQL, Redis.

## Development Setup

To run the individual services in development mode with live-reloading:

### 1. Prerequisites

Ensure you have the following installed:
- [Go](https://golang.org/doc/install) (1.21+)
- [Node.js](https://nodejs.org/en/download/) (18+)
- [Docker](https://docs.docker.com/get-docker/)
- [Air](https://github.com/cosmtrek/air) (for Go live-reloading)

### 2. Run Infrastructure (Databases)

Start only the supporting services (PostgreSQL and Redis) in the background:

```bash
docker compose up postgres redis -d
```

### 3. Start Backend Services

In separate terminal tabs, run the following to start the backend with live-reloading:

**API Service (Port 5000):**
```bash
make dev-api
```

**Redirect Service (Port 5001):**
```bash
make dev-redirect
```

### 4. Start Frontend

In another terminal tab, start the Next.js development server:

```bash
make dev-frontend
```

## Makefile Commands

| Command | Description |
| :--- | :--- |
| `make install` | Install all dependencies (Go + NPM) |
| `make build` | Build backend binary and frontend bundle |
| `make test` | Run all backend tests |
| `make up` | Start full environment in Docker |
| `make down` | Stop and remove all containers |
| `make logs` | Follow Docker logs |
| `make migrate-up`| Apply DB migrations |

## Documentation

- **Backend API**: See `backend/docs/endpoints/INDEX.md`
- **Postman Collection**: `backend/docs/postman.json`
