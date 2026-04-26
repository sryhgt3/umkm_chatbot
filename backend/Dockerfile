# Stage 1: Build the Go binary
FROM golang:1.25-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the application
RUN go build -o server cmd/server/main.go

# Stage 2: Create the final image
FROM alpine:latest

# Set the working directory
WORKDIR /app

# Copy the binary from the builder stage
COPY --from=builder /app/server .
# Copy migrations from the builder stage
COPY --from=builder /app/migrations ./migrations

# Expose the port the app runs on
EXPOSE 8080

# Run the binary
ENTRYPOINT ["./server"]
