.PHONY: lint fix build docs docs-check onboard update
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

# update - Update an existing service integration by fetching latest SDK and regenerating descriptions
# Usage: make update SERVICE=fabricv4
# This will:
# 1. Update the SDK package to the latest version
# 2. Extract SDK descriptions and save to cmd/descriptions/<service>.json
update:
	@if [ -z "$(SERVICE)" ]; then \
		echo "Error: SERVICE parameter is required"; \
		echo "Usage: make update SERVICE=fabricv4"; \
		exit 1; \
	fi
	@echo "Updating service: $(SERVICE)"
	@echo "Step 1: Updating SDK package..."
	@go get -u github.com/equinix/equinix-sdk-go/services/$(SERVICE)
	@go mod tidy
	@echo ""
	@echo "Step 2: Extracting SDK descriptions..."
	@mkdir -p cmd/descriptions
	@SDK_PATH=$$(go list -f '{{.Dir}}' -m github.com/equinix/equinix-sdk-go)/services/$(SERVICE); \
	if [ -z "$$SDK_PATH" ] || [ ! -d "$$SDK_PATH" ]; then \
		echo "Error: Could not find SDK path for $(SERVICE)"; \
		echo "Make sure github.com/equinix/equinix-sdk-go/services/$(SERVICE) is a valid module"; \
		exit 1; \
	fi; \
	echo "Extracting descriptions from: $$SDK_PATH"; \
	go run cmd/extract-descriptions/main.go --sdk-path "$$SDK_PATH" --output cmd/descriptions/$(SERVICE).json
	@echo ""
	@echo "Service $(SERVICE) updated successfully!"
	@echo "Description file: cmd/descriptions/$(SERVICE).json"
	@echo ""
	@echo "Next steps:"
	@echo "1. Review the updated SDK integration"
	@echo "2. Run 'make build' to verify the changes"
	@echo "3. Run 'make docs' to update documentation"
	@echo "4. Commit the changes including go.mod, go.sum, and cmd/descriptions/$(SERVICE).json"

# onboard - Scaffold a new service integration
# Usage: make onboard SERVICE=fabricv5
# This will create cmd/<service>.go, internal/<service>/<service>.go, and extract SDK descriptions
onboard:
	@if [ -z "$(SERVICE)" ]; then \
		echo "Error: SERVICE parameter is required"; \
		echo "Usage: make onboard SERVICE=fabricv5"; \
		exit 1; \
	fi
	@echo "Onboarding new service: $(SERVICE)"
	@echo ""
	@echo "Step 1: Creating service scaffolding..."
	@mkdir -p cmd
	@sed -e 's/{{SERVICE}}/$(SERVICE)/g' \
	     -e 's/{{SERVICE_DISPLAY}}/$(shell echo $(SERVICE) | sed 's/\([a-z]\)\([A-Z]\)/\1 \2/g' | sed 's/v\([0-9]\)/v\1/g' | sed 's/\b\(.\)/\u\1/g')/g' \
	     -e 's/{{SERVICE_ALIAS}}/$(shell echo $(SERVICE) | sed 's/v[0-9]*$$//')/g' \
	     templates/cmd/service.go.tmpl > cmd/$(SERVICE).go
	@echo "  - Created cmd/$(SERVICE).go"
	@mkdir -p internal/$(SERVICE)
	@sed -e 's/{{SERVICE}}/$(SERVICE)/g' \
	     -e 's/{{SERVICE_DISPLAY}}/$(shell echo $(SERVICE) | sed 's/\([a-z]\)\([A-Z]\)/\1 \2/g' | sed 's/v\([0-9]\)/v\1/g' | sed 's/\b\(.\)/\u\1/g')/g' \
	     templates/internal/service.go.tmpl > internal/$(SERVICE)/$(SERVICE).go
	@echo "  - Created internal/$(SERVICE)/$(SERVICE).go"
	@echo ""
	@echo "Step 2: Fetching SDK and extracting descriptions..."
	@$(MAKE) update SERVICE=$(SERVICE)
	@echo ""
	@echo "Service $(SERVICE) onboarded successfully!"
	@echo ""
	@echo "Files created:"
	@echo "  - cmd/$(SERVICE).go"
	@echo "  - internal/$(SERVICE)/$(SERVICE).go"
	@echo "  - cmd/descriptions/$(SERVICE).json"
	@echo ""
	@echo "Next steps:"
	@echo "1. Review and adjust the generated files as needed"
	@echo "2. Add service-specific aliases in cmd/$(SERVICE).go if desired"
	@echo "3. Run 'make build' to verify the integration"
	@echo "4. Run 'make docs' to generate documentation"
	@echo "5. Update README.md with information about the new service"