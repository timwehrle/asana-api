
# Commands
GOCMD := go
GOMOD := $(GOCMD) mod
GOVET := $(GOCMD) vet
GOTEST := $(GOCMD) test
GOBUILD := $(GOCMD) build
GORUN := $(GOCMD) run
GOCLEAN := $(GOCMD) clean
LINT := golangci-lint run
GOFMT := gofmt
VULN := golang.org/x/vuln/cmd/govulncheck@latest


.PHONY: lint
lint: ## Run linter
	@echo "Running lint..."
	$(LINT)