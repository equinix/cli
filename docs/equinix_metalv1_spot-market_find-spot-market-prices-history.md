## equinix metalv1 spot-market find-spot-market-prices-history

Get spot market prices for a given period of time

### Synopsis

# Get spot market prices for a given plan and facility in a fixed period of time *Note: In the `200` response, the property `datapoints` contains arrays of `[float, integer]`.*

Use --request flag to provide optional JSON payload fields.

```
equinix metalv1 spot-market find-spot-market-prices-history [flags]
```

### Options

```
      --facility string   facility field
      --from string       from field
  -h, --help              help for find-spot-market-prices-history
      --metro string      metro field
      --plan string       plan field
      --request string    JSON payload for additional optional fields not exposed as flags
      --until string      until field
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug           Enable debug logging for HTTP requests
  -f, --format string   Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix metalv1 spot-market](equinix_metalv1_spot-market.md)	 - Manage spot-market resources

