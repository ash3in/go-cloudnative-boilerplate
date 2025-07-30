# Variables
SERVICE_NAME ?= app
IMAGE_NAME ?= $(SERVICE_NAME):local

# Targets
.PHONY: help build run clean

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "  \033[36m%-14s\033[0m %s\n", $$1, $$2}'

dep:
	go mod tidy
	go mod vendor

build:
	docker build --tag $(IMAGE_NAME) .

run: build
	docker run --rm -p 8080:8080 $(IMAGE_NAME)

clean:
	docker rmi $(IMAGE_NAME) || true
