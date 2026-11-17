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
brew trust --cask equinix/tap/equinix
brew install equinix
```

## Usage

Run `equinix init` to create or enter API credentials and save them to
`~/.config/equinix/equinix.yaml`. Credentials can also be supplied with the
`EQUINIX_CLIENT_ID` and `EQUINIX_CLIENT_SECRET` environment variables, and a
non-production API can be selected with `EQUINIX_API_ENDPOINT` (or
`equinix init --api-endpoint`).

The full CLI documentation can be found [in the docs directory](docs/equinix.md).
