DOCKER_COMPOSE=docker compose

DOCKER_MAC_BIN := /Applications/Docker.app/Contents/Resources/bin/
PATH := $(PATH):$(DOCKER_MAC_BIN)

OPENAPI_SPEC=api/v1/openapi.yaml
SWAGGER_DIR=internal/infrastructure/server/rest/handler/swagger

.PHONY: start stop generate

build:
	@echo "Building ..."
	$(DOCKER_COMPOSE) build

start:
	@echo "Starting services ..."
	$(DOCKER_COMPOSE) up -d
	@echo "Starting application ..."

stop:
	@echo "Stopping services ..."
	$(DOCKER_COMPOSE) down

test:
	@echo "Run test ..."
	go test ./...

generate:
	@echo "Generating OpenAPI server code ..."
	cp $(OPENAPI_SPEC) $(SWAGGER_DIR)/openapi.yaml
	oapi-codegen -config api/v1/oapi-codegen.yaml $(OPENAPI_SPEC)
