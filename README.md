# Equinix CLI

`equinix` is the command line interface (CLI) for interacting with Equinix resources and services.

## Installation

### Install from GitHub Releases

[Equinix CLI releases](https://github.com/equinix/cli/releases/latest) contain precompiled binaries for a variety of operating systems and architectures.

1. Download the appropriate zip archive for your platform from the desired Equinix CLI release
2. Run `unzip <zip file name>` to expand the zip archive you downloaded
3. You can run the command immediately as `./equinix` or, if you prefer, `cp equinix <directory>` where `<directory>` is a location on your local filesystem that is already included in your `PATH`.

### Install with Homebrew

If you prefer, you can also install `equinix` with Homebrew by running the following commands:

```sh
brew tap equinix/homebrew-tap
brew install equinix
```

## Usage

The full CLI documentation can be found [in the docs directory](docs/equinix.md).

### Dynamic Command Registration

The CLI uses reflection to automatically register commands from the Equinix SDK. This approach:

- **Automatically discovers** all API services in the SDK client
- **Generates commands** for each service and method at build time
- **Reduces maintenance** - new SDK services are automatically available
- **Ensures consistency** - command structure mirrors the SDK structure

## Development

### Adding New Services

To onboard a new Equinix service (e.g., fabricv5), use the `onboard` target:

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

### Building

```sh
make build
```

### Linting

```sh
make lint
```

### Generating Documentation

```sh
make docs
```
