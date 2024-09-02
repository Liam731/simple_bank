# Build stage
FROM golang:1.22-alpine3.19 AS builder

# Create a directory for the app
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the Go application
RUN go build -o main main.go

# Run stage
FROM alpine:3.19

# Create a directory for the app
WORKDIR /app

# Copy the binary from the builder stage
COPY --from=builder /app/main .

# Expose the port the app will run on
EXPOSE 8080

# Run the Go application
CMD ["/app/main"]
