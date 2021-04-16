NAME := avoxiinterview
SHELL := /bin/bash
FLAGS :=
CONTAINER := billdoss/$(NAME)
BUILD_DIR := dist
DATA_DIR := data

.PHONY: build test docker dist/avox-interview

build:
	CGO_ENABLED=0 GOOS=linux go build -o $(BUILD_DIR)/$(NAME)
	cp -R $(DATA_DIR) $(BUILD_DIR)

run:
	go run main.go

docker: dist/$(NAME)
	docker build --pull --no-cache -t $(CONTAINER) .

docker-run:
	docker run --rm -p 8080:10000 $(CONTAINER)
	
test:
	go test ./...

clean:
	go clean
	docker -rmi $(CONTAINER)

