SHELL := /bin/bash -o pipefail

UNAME_OS := $(shell uname -s)
UNAME_ARCH := $(shell uname -m)

.PHONY: env
env:
	@echo "GOPATH: $(GOPATH)"

.PHONY: lint
lint:
	@echo "## Run GolangCI Lint"
	golangci-lint -E bodyclose,misspell,gocyclo,dupl,gofmt,golint,unconvert,goimports,depguard,gocritic,funlen run

.PHONY: build
build:
	@echo "## Build API"
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -a -installsuffix cgo -o build/api cmd/api/main.go
	@echo "## Build Migration"
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -a -installsuffix cgo -o build/migration cmd/migration/main.go
	@echo "## Build Worker"
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -a -installsuffix cgo -o build/worker cmd/worker/main.go

.PHONY: build_docker
build_docker:
	@echo "## Build Docker Image"
	docker build -t go-service-template:latest -f Dockerfile .

.PHONY: test
test:
	@echo "## Run Unit Tests"
	go test -covermode=atomic -coverprofile=coverage.out ./... -v
	go tool cover -html=coverage.out -o coverage.html

.PHONY: migrate
migrate:
	@echo "## Migrate DB"
	go run cmd/migration/main.go

.PHONY: api
api:
	@echo "## Start API"
	go run cmd/api/main.go

.PHONY: worker
worker:
	@echo "## Start Worker"
	go run cmd/worker/main.go
