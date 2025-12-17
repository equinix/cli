## equinix fabricv4 route-aggregations patch-route-aggregation-by-uuid

Patch Aggregation

### Synopsis

This API provides capability to partially update a Route Aggregation

Use --request flag to provide optional JSON payload fields.

```
equinix fabricv4 route-aggregations patch-route-aggregation-by-uuid [flags]
```

### Options

```
  -h, --help                                           help for patch-route-aggregation-by-uuid
      --request string                                 JSON payload for additional optional fields not exposed as flags
      --route-aggregation-id string                    Route Aggregations Id (required)
      --route-aggregations-patch-request-item string   route-aggregations-patch-request-item field (JSON or string)
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug           Enable debug logging for HTTP requests
  -f, --format string   Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix fabricv4 route-aggregations](equinix_fabricv4_route-aggregations.md)	 - Manage route-aggregations resources

