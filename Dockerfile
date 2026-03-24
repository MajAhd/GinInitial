# Build stage
FROM golang:alpine AS builder

WORKDIR /app

# Install build dependencies
RUN apk add --no-cache build-base

# Copy module files and install dependencies
COPY go.mod go.sum* ./
RUN go mod download

# Copy application source code
COPY . .

# Build the application statically compiled
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o server ./cmd/server

# Final stage
FROM alpine:latest

# Certificates for external API calls, and timezone data
RUN apk --no-cache add ca-certificates tzdata

WORKDIR /app

# Copy binary from builder
COPY --from=builder /app/server .
# Copy .env configuration if exists
COPY --from=builder /app/.env* ./

EXPOSE 8080

# Run the binary
CMD ["./server"]
