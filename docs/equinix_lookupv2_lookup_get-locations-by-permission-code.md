## equinix lookupv2 lookup get-locations-by-permission-code

Get Locations by permission code

### Synopsis

This method retrieves all the user’s IBX locations, accounts, cages and cabinets information according to their ordering permission. This includes A-side and Z-side location information for a cross connect order.

Use --request flag to provide optional JSON payload fields.

```
equinix lookupv2 lookup get-locations-by-permission-code [flags]
```

### Options

```
      --a-side-ibx string                a-side-ibx field
      --connection-service string        connection-service field
      --details                          details field
  -h, --help                             help for get-locations-by-permission-code
      --ibxs string                      ibxs field (JSON or string)
      --permission-code string           permission-code field
      --provider-account-number string   provider-account-number field
      --request string                   JSON payload for additional optional fields not exposed as flags
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug           Enable debug logging for HTTP requests
  -f, --format string   Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix lookupv2 lookup](equinix_lookupv2_lookup.md)	 - Manage lookup resources

