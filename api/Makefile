test:
	go test -v ./...

test-cover:
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out

integration-test:
	go test -v ./tests/integration/... -tags=integration