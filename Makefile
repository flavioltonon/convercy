ENVIRONMENT ?= development

build:
	go build -o ./build/bin/convercy ./cmd/main.go

check:
	go vet ./...	

image:
	docker build -f Dockerfile -t flavioltonon/convercy:latest .

push:
	docker push flavioltonon/convercy:latest

start: stop
	docker-compose up --build

stop:
	docker-compose down --remove-orphans

tests:
	go test -cover ./...

tidy:
	go fmt ./...