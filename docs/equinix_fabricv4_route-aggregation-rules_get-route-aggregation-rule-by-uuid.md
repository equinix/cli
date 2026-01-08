## equinix fabricv4 route-aggregation-rules get-route-aggregation-rule-by-uuid

GetRARule By UUID

### Synopsis

This API provides capability to view a Route Aggregation Rule by UUID

Use --request flag to provide optional JSON payload fields.

```
equinix fabricv4 route-aggregation-rules get-route-aggregation-rule-by-uuid [flags]
```

### Options

```
  -h, --help                               help for get-route-aggregation-rule-by-uuid
      --request string                     JSON payload for request body fields
      --route-aggregation-id string        Route Aggregations Id (required)
      --route-aggregation-rule-id string   Route Aggregation Rules Id (required)
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug           Enable debug logging for HTTP requests
  -f, --format string   Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix fabricv4 route-aggregation-rules](equinix_fabricv4_route-aggregation-rules.md)	 - Manage route-aggregation-rules resources

