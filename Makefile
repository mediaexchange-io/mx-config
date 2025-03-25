.PHONY: help
help: ## Display valid Makefile targets.
	@echo 'targets:'
	@echo
	@grep -E '^[a-zA-Z_-]+:\s*##\s*.*$$' $(MAKEFILE_LIST) | sed -H 's/:\s*##\s*/#/' | column -s '#' -t

.PHONY: clean
clean: ## Clean up build artifacts
	@rm -f coverage.out coverage.html

.PHONY: update
update: ## Updates go module dependencies
	@go clean -modcache
	@go get -u -t ./...
	@go mod tidy

.PHONY: test
test: ## Run unit tests
	@go test ./...

.PHONY: coverage
coverage: ## Run unit tests with coverage analysis
	@go test -covermode=atomic -coverprofile=coverage.out ./...
	@go tool cover -html=coverage.out -o coverage.html
	@echo 'open coverage.html for coverage heat map'
