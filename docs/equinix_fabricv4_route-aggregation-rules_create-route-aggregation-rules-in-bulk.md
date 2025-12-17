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
      --route-aggregation-rules-post-request-additional-properties string   route-aggregation-rules-post-request-additional-properties (JSON)
      --route-aggregation-rules-post-request-data string                    Route Aggregation Rule configuration (JSON array)
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug           Enable debug logging for HTTP requests
  -f, --format string   Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix fabricv4 route-aggregation-rules](equinix_fabricv4_route-aggregation-rules.md)	 - Manage route-aggregation-rules resources

