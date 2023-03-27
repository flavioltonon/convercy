# build builds the application. It is intended to be executed only inside a Docker container context
build:
	@echo "Compiling the application..."
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o ./build/bin/convercy ./cmd/main.go

check:
	@echo "# Checking for suspicious, abnormal, or useless code..."
	@go vet ./...

image:
	@echo "# Building the application image..."
	@docker build --rm -f Dockerfile -t flavioltonon/convercy:latest .

install:
	@echo "# Installing dependencies..."
	@go mod tidy

push:
	@echo "# Pushing image to registry..."
	@docker push flavioltonon/convercy:latest

start: stop
	@echo "Initializing application and its dependencies..."
	@docker-compose up --build

stop:
	@echo "Stopping application and its dependencies..."
	@docker-compose down --remove-orphans

tests:
	@echo "# Running tests..."
	@go test -cover ./application/services/... ./domain/services/...

tidy:
	@echo "# Formatting code..."
	@go fmt ./...