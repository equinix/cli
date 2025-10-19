## equinix fabricv4 route-aggregation-rules get-route-aggregation-rule-changes

Get All Changes

### Synopsis

This API provides capability to retrieve all of a Route Aggregation Rule's Changes

Use --request flag to provide optional JSON payload fields.

```
equinix fabricv4 route-aggregation-rules get-route-aggregation-rule-changes [flags]
```

### Options

```
  -h, --help                               help for get-route-aggregation-rule-changes
      --limit int                          limit field (required)
      --offset int                         offset field (required)
      --request string                     JSON payload for additional optional fields not exposed as flags
      --route-aggregation-id string        Route Aggregations Id (required)
      --route-aggregation-rule-id string   Route Aggregation Rules Id (required)
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/equinix/equinix.yaml)
      --debug           Enable debug mode to show HTTP requests and responses
```

### SEE ALSO

* [equinix fabricv4 route-aggregation-rules](equinix_fabricv4_route-aggregation-rules.md)	 - Manage route-aggregation-rules resources

