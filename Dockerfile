# Stage 1: Build the Go application
FROM golang:1.23.2 AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy the go.mod and go.sum files to download dependencies first
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the application, specifying the entry point in cmd/server/main.go
RUN go build -o vlink_backend ./cmd/server/main.go

# Stage 2: Run the application
FROM golang:1.23.2

# Set the working directory inside the container
WORKDIR /app

# Copy the compiled binary from the builder stage
COPY --from=builder /app/vlink_backend /app/vlink_backend

# Copy any necessary configuration files (e.g., .env) into the container
COPY .env .env

# Expose the application port (change if different)
EXPOSE 5000

# Set environment variables if necessary, e.g., ENV=development
ENV ENV=development

# Run the application
CMD ["/app/vlink_backend"]