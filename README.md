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

The `brew trust` step is required because `equinix` is published as a cask from a
non-official tap. Since Homebrew 6.0.0, formulae and casks from non-official taps
are not loaded until that tap — or the specific item — has been trusted. Skipping
this step fails with:

```text
Error: Refusing to load cask equinix/tap/equinix from untrusted tap equinix/tap.
Run `brew trust --cask equinix/tap/equinix` or `brew trust equinix/tap` to trust it.
```

`brew trust --cask equinix/tap/equinix` trusts only the `equinix` cask. To trust
every current and future formula and cask published from the tap, run
`brew trust equinix/tap` instead. See [Tap Trust](https://docs.brew.sh/Tap-Trust)
for details.

## Usage

The full CLI documentation can be found [in the docs directory](docs/equinix.md).
