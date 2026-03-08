#!/bin/sh
set -e

echo "Stopping existing container..."
docker stop GoBackend 2>/dev/null || true
docker rm GoBackend 2>/dev/null || true

echo "Pulling latest code..."
git pull

echo "Building new image..."
docker build . -t go-backend

echo "Starting container..."
docker run \
    --name GoBackend \
    --env-file .env \
    --network host \
    -p 5000:5000 \
    --restart unless-stopped \
    -d go-backend

echo "Waiting for health check..."
sleep 5

if curl -sf http://localhost:5000/health > /dev/null 2>&1; then
    echo "Deployment successful - service is healthy"
else
    echo "WARNING: Health check failed. Check logs with: docker logs GoBackend"
    exit 1
fi
