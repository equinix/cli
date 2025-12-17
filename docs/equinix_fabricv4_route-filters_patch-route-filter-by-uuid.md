## equinix fabricv4 route-filters patch-route-filter-by-uuid

Patch Route Filter

### Synopsis

This API provides capability to partially update a Route Filter

Use --request flag to provide optional JSON payload fields.

```
equinix fabricv4 route-filters patch-route-filter-by-uuid [flags]
```

### Options

```
  -h, --help                                      help for patch-route-filter-by-uuid
      --request string                            JSON payload for additional optional fields not exposed as flags
      --route-filter-id string                    Route Filters Id (required)
      --route-filters-patch-request-item string   route-filters-patch-request-item field (JSON or string)
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug           Enable debug logging for HTTP requests
  -f, --format string   Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix fabricv4 route-filters](equinix_fabricv4_route-filters.md)	 - Manage route-filters resources

