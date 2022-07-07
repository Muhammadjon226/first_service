.PHONY: build
PKG_LIST := $(shell go list ./... | grep -v /vendor/)
CURRENT_DIR=$(shell pwd)
APP=template
APP_CMD_DIR=./cmd

ifneq (,$(wildcard ./.env))
	include .env
endif

build:
	CGO_ENABLED=0 GOOS=linux go build -mod=vendor -a -installsuffix cgo -o ${CURRENT_DIR}/bin/${APP} ${APP_CMD_DIR}/main.go

set-env:
	./scripts/set-env.sh ${CURRENT_DIR}

proto-gen:
	./scripts/gen-proto.sh	${CURRENT_DIR}

lint:
	golint -set_exit_status ${PKG_LIST}

migrate-jeyran: set-env
	env POSTGRES_HOST=${POSTGRES_HOST} env POSTGRES_PORT=${POSTGRES_PORT} env POSTGRES_USER=${POSTGRES_USER} env POSTGRES_PASSWORD=${POSTGRES_PASSWORD} env POSTGRES_DB=${POSTGRES_DB} ./scripts/migrate-jeyran.sh

pull-proto-module:
	git submodule update --init --recursive

update-proto-module:
	git submodule update --remote --rebase