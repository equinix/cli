## equinix fabricv4 routing-protocols get-connection-routing-protocols-changes

Get Changes

### Synopsis

This API provides capability to retrieve user's Routing Protocol Changes

Use --request flag to provide optional JSON payload fields.

```
equinix fabricv4 routing-protocols get-connection-routing-protocols-changes [flags]
```

### Options

```
      --connection-id string         Connection Id (required)
  -h, --help                         help for get-connection-routing-protocols-changes
      --limit int                    limit field
      --offset int                   offset field
      --request string               JSON payload for additional optional fields not exposed as flags
      --routing-protocol-id string   Routing Protocol Id (required)
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug           Enable debug logging for HTTP requests
  -f, --format string   Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix fabricv4 routing-protocols](equinix_fabricv4_routing-protocols.md)	 - Manage routing-protocols resources

