# Build stage
FROM golang:1.24-alpine AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Argument to determine which service to build
ARG SERVICE_NAME

# Build the Go app
RUN go build -o main ./cmd/${SERVICE_NAME}

# Final stage
FROM alpine:latest

WORKDIR /app

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/main .

# Expose port (this is just documentation, docker-compose will handle actual mapping)
EXPOSE 8080

# Command to run the executable
CMD ["./main"]
