# Convercy

A currency conversion management API written in Go.

## Pre-requisites

- golang@v1.18 or newer
- docker-compose@v2.15.1

## Features

### Managing registered currencies

- Register currency
- Unregister currency

### Converting a currency

- Convert currency

## Running the project

### Unix-based Operational Systems

> make start

### Other Operational Systems

> docker-compose up --build

## Documentation

The API complete specification can be found at docs/Convercy.postman_collection.json

## Running tests

### Unix-based Operational Systems

> make tests

### Other Operational Systems

> go test -cover ./...