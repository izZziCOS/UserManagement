.PHONY: test test-unit test-integration test-cover

# Run only unit tests (excludes integration tests)
test-unit:
	go test -v ./...

# Run only integration tests
test-integration:
	go test -v -tags=integration ./...

# Run all tests (both unit and integration)
test-all: test-unit test-integration

# Alias for test-unit (so 'make test' only runs unit tests)
test: test-unit

# Test coverage (unit tests only)
test-cover:
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out

# Integration test coverage
test-integration-cover:
	go test -coverprofile=coverage-integration.out -tags=integration ./...
	go tool cover -html=coverage-integration.out