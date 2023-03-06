#!/usr/bin/make
# Makefile readme (ru): <http://linux.yaroslavl.ru/docs/prog/gnu_make_3-79_russian_manual.html>
# Makefile readme (en): <https://www.gnu.org/software/make/manual/html_node/index.html#SEC_Contents>
include .env
export

docker_bin := $(shell command -v docker 2> /dev/null)
docker_compose_bin := $(shell command -v docker-compose 2> /dev/null)

.DEFAULT_GOAL := help

# This will output the help for each task. thanks to https://marmelab.com/blog/2016/02/29/auto-documented-makefile.html
help: ## Show this help
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)
	@echo "\n  Allowed for overriding next properties:\n\n\
	    PULL_TAG - Tag for pulling images before building own\n\
	              ('latest' by default)\n\
	    PUBLISH_TAGS - Tags list for building and pushing into remote registry\n\
	                   (delimiter - single space, 'latest' by default)\n\n\
	  Usage example:\n\
	    make PULL_TAG='v1.2.3' PUBLISH_TAGS='latest v1.2.3 test-tag' app-push"

build:
	$(docker_compose_bin) -f docker-compose.yml -f docker-compose.override.yml build --no-cache --parallel --force-rm
	$(docker_compose_bin) up --remove-orphans --force-recreate

up:
	$(docker_compose_bin) -f docker-compose.yml -f docker-compose.override.yml up --no-recreate -d

debug:
	$(docker_compose_bin) up --no-recreate

prod:
	$(docker_bin) volume prune --force
	DOCKER_BUILDKIT=0 $(docker_compose_bin) -f docker-compose.yml up --build --remove-orphans --force-recreate

dev:
	$(docker_bin) volume prune --force
	DOCKER_BUILDKIT=0 $(docker_compose_bin) -f docker-compose.yml -f docker-compose.override.yml up --build --remove-orphans --force-recreate

down:
	$(docker_compose_bin) down

swagger_gen:
	swag init --dir cmd,internal --generalInfo api/main.go

migrate-up:
	migrate -source="file://./database/migrations/" -database="$(DB_DRIVER)://$(DB_USERNAME):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_DATABASE)?sslmode=disable" up

migrate-down:
	migrate -source="file://./database/migrations/" -database="$(DB_DRIVER)://$(DB_USERNAME):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_DATABASE)?sslmode=disable" down -all

mockgen:
	mockgen -source=./internal/core/ports/user.go -package mocks > ./internal/mocks/user.go

air:
	$(go env GOPATH)/bin/air

local:
	PGX_DATABASE=postgres://$(DB_USERNAME):$(DB_PASSWORD)@0.0.0.0:5432/$(DB_DATABASE) ~/go/bin/air

#tests:
#	$(docker_compose_bin) exec -it envs_api go tests ./internal/core/services -v -race && go tests ./tests -v -race

tests:
	go test ./internal/core/services -v -race
	go test ./tests -v -race

.PHONY: tests
