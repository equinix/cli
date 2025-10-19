## equinix fabricv4 route-aggregation-rules create-route-aggregation-rules-in-bulk

Bulk RARules

### Synopsis

This API provides capability to create bulk route aggregation rules

Use --request flag to provide optional JSON payload fields.

```
equinix fabricv4 route-aggregation-rules create-route-aggregation-rules-in-bulk [flags]
```

### Options

```
  -h, --help                                                                help for create-route-aggregation-rules-in-bulk
      --request string                                                      JSON payload for additional optional fields not exposed as flags
      --route-aggregation-id string                                         Route Aggregations Id (required)
      --route-aggregation-rules-post-request-additional-properties string   route-aggregation-rules-post-request-additional-properties (required) (JSON)
      --route-aggregation-rules-post-request-data string                    route-aggregation-rules-post-request-data (JSON array)
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/equinix/equinix.yaml)
      --debug           Enable debug mode to show HTTP requests and responses
```

### SEE ALSO

* [equinix fabricv4 route-aggregation-rules](equinix_fabricv4_route-aggregation-rules.md)	 - Manage route-aggregation-rules resources

