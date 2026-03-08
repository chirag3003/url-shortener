.PHONY: build run dev test lint docker-build docker-run clean tidy mocks migrate-up migrate-down migrate-status

BACKEND_DIR := backend

# Build the application binary
build:
	go build -o ./$(BACKEND_DIR)/bin/server ./$(BACKEND_DIR)

# Run the application locally
run: build
	./$(BACKEND_DIR)/bin/server

# Run with Air (live reload)
dev:
	cd $(BACKEND_DIR) && air

# Run tests
test:
	go test -v -race -cover ./$(BACKEND_DIR)/...

# Run tests with coverage report
test-coverage:
	go test -v -race -coverprofile=$(BACKEND_DIR)/coverage.out ./$(BACKEND_DIR)/...
	go tool cover -html=$(BACKEND_DIR)/coverage.out -o $(BACKEND_DIR)/coverage.html

# Run linter
lint:
	golangci-lint run ./$(BACKEND_DIR)/...

# Tidy go modules
tidy:
	cd $(BACKEND_DIR) && go mod tidy

# Build Docker image
docker-build:
	docker build -t go-backend ./$(BACKEND_DIR)

# Run with Docker Compose
docker-run:
	docker compose up --build -d

# Stop Docker Compose services
docker-stop:
	docker compose down

# Clean build artifacts
clean:
	rm -rf ./$(BACKEND_DIR)/bin ./$(BACKEND_DIR)/tmp $(BACKEND_DIR)/coverage.out $(BACKEND_DIR)/coverage.html

# Generate mocks
mocks:
	go generate ./$(BACKEND_DIR)/repository/...

# Apply all pending DB migrations
migrate-up:
	cd $(BACKEND_DIR) && go run ./cmd/migrate up

# Roll back one DB migration
migrate-down:
	cd $(BACKEND_DIR) && go run ./cmd/migrate down

# Show DB migration status
migrate-status:
	cd $(BACKEND_DIR) && go run ./cmd/migrate status
