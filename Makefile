# Run the application
run:
	go run main.go

# Build the application
build:
	go build -o bin/app main.go

# Install dependencies
install:
	go mod tidy

# Format code
fmt:
	go fmt ./...

# Run all checks before committing
check: fmt lint test
