# Riseplus-go-boilerplate

<p align="center">
  <img src="https://go.dev/blog/go-brand/Go-Logo/SVG/Go-Logo_Blue.svg" alt="Go Logo" width="500" height="300">
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

4. Run the application:

   ```
   make run
   ```

5. (Optional) Build and run with Docker:
   ```
   make docker-build
   make docker-run
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
│   ├── repository/        # Data access layer
│   ├── service/           # Business logic layer
│   └── utils/             # Utility functions and helpers
├── pkg/                   # Library code that's ok to use by external applications
├── scripts/               # Scripts for development, CI/CD, etc.
├── .gitignore             # Specifies intentionally untracked files to ignore
├── Dockerfile             # Instructions for building a Docker container
├── go.mod                 # Go module definition and dependency management
├── go.sum                 # Checksums of the dependencies
└── README.md              # Project documentation and information
```
