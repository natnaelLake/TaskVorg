# Use a specific Go version that matches your requirements
FROM golang:latest AS builder

# Set the working directory in the container
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the rest of the application source code
COPY . .

# Build the Go application
RUN go build -o main .

# Use a minimal base image for the final stage
FROM gcr.io/distroless/base:latest

# Set the working directory in the container
WORKDIR /app

# Copy the compiled binary from the builder stage
COPY --from=builder /app/main .

# Command to run the executable
ENTRYPOINT ["/app/main"]
