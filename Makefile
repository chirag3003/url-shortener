# Variables
BACKEND_DIR := backend
FRONTEND_DIR := frontend
DOCKER_COMPOSE := docker compose

.PHONY: help install build test lint tidy clean dev-api dev-redirect dev-frontend up down logs ps migrate-up migrate-down migrate-status

# Default help command
help:
	@echo "Available commands:"
	@echo "  install        - Install dependencies for both backend and frontend"
	@echo "  build          - Build backend binary and frontend production bundle"
	@echo "  test           - Run all backend tests with race detection"
	@echo "  test-coverage  - Run backend tests and generate HTML coverage report"
	@echo "  lint           - Run golangci-lint on the backend"
	@echo "  tidy           - Tidy Go modules"
	@echo "  dev-api        - Run API service with Air (port 5000)"
	@echo "  dev-redirect   - Run Redirect service with Air (port 5001)"
	@echo "  dev-frontend   - Run Frontend (Next.js)"
	@echo "  up             - Start all services with Docker Compose (detached)"
	@echo "  down           - Stop all services and remove containers"
	@echo "  logs           - Follow Docker Compose logs"
	@echo "  ps             - Show status of Docker Compose services"
	@echo "  migrate-up     - Apply all pending DB migrations"
	@echo "  migrate-down   - Roll back one DB migration"
	@echo "  migrate-status - Show DB migration status"
	@echo "  clean          - Remove build artifacts and temporary files"

# --- Setup & Maintenance ---

install:
	cd $(BACKEND_DIR) && go mod tidy
	cd $(FRONTEND_DIR) && npm install

tidy:
	cd $(BACKEND_DIR) && go mod tidy

# --- Development ---

dev-api:
	cd $(BACKEND_DIR) && air

dev-redirect:
	cd $(BACKEND_DIR) && air -c .air.redirect.toml

dev-frontend:
	cd $(FRONTEND_DIR) && npm run dev

# --- Build & Test ---

build-backend:
	go build -o ./$(BACKEND_DIR)/bin/server ./$(BACKEND_DIR)/cmd/server

build-frontend:
	cd $(FRONTEND_DIR) && npm run build

build: build-backend build-frontend

test:
	go test -v -race -cover ./$(BACKEND_DIR)/...

test-coverage:
	go test -v -race -coverprofile=$(BACKEND_DIR)/coverage.out ./$(BACKEND_DIR)/...
	go tool cover -html=$(BACKEND_DIR)/coverage.out -o $(BACKEND_DIR)/coverage.html

lint:
	golangci-lint run ./$(BACKEND_DIR)/...

# --- Docker Orchestration ---

up:
	$(DOCKER_COMPOSE) up --build -d

down:
	$(DOCKER_COMPOSE) down

logs:
	$(DOCKER_COMPOSE) logs -f

ps:
	$(DOCKER_COMPOSE) ps

# --- Database Migrations ---

migrate-up:
	cd $(BACKEND_DIR) && go run ./cmd/migrate up

migrate-down:
	cd $(BACKEND_DIR) && go run ./cmd/migrate down

migrate-status:
	cd $(BACKEND_DIR) && go run ./cmd/migrate status

# --- Cleanup ---

clean:
	rm -rf ./$(BACKEND_DIR)/bin ./$(BACKEND_DIR)/tmp $(BACKEND_DIR)/coverage.out $(BACKEND_DIR)/coverage.html
	rm -rf ./$(FRONTEND_DIR)/.next ./$(FRONTEND_DIR)/out
