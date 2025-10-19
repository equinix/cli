## equinix fabricv4 route-aggregation-rules delete-route-aggregation-rule-by-uuid

DeleteRARule

### Synopsis

This API provides capability to delete a Route aggregation Rule

Use --request flag to provide optional JSON payload fields.

```
equinix fabricv4 route-aggregation-rules delete-route-aggregation-rule-by-uuid [flags]
```

### Options

```
  -h, --help                               help for delete-route-aggregation-rule-by-uuid
      --request string                     JSON payload for request body fields
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

