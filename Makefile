ENVIRONMENT ?= development

# build builds the application. It is intended to be executed only inside a Docker container context
build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o ./build/bin/convercy ./cmd/main.go

check:
	go vet ./...	

image:
	docker build --rm -f Dockerfile -t flavioltonon/convercy:latest .

push:
	docker push flavioltonon/convercy:latest

release: image push

start: stop
	docker-compose up --build

stop:
	docker-compose down --remove-orphans

tests:
	go test -cover ./...

tidy:
	go fmt ./...