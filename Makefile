.PHONY: lint fix build docs docs-check
BINARY=equinix

GOLANGCI_LINT_VERSION=v2.3.0
GOLANGCI_LINT=go run github.com/golangci/golangci-lint/v2/cmd/golangci-lint@${GOLANGCI_LINT_VERSION}

lint:
	$(GOLANGCI_LINT) run -v

fix:
	$(GOLANGCI_LINT) run -v --fix

build:
	go build $(LDFLAGS) -o bin/$(BINARY) main.go

docs:
	rm -rf docs/
	mkdir -p docs
	go run main.go docs ./docs

docs-check: docs
	if git status --porcelain | grep docs; then \
		echo "Uncommitted changes detected. Run 'make docs' and commit changes."; \
		exit 1; \
	fi