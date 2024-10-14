# Build stage
FROM golang:1.23.2-alpine AS builder

# Set the working directory
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the application with optimizations
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o server ./cmd/api

# Final stage
FROM gcr.io/distroless/static-debian12

# Set the working directory
WORKDIR /

# Copy the binary from the builder stage
COPY --from=builder /app/server .

# Expose the port the app runs on
EXPOSE 8080

# Set environment variables for better performance
ENV GOMAXPROCS=1

# Command to run the executable
CMD ["/server"]
