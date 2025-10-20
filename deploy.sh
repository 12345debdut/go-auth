#!/bin/bash
set -e

APP_NAME="go_app"
IMAGE_NAME="myapp:latest"

echo "🚀 Building Go app with Docker..."

# Build Docker image
docker build -t $IMAGE_NAME .

echo "🧹 Stopping and removing old container (if any)..."
docker stop $APP_NAME || true
docker rm $APP_NAME || true

echo "🏗️ Starting new container..."
docker run -d \
  --name $APP_NAME \
  -p 8080:8080 \
  -e PORT=8080 \
  $IMAGE_NAME

echo "✅ Deployment complete. App running at http://localhost:8080"