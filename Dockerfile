# --- Stage 1: Build the Go binary ---
FROM golang:1.22 AS builder
LABEL authors="12345debdut"

# Set destination for COPY commands
WORKDIR /app

# Configure module proxy with direct fallback to better handle restricted networks
ENV GOPROXY=https://proxy.golang.org,direct

# Copy go.mod and go.sum first for caching
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the entire app source
COPY . .

# Build the Go binary
RUN CGO_ENABLED=0 GOOS=linux go build -o myapp .

# --- Stage 2: Run the Go binary ---
FROM debian:bookworm-slim

WORKDIR /app

# Ensure CA certificates present in runtime image
RUN apt-get update && apt-get install -y --no-install-recommends ca-certificates && rm -rf /var/lib/apt/lists/*

# Copy binary from builder
COPY --from=builder /app/myapp .

# Expose port
EXPOSE 8080

# Set environment variables
ENV PORT=8080

# Run app
CMD ["./myapp"]