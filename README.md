# Riseplus-go-boilerplate

[![Go Test](https://github.com/pattaradol9/riseplus-go-boilerplate/actions/workflows/go-test.yml/badge.svg)](https://github.com/pattaradol9/riseplus-go-boilerplate/actions/workflows/go-test.yml)

[![Go Report Card](https://goreportcard.com/badge/github.com/pattaradol9/riseplus-go-boilerplate)](https://goreportcard.com/report/github.com/pattaradol9/riseplus-go-boilerplate)

[![Release](https://img.shields.io/github/v/release/pattaradol9/riseplus-go-boilerplate.svg)](https://github.com/pattaradol9/riseplus-go-boilerplate/releases)

<p align="center">
  <img src="https://gofiber.io/assets/images/logo.svg" alt="Go Fiber Logo" width="400" height="200">
</p>

## Description

Riseplus-go-boilerplate is a robust and scalable boilerplate for Go projects, optimized for Go Fiber. It provides a solid foundation for building modern, efficient, and maintainable web applications using the high-performance Fiber web framework.

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
   cd riseplus-go-boilerplate
   ```

3. Install dependencies:

   ```
   make tidy
   ```

4. Run tests:

   ```
   make test
   ```

5. Run the application:

   ```
   make run
   ```

6. (Optional) Build and run with Docker:

   ```
   make docker-build
   make docker-run
   ```

7. (Optional) Run tests in Docker:
   ```
   make docker-test
   ```

## Quick Start

Before you begin, it's recommended to change the module name in `go.mod` to match your project:

1. Open `go.mod` file
2. Replace the module name:
   ```
   module github.com/yourusername/your-project-name
   ```
3. Run `go mod tidy` to update dependencies

This ensures that your project has a unique module name, which is important for proper Go package management.

## Project Structure

```
riseplus-go-boilerplate/
├── cmd/                   # Command-line applications and entry points
│   └── api/
│       └── main.go        # Main entry point for the API
├── internal/              # Private application and library code
│   ├── api/               # API handlers and routes
│   ├── config/            # Configuration management
│   ├── models/            # Data models and structures
│   ├── middleware/        # Middleware layer
│   ├── repositories/      # Data access layer
│   ├── services/          # Business logic layer
│   ├── server/            # Server layer
│   ├── handlers/          # Handler layer
│   └── utils/             # Utility functions and helpers
├── pkg/                   # Library code that's ok to use by external applications
├── .gitignore             # Specifies intentionally untracked files to ignore
├── Dockerfile             # Instructions for building a Docker container
├── go.mod                 # Go module definition and dependency management
├── go.sum                 # Checksums of the dependencies
└── README.md              # Project documentation and information
```

## Makefile

The Makefile is used to manage the project. It provides a set of commands to build, clean, run, test, and lint the project.

### Makefile Commands

- `make build`: Build the project
- `make clean`: Clean up build artifacts
- `make run`: Run the project
- `make test`: Run tests
- `make coverage`: Run tests with coverage
- `make deps`: Check and get dependencies
- `make lint`: Run linter
- `make vet`: Run vet
- `make fmt`: Run fmt
- `make tidy`: Run tidy
- `make upgrade`: Upgrade dependencies to latest versions
- `make help`: Show help

## Contact

If you have any questions or need further assistance, please contact me at [pattaradol.n@riseplus.tech](mailto:pattaradol.n@riseplus.tech).
