## equinix fabricv4 route-aggregations detach-connection-route-aggregation

Detach Aggregation

### Synopsis

This API provides capability to detach a Route Aggregation from a Connection

Use --request flag to provide optional JSON payload fields.

```
equinix fabricv4 route-aggregations detach-connection-route-aggregation [flags]
```

### Options

```
      --connection-id string          Connection Id (required)
  -h, --help                          help for detach-connection-route-aggregation
      --request string                JSON payload for request body fields
      --route-aggregation-id string   Route Aggregations Id (required)
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug           Enable debug logging for HTTP requests
  -f, --format string   Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix fabricv4 route-aggregations](equinix_fabricv4_route-aggregations.md)	 - Manage route-aggregations resources

