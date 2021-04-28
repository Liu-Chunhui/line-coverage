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
	rm -rf ./dist

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
build: clean 
	go build -v \
		-ldflags="\
		-X main.Version=${VERSIOIN}" \
		-o ./dist/line-coverage \
		.

##################################################
# Runs unit tests and generates a coverage file at coverage.out
##################################################
COVERFILE?=./gen/coverage.out
COVER_HTML?=./gen/coverage.html

.PHONY: test
test:
	@mkdir -p gen  ## Creating a gen folder if it doesn't exist
	go test `go list ./... | grep -vE "/test/"` -race -covermode=atomic -coverprofile=$(COVERFILE)
	go tool cover -func=$(COVERFILE) 
	go tool cover -html=$(COVERFILE) -o $(COVER_HTML)
