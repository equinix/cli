## equinix fabricv4 route-aggregation-rules patch-route-aggregation-rule-by-uuid

PatchRARule

### Synopsis

This API provides capability to partially update a Route Aggregation Rule

Use --request flag to provide optional JSON payload fields.

```
equinix fabricv4 route-aggregation-rules patch-route-aggregation-rule-by-uuid [flags]
```

### Options

```
  -h, --help                                                help for patch-route-aggregation-rule-by-uuid
      --request string                                      JSON payload for additional optional fields not exposed as flags
      --route-aggregation-id string                         Route Aggregations Id (required)
      --route-aggregation-rule-id string                    Route Aggregation Rules Id (required)
      --route-aggregation-rules-patch-request-item string   route-aggregation-rules-patch-request-item field (JSON or string)
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/equinix/equinix.yaml)
      --debug           Enable debug mode to show HTTP requests and responses
```

### SEE ALSO

* [equinix fabricv4 route-aggregation-rules](equinix_fabricv4_route-aggregation-rules.md)	 - Manage route-aggregation-rules resources

