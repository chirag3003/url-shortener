# URL Shortener (Full-Stack)

A high-performance, production-ready URL shortener built with a Go (Fiber v3) backend and a Next.js frontend. This project features a distributed ID generation system, real-time analytics via Redis Streams, and a containerized microservices architecture.

## 🚀 Quick Start

The fastest way to get the entire stack running is using Docker Compose:

```bash
docker compose up --build
```

- **Frontend**: [http://localhost:3010](http://localhost:3010)
- **API Service**: [http://localhost:5010](http://localhost:5010)
- **Redirect Service**: [http://localhost:5011](http://localhost:5011)

---

## 🏗️ Architecture & Design Decisions

This project is split into three main components to ensure scalability and low-latency redirects:

1.  **API Service (`backend/cmd/server`)**: Handles user authentication, link management (CRUD), and analytics retrieval.
2.  **Redirect Service (`backend/cmd/redirect`)**: A ultra-lightweight service dedicated solely to resolving short codes and redirecting users.
3.  **Frontend (`frontend/`)**: A modern Next.js application for managing links and viewing analytics.

### Key Architectural Choices:

- **Hyperflake ID Generation**: We use **Hyperflake** (a Snowflake-inspired distributed ID generator) to create unique, 64-bit time-ordered integers for link IDs. This ensures high-performance ID generation across multiple nodes without coordination.
- **Asynchronous Analytics**: To keep redirects as fast as possible, the Redirect Service publishes click metadata to a **Redis Stream** and immediately returns the HTTP 301/302 response. A background worker (integrated or separate) consumes these streams to update PostgreSQL.
- **Service Separation**: The API and Redirect services are separated to allow independent scaling. The Redirect service can be scaled horizontally to handle massive traffic spikes without affecting the management API.

---

## 🛠️ Tech Stack

- **Backend**: Go 1.22+, Fiber v3, GORM (PostgreSQL), Redis (Streams & Caching)
- **Frontend**: Next.js 14 (App Router), Tailwind CSS, TypeScript
- **Infrastructure**: Docker, Docker Compose, PostgreSQL 16, Redis 7

---

## 💻 Local Development

If you prefer to run services individually for real-time coding:

### 1. Infrastructure
```bash
docker compose up postgres redis -d
```

### 2. Backend (API & Redirect)
Requires [Air](https://github.com/air-verse/air) for live-reloading.
```bash
make dev-api      # Terminal 1
make dev-redirect # Terminal 2
```

### 3. Frontend
```bash
make dev-frontend # Terminal 3
```

---

## 📜 Documentation

- **API Reference**: Detailed endpoint documentation can be found in `backend/docs/endpoints/INDEX.md`.
- **Postman**: A pre-configured collection is available at `backend/docs/postman.json`.
