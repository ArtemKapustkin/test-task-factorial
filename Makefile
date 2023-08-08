default: help

# https://dwmkerr.com/makefile-help-command/
.PHONY: help
help: # Show help for each of the Makefile recipes.
	@grep -E '^[a-zA-Z0-9 -]+:.*#'  Makefile | sort | while read -r l; do printf "\033[1;32m$$(echo $$l | cut -f 1 -d':')\033[00m:$$(echo $$l | cut -f 2- -d'#')\n"; done

.PHONY: run
run: # Start http server.
	go run cmd/main.go

.PHONY: test
test: # Run all tests.
	go test ./...

.PHONY: lint
lint: # Run linter.
	golangci-lint run ./...