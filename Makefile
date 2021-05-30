VERSION ?= dev
COMMIT_SHA ?= $(shell git rev-parse --verify HEAD)
SRC_CODE ?= $(shell find . -type f -name '*.go' -not -path "*/vendor/*")

##################################################
# Defaults
##################################################
all: clean vendor lint build test

##################################################
# Runs go clean and removes generated binaries and coverfiles
##################################################
.PHONY: clean
clean:
	go clean ./...
	rm -rf ./gen
	rm -rf ./bin

##################################################
# Cleans up go mod dependencies and vendor's all dependencies
##################################################
.PHONY: vendor
vendor:
	go mod tidy
	go mod vendor

##################################################
# Examines Go source code and reports suspicious constructs
##################################################
.PHONY: vet
vet: 
	go vet ./...

##################################################
# Formats Go programs
##################################################
.PHONY: format
format: 
	gofmt -s -w -e ${SRC_CODE}

##################################################
# Updates import lines, adding missing ones and removing unreferenced ones.
##################################################
.PHONY: import
import:
	goimports -v -w -e ${SRC_CODE}

##################################################
# Runs the golangci-lint checker
##################################################
.PHONY: lint
lint: vet format import
	golangci-lint run

##################################################
# Build the project and generate binary file
##################################################
.PHONY: build
build: 
	go build -v \
		-ldflags="\
		-X 'main.Version=${VERSION}' \
		-X 'main.Commit=${COMMIT_SHA}'" \
		-o ./bin/line-coverage \
		.

##################################################
# Runs unit tests and generates a coverage file at coverage.out
##################################################
COVERFILE?=./gen/coverage.out
COVER_HTML?=./gen/coverage.html

.PHONY: test
test:
	$(call print-target)
	@mkdir -p gen  ## Creating a gen folder if it doesn't exist
	go test `go list ./... | grep -vE "/test/"` -race -covermode=atomic -coverprofile=$(COVERFILE)
	go tool cover -func=$(COVERFILE) 
	go tool cover -html=$(COVERFILE) -o $(COVER_HTML)


.PHONY: help
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

define print-target
    @printf "Executing target: \033[36m$@\033[0m\n"
endef
