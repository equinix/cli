# Development

## Dynamic Command Registration

The CLI uses reflection to automatically register commands from the Equinix SDK. This approach:

- **Automatically discovers** all API services in the SDK client
- **Generates commands** for each service and method at build time
- **Reduces maintenance** - new SDK services are automatically available
- **Ensures consistency** - command structure mirrors the SDK structure

## Adding one New Service

To onboard a single new Equinix service (e.g., fabricv5), use the `onboard` target:

```sh
make onboard SERVICE=fabricv5
```

This will scaffold:
- `cmd/<service>.go` - Command registration
- `internal/api/<service>.go` - API client setup
- `cmd/descriptions/<service>.json` - Field descriptions to embed for help

After scaffolding, you'll need to:
1. Review and adjust the generated files
2. Ensure the SDK package exists in `github.com/equinix/equinix-sdk-go/services/<service>`
3. Run `make build` to verify the integration

## Adding all services defined in equinix-sdk-go

To onboard/update every service defined in equinix-sdk-go, use the `generate-all` target:

```sh
make generate-all
```

After generating, you'll need to:
1. Review and adjust the generated files
2. Run `make build` to verify the integration

## Building

```sh
make build
```

## Linting

```sh
make lint
```


## Fixing linter issues

```sh
make fix
```

**NOTE:** Some linter issues may require hands-on work to address.  Where possible, the make task abov will automatically address linter issues in generated code.  It's a good idea to carefully review lint issues first to see if any of them can be fixed proactively with changes to template files, etc.. 

## Generating Documentation

```sh
make docs
```
