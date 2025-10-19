## equinix fabricv4 route-aggregations get-connection-route-aggregation-by-uuid

Get Aggregation

### Synopsis

This API provides capability to view a specific Route Aggregation attached to a Connection

Use --request flag to provide optional JSON payload fields.

```
equinix fabricv4 route-aggregations get-connection-route-aggregation-by-uuid [flags]
```

### Options

```
      --connection-id string          Connection Id (required)
  -h, --help                          help for get-connection-route-aggregation-by-uuid
      --request string                JSON payload for request body fields
      --route-aggregation-id string   Route Aggregations Id (required)
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/equinix/equinix.yaml)
      --debug           Enable debug mode to show HTTP requests and responses
```

### SEE ALSO

* [equinix fabricv4 route-aggregations](equinix_fabricv4_route-aggregations.md)	 - Manage route-aggregations resources

