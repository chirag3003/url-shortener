# Build stage
FROM golang:1.25-bookworm AS builder

# Install libvips for image optimization
RUN apt-get update && apt-get install -y libvips-dev && rm -rf /var/lib/apt/lists/*

WORKDIR /app

# Copy go mod files and download dependencies (cached layer)
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build the server binary
RUN CGO_ENABLED=1 go build -o /app/server ./cmd/server

# Build the redirect binary
RUN CGO_ENABLED=1 go build -o /app/redirect ./cmd/redirect

# Runtime stage
FROM debian:bookworm-slim

# Install runtime dependencies for libvips
RUN apt-get update && \
    apt-get install -y --no-install-recommends libvips ca-certificates && \
    rm -rf /var/lib/apt/lists/*

WORKDIR /app

# Copy binaries from builder
COPY --from=builder /app/server .
COPY --from=builder /app/redirect .

# Create non-root user
RUN useradd -r -s /bin/false appuser
USER appuser

EXPOSE 5000 5001

CMD ["./server"]
