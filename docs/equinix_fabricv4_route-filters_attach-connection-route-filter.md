## equinix fabricv4 route-filters attach-connection-route-filter

Attach Route Filter

### Synopsis

This API provides capability to attach a Route Filter to a Connection

Use --request flag to provide optional JSON payload fields.

```
equinix fabricv4 route-filters attach-connection-route-filter [flags]
```

### Options

```
      --connection-id string                                         Connection Id (required)
      --connection-route-filters-base-additional-properties string   connection-route-filters-base-additional-properties (JSON)
      --connection-route-filters-base-direction string               connection-route-filters-base-direction
  -h, --help                                                         help for attach-connection-route-filter
      --request string                                               JSON payload for additional optional fields not exposed as flags
      --route-filter-id string                                       Route Filters Id (required)
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug           Enable debug logging for HTTP requests
  -f, --format string   Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix fabricv4 route-filters](equinix_fabricv4_route-filters.md)	 - Manage route-filters resources

