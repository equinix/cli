## equinix init

Configure API credentials

### Synopsis

Interactively configure the client ID and secret the CLI uses to
authenticate with the Equinix API, and save them to the config file.

The credentials are checked by requesting an OAuth2 token before they are
saved. Leaving an answer blank, or pressing Ctrl-C, leaves the config file
unchanged.

```
equinix init [flags]
```

### Options

```
      --api-endpoint string   Equinix API base URL to save to the config file (env EQUINIX_API_ENDPOINT, default https://api.equinix.com)
  -h, --help[=true|false]     help for init
      --sts-endpoint string   Equinix STS base URL to save to the config file (env EQUINIX_STS_ENDPOINT, default https://sts.eqix.equinix.com)
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.config/equinix/equinix.yaml)
```

### SEE ALSO

* [equinix](equinix.md)	 - Equinix CLI

