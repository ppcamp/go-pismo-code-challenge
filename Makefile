default: help

.PHONY: run
.PHONY: help
.PHONY: build
.PHONY: docker
.PHONY: migrate
.PHONY: create_migration
.PHONY: revert_migrations
.PHONY: generate
.PHONY: setup_dev

# Inner vars
SHELL := /bin/bash
ESC = \x1b

REGEX_COLUMN_SEP = :
REGEX_MAKEFILE_DOC = ^([a-zA-Z_-]+):.* \#\# (.*)$$
REGEX_MD_ITALIC = \*{1}([a-zA-Z0-9=\_\ \-]+)\*{1}
REGEX_MD_BOLD = \*{2}([a-zA-Z0-9=\_\ \-]+)\*{2}
REGEX_MD_MONO = `([a-zA-Z0-9=\_\ \-]+)`
REGEX_MD_LINK = \[([a-zA-Z0-9=\_\ \-]+)\]
REGEX_MD_UNDERLINE = \_{1}([a-zA-Z0-9=\_\ \-]+)\_{1}

# fonts
LD = ${ESC}[0m#   default
LB = ${ESC}[1m#   bold
FF = ${ESC}[2m#   faint
FI = ${ESC}[3m#   italic
FU = ${ESC}[4m#   underline
# foreground colors
F0 = ${ESC}[30m#  black
F1 = ${ESC}[31m#  red
F2 = ${ESC}[32m#  green
F3 = ${ESC}[33m#  yellow
F4 = ${ESC}[34m#  blue
F5 = ${ESC}[35m#  magenta
F6 = ${ESC}[36m#  cyan
F7 = ${ESC}[37m#  light gray
F8 = ${ESC}[90m#  gray
F9 = ${ESC}[91m#  light red
F10 = ${ESC}[92m# light green
F11 = ${ESC}[93m# light yellow
F12 = ${ESC}[93m# light blue
F13 = ${ESC}[94m# light blue
F14 = ${ESC}[95m# light magenta
F15 = ${ESC}[96m# light cyan
F16 = ${ESC}[97m# white
# background colors
B0 = ${ESC}[40m#   black
B1 = ${ESC}[41m#   red
B2 = ${ESC}[42m#   green
B3 = ${ESC}[43m#   yellow
B4 = ${ESC}[44m#   blue
B5 = ${ESC}[45m#   magenta
B6 = ${ESC}[46m#   cyan
B7 = ${ESC}[47m#   light gray
B8 = ${ESC}[100m#  gray
B9 = ${ESC}[101m#  light red
B10 = ${ESC}[102m# light green
B11 = ${ESC}[103m# light yellow
B12 = ${ESC}[103m# light blue
B13 = ${ESC}[104m# light blue
B14 = ${ESC}[105m# light magenta
B15 = ${ESC}[106m# light cyan
B16 = ${ESC}[107m# white


ifeq ($(shell test -f .env && echo -n EXIST_ENV), EXIST_ENV)
    include .env
    export
endif


# Migrations
N?=1

# Docker
TAG=latest
BRANCH_TAG=`git describe --abbrev=0 --tags | sort | head -n1`
ifeq ($(strip $(BRANCH_TAG)),)
    TAG="${BRANCH_TAG}"
endif
IMAGE=ppcamp/go-microservice-authentication:${TAG}


up: ## Run docker compose
	docker compose up -d

down: ## Remove docker compose containers
	docker compose down

run: ## **Run** the server
	go run cmd/main.go


build: ## Build the server locally
	go build -race cmd/main.go


lint: ## Run *linters* to this project. Remember to run `make setup_dev`
	@echo "Running linters"
	golangci-lint run ./...


docker: ## Create [docker] image
	@echo "Building ${IMAGE}"
	@docker build --no-cache -f Dockerfile -t ${IMAGE} .


migrate: ## Run migrations created with `make create_migration`. Remember to `make setup_dev`
	@echo "Running migrations"
	@migrate -path migrations -database "postgresql://${POSTGRES_USER}:${POSTGRES_PASSWORD}@localhost:5432/${POSTGRES_DATABASE}?sslmode=disable&application_name=authmigration" -verbose up


create_migration: ## Create a new migration, e.g `name=teste make create_migration`. Remember to `make setup_dev`
	@echo "Creating migration"
	@migrate create -ext sql -dir migrations -seq ${name}


revert_migrations: ## Revert a given migration, e.g `N=2 make revert_migrations`, by default 1. Remember to `make setup_dev`
	@echo "Reverting migrations"
	@migrate -path migrations -database "postgresql://${POSTGRES_USER}:${POSTGRES_PASSWORD}@localhost:5432/${POSTGRES_DATABASE}?sslmode=disable&application_name=authmigration" -verbose down ${N}


generate: ## Create [golang] protos
	@echo "Generating protos"
	@echo " - Generating messages"
	@protoc --go_out=. protos/*.proto
	@echo " - Generating services"
	@protoc --go-grpc_out=. protos/*.proto


setup_dev: ## Install _dev_ dependencies
	@echo "Installing go-migrate"
	@go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
	@echo "Installing linters"
	@curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.64.7



help:
	@printf "$(FF) Available methods: $(LD)\n\n"
    # 1. read makefile
    # 2. get lines that can have a method description
    # 3. color method names and add a COLUMN_SEPARATOR
    # 4. colour backticks (``)
    # 5. colour brackets ([])
    # 6. make it bold
    # 7. make it italic
    # 8. make it underline
    # 9. show as table
	@cat $(MAKEFILE_LIST) | \
	 	grep -E "$(REGEX_MAKEFILE_DOC)" | \
		sed -En 's/$(REGEX_MAKEFILE_DOC)/$(F2)\1$(REGEX_COLUMN_SEP)$(LD)\2/p' | \
		sed -En 's/$(REGEX_MD_MONO)/$(F3)\1$(LD)/g;p' | \
		sed -En 's/$(REGEX_MD_LINK)/$(F6)\1$(LD)/g;p' | \
		sed -En 's/$(REGEX_MD_BOLD)/$(LB)\1$(LD)/g;p' | \
		sed -En 's/$(REGEX_MD_ITALIC)/$(FI)\1$(LD)/g;p' | \
		sed -En 's/$(REGEX_MD_UNDERLINE)/$(FU)\1$(LD)/g;p' | \
		column -t -s "$(REGEX_COLUMN_SEP)"
