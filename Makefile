VERSIOIN ?= dev
COMMIT_SHA ?= $(shell git rev-parse --verify HEAD)

##################################################
# Defaults
##################################################
all: clean vendor tidy build test

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
# Formats Go programs
##################################################
.PHONY: fmt
fmt: 
	gofmt -s -w -e $(shell find . -type f -name '*.go' -not -path "*/vendor/*")

##################################################
# Simplify code and reorders imports
##################################################
.PHONY: tidy
tidy: fmt 
	goimports -v -w -e $(shell find . -type f -name '*.go' -not -path "*/vendor/*")


##################################################
# Runs the golangci-lint checker
##################################################
.PHONY: lint
lint:
	golangci-lint run

##################################################
# Build the project and generate binary file
##################################################
.PHONY: build
build: 
	go build -v \
		-ldflags="\
		-X 'main.Version=${VERSIOIN}' \
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