.POSIX:

.PHONY: help
help: ## Show this help
	@egrep -h '\s##\s' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

.PHONY: init
init: ## Get dependencies
	go get -v -t -d ./...

.PHONY: fmt
fmt: ## Run all formatings
	go mod vendor
	go mod tidy
	go fmt ./...

.PHONY: run
run: ## Run all examples
	go run ./examples/simple/main.go
	@echo "\nVisit http://localhost:8080"
	@go run ./examples/http/main.go

.PHONY: test
test: ## Run all test
	go test -v ./...

.PHONY: docs
docs: ## Start local docs server
	@echo "See Documentation:"
	@echo "\thttp://localhost:6060/pkg/github.com/erdaltsksn/jerr"
	@echo "\n"
	@godoc -http=:6060

.PHONY: build
build: ## Build the app
	go build -v ./...

.PHONY: clean
clean: ## Clean all generated files
	rm -rf ./vendor/
