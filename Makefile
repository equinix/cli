.PHONY: lint fix build docs docs-check onboard
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

# onboard - Scaffold a new service integration
# Usage: make onboard SERVICE=fabricv5
# This will create cmd/<service>.go with dynamic command registration
onboard:
	@if [ -z "$(SERVICE)" ]; then \
		echo "Error: SERVICE parameter is required"; \
		echo "Usage: make onboard SERVICE=fabricv5"; \
		exit 1; \
	fi
	@echo "Onboarding new service: $(SERVICE)"
	@echo "Creating cmd/$(SERVICE).go..."
	@sed -e 's/fabricv4/$(SERVICE)/g' \
	     -e 's/Fabric v4/$(shell echo $(SERVICE) | sed 's/\([a-z]\)\([A-Z]\)/\1 \2/g' | sed 's/v\([0-9]\)/v\1/g')/g' \
	     cmd/fabricv4.go > cmd/$(SERVICE).go
	@echo "Creating internal/api/$(SERVICE).go..."
	@sed -e 's/fabricv4/$(SERVICE)/g' \
	     -e 's/FabricV4/$(shell echo $(SERVICE) | sed 's/\(.\)/\u\1/' | sed 's/v\([0-9]\)/V\1/g')/g' \
	     -e 's/Fabric v4/$(shell echo $(SERVICE) | sed 's/\([a-z]\)\([A-Z]\)/\1 \2/g' | sed 's/v\([0-9]\)/v\1/g')/g' \
	     internal/api/fabricv4.go > internal/api/$(SERVICE).go
	@echo ""
	@echo "Service $(SERVICE) scaffolded successfully!"
	@echo ""
	@echo "Next steps:"
	@echo "1. Review and adjust cmd/$(SERVICE).go"
	@echo "2. Review and adjust internal/api/$(SERVICE).go"
	@echo "3. Ensure the SDK package exists: github.com/equinix/equinix-sdk-go/services/$(SERVICE)"
	@echo "4. Run 'make build' to verify the integration"
	@echo "5. Update documentation in README.md"