# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOMOD=$(GOCMD) mod
BINARY_NAME=riseplus-go-boilerplate
BINARY_UNIX=$(BINARY_NAME)_unix
MAIN_PATH=cmd/api/main.go

# Make is verbose in Linux. Make it silent.
MAKEFLAGS += --silent

.PHONY: all build clean run test coverage deps lint vet fmt tidy help upgrade

all: test build ## Run tests and build the project

build: ## Build the project
	@echo "Building..."
	$(GOBUILD) -o $(BINARY_NAME) -v $(MAIN_PATH)

clean: ## Clean up build artifacts
	@echo "Cleaning..."
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
	rm -f $(BINARY_UNIX)

run: ## Run the project
	@echo "Running..."
	$(GOCMD) run $(MAIN_PATH)

test: ## Run tests
	@echo "Running tests..."
	$(GOTEST) -v ./...

coverage: ## Run tests with coverage
	@echo "Running tests with coverage..."
	$(GOTEST) -v -coverprofile=coverage.out ./...
	$(GOCMD) tool cover -html=coverage.out

deps: ## Check and get dependencies
	@echo "Checking dependencies..."
	$(GOGET) -v -t -d ./...

lint: ## Run linter
	@echo "Linting..."
	golangci-lint run

vet: ## Run go vet
	@echo "Vetting..."
	$(GOCMD) vet ./...

fmt: ## Format code
	@echo "Formatting..."
	$(GOCMD) fmt ./...

tidy: ## Tidy up modules
	@echo "Tidying up modules..."
	$(GOMOD) tidy

upgrade: ## Upgrade dependencies to latest versions
	@echo "Upgrading dependencies..."
	$(GOGET) -u ./...
	$(GOMOD) tidy

build-linux: ## Cross-compile for Linux
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD)

build-docker: ## Build Docker image
	@echo "Building Docker image..."
	docker build -t riseplus-go-boilerplate:latest . --no-cache --platform linux/amd64

run-docker: ## Run Docker container
	@echo "Running Docker container..."
	docker run -p 8080:8080 riseplus-go-boilerplate:latest

docker-push: ## Push Docker image to registry
	@echo "Pushing Docker image to registry..."
	docker push riseplus-go-boilerplate:latest

help: ## Display available commands
	@echo "Available commands:"
	@echo "Command               Description"
	@echo "-------               -----------"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'
