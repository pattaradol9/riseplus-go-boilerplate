# Riseplus-go-boilerplate

## Description

Riseplus-go-boilerplate is a robust and scalable boilerplate for Go projects, providing a solid foundation for building modern, efficient, and maintainable Go applications.

## Features

- Clean architecture structure
- RESTful API setup
- Docker support
- Logging and error handling
- Optimized for performance

## Prerequisites

- Go 1.23.2 or higher
- Docker (optional, for containerization)

## Getting Started

1. Clone the repository:

   ```
   git clone https://github.com/pattaradol9/riseplus-go-boilerplate.git
   ```

2. Navigate to the project directory:

   ```
   cd Riseplus-go-boilerplate
   ```

3. Install dependencies:

   ```
   go mod tidy
   ```

4. Run the application:

   ```
   go run cmd/api/main.go
   ```

5. (Optional) Build and run with Docker:
   ```
   docker build -t riseplus-go-boilerplate .
   docker run -p 8080:8080 riseplus-go-boilerplate
   ```

## Project Structure

Riseplus-go-boilerplate/
├── cmd/
│ └── api/
│ └── main.go
├── internal/
│ ├── api/
│ ├── config/
│ ├── models/
│ ├── repository/
│ ├── service/
│ └── utils/
├── pkg/
├── scripts/
├── .gitignore
├── Dockerfile
├── go.mod
├── go.sum
└── README.md
