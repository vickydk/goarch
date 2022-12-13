#!/bin/bash

export REPO_NAME=goarch

build:
	@echo "${NOW} Building HTTP Server"
	@go build -o ./bin/${REPO_NAME}-http cmd/http/main.go

build-grpc:
	@echo "${NOW} Building GRPC Server"
	@go build -o ./bin/${REPO_NAME}-grpc cmd/grpc/main.go

build-app:
	@echo "${NOW} Building GRPC & HTTP Server"
	@go build -o ./bin/${REPO_NAME}-grpc cmd/app/main.go

docker-build:
	@ echo "Building goarch image"
	@ docker build -f Dockerfile -t goarch .

run: build
	@./bin/${REPO_NAME}-http

run-grpc: build-grpc
	@./bin/${REPO_NAME}-grpc

run-app: build-app
	@./bin/${REPO_NAME}-app