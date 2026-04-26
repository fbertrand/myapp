# Makefile
APP=myapp
IMAGE=ghcr.io/tonuser/$(APP)
VERSION=$(shell git describe --tags --always)

.PHONY: build test lint docker-build run clean

build:
	go build -o bin/$(APP) ./cmd/main.go

test:
	go test -race -coverprofile=coverage.out ./...
	go tool cover -func=coverage.out

lint:
	go vet ./...
	staticcheck ./...   # go install honnef.co/go/tools/cmd/staticcheck@latest

docker-build:
	docker build -t $(IMAGE):$(VERSION) .

run:
	docker compose up --build

clean:
	rm -rf bin/ coverage.out