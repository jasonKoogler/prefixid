.PHONY: test test-verbose test-cover test-race fmt lint clean

# Default target
all: fmt lint test

# Run all tests
test:
	go test ./tests/...

# Run tests with verbose output
test-verbose:
	go test -v ./tests/...

# Run tests with race detector
test-race:
	go test -race ./tests/...

# Run tests with coverage
test-cover:
	go test -cover ./tests/... -coverpkg ./...

# Generate coverage report and open in browser
cover:
	go test -coverprofile=coverage.out -coverpkg ./... ./tests/...
	go tool cover -html=coverage.out -o coverage.html

# Format code
fmt:
	go fmt ./...

# Lint code
lint:
	go vet ./...

# Clean up
clean:
	rm -f coverage.out


# Run all tests and generate a coverage report
test-all: fmt lint test-race test-cover cover 