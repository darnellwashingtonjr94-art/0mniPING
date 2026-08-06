.PHONY: build test clean lint docker bench

BINARY_NAME=omniping
VERSION?=v1.0.0
BUILD_TIME=$(shell date -u '+%Y-%m-%d_%H:%M:%S')

build:
	@echo "Building binary..."
	mkdir -p bin/
	go build -ldflags "-X main.Version=$(VERSION) -X main.BuildTime=$(BUILD_TIME)" -o bin/$(BINARY_NAME) cmd/omniping/*.go

test:
	@echo "Running tests..."
	go test -v -cover ./...

clean:
	@echo "Cleaning artifacts..."
	rm -rf bin/ coverage.out *.log data/

lint:
	@echo "Linting code..."
	golangci-lint run ./... || go vet ./...

docker:
	@echo "Building Docker image..."
	docker build -t $(BINARY_NAME):$(VERSION) .

bench:
	@echo "Running benchmarks..."
	go test -bench=. -benchmem ./tests/benchmarks/...
