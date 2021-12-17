#!/usr/bin/make

.PHONY: help
.SECONDEXPANSION:

################################################################################
# Set common variables
PROJECT                             := game-of-life
BUILD_OUTPUT_DIR                    ?= dist
SERVICE_PORT						:= 8080

################################################################################
help: ## Print this help message.
	@echo "Usable make targets:"
	@echo "$$(grep -hE '^\S+:.*##' $(MAKEFILE_LIST) | sed -e 's/:.*##\s*/:/' -e 's/^\(.\+\):\(.*\)/\1:\2/' | column -c2 -t -s : | sort)"

################################################################################
# Build, Package, Test and Code Quality Make Targets

download-deps:
	go mod download -x

build-local:
	cd ./cmd; go build -o ../dist/gameoflife

test-local:
	cd ./cmd; go test ./... -coverprofile ../dist/coverage.out

start-local:
	export SERVICE_PORT=$(SERVICE_PORT)
	./dist/gameoflife

build:
	docker build \
		--network=host \
		-f docker/build.Dockerfile \
		-t "$(PROJECT)" \
		.

start:
	docker run -p $(SERVICE_PORT):$(SERVICE_PORT) --env SERVICE_PORT=$(SERVICE_PORT) $(PROJECT)

gen-mocks:
	mockgen -destination=cmd/publish/publisher_mock.go -package="publish" -source=cmd/publish/publisher.go