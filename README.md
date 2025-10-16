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