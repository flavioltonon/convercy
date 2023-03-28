check:
	@echo "# Checking for suspicious, abnormal, or useless code..."
	@go vet ./...

install:
	@echo "# Installing dependencies..."
	@go mod tidy

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