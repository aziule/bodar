# Commands
GO-LINT = docker run --env=GOFLAGS=-mod=vendor --rm -v $(CURDIR):/app -w /app golangci/golangci-lint:v1.18.0 golangci-lint

default: help
.PHONY: default

## lint: Lint the Go code
lint:
	$(GO-LINT) run -E golint,gofmt,unparam,goconst,prealloc --exclude-use-default=false --deadline=5m
.PHONY: lint

## help: Show this help
help: Makefile
	@echo
	@echo "Available targets:"
	@echo
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /' | LANG=C sort
	@echo
.PHONY: help