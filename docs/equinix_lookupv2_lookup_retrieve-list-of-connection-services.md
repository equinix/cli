## equinix lookupv2 lookup retrieve-list-of-connection-services

Retrieve list of connection services

### Synopsis

This method retrieves a list of supported connection services such as media types, protocol types, connector types and circuit counts.

Use --request flag to provide optional JSON payload fields.

```
equinix lookupv2 lookup retrieve-list-of-connection-services [flags]
```

### Options

```
  -h, --help             help for retrieve-list-of-connection-services
      --ibx string       ibx field
      --request string   JSON payload for additional optional fields not exposed as flags
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug           Enable debug logging for HTTP requests
  -f, --format string   Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix lookupv2 lookup](equinix_lookupv2_lookup.md)	 - Manage lookup resources

