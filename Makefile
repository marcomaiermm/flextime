build:
	@echo "Building..."
	@go build -o flextime ./cmd/cli/main.go

run:
	@echo "Running..."
	@go run ./cmd/cli/main.go
