## equinix metalv1 spot-market find-spot-market-request-by-id

Retrieve a spot market request

### Synopsis

Returns a single spot market request

Use --request flag to provide optional JSON payload fields.

```
equinix metalv1 spot-market find-spot-market-request-by-id [flags]
```

### Options

```
  -h, --help             help for find-spot-market-request-by-id
      --id string        SpotMarketRequest UUID (required)
      --include string   include field (JSON or string)
      --request string   JSON payload for additional optional fields not exposed as flags
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug           Enable debug logging for HTTP requests
  -f, --format string   Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix metalv1 spot-market](equinix_metalv1_spot-market.md)	 - Manage spot-market resources

