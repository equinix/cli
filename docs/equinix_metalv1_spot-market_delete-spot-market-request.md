## equinix metalv1 spot-market delete-spot-market-request

Delete the spot market request

### Synopsis

Deletes the spot market request.

Use --request flag to provide optional JSON payload fields.

```
equinix metalv1 spot-market delete-spot-market-request [flags]
```

### Options

```
      --force-termination   force-termination field
  -h, --help                help for delete-spot-market-request
      --id string           SpotMarketRequest UUID (required)
      --request string      JSON payload for additional optional fields not exposed as flags
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug           Enable debug logging for HTTP requests
  -f, --format string   Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix metalv1 spot-market](equinix_metalv1_spot-market.md)	 - Manage spot-market resources

