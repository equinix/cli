## equinix fabricv4 route-aggregations get-route-aggregation-changes

Get All Changes

### Synopsis

This API provides capability to retrieve all of a Route Aggregation's Changes

Use --request flag to provide optional JSON payload fields.

```
equinix fabricv4 route-aggregations get-route-aggregation-changes [flags]
```

### Options

```
  -h, --help                          help for get-route-aggregation-changes
      --limit int                     limit field
      --offset int                    offset field
      --request string                JSON payload for additional optional fields not exposed as flags
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

