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
      --request string                JSON payload for request body. Available fields: limit (int32), offset (int32)
      --route-aggregation-id string   Route Aggregations Id (required)
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/equinix/equinix.yaml)
      --debug           Enable debug mode to show HTTP requests and responses
```

### SEE ALSO

* [equinix fabricv4 route-aggregations](equinix_fabricv4_route-aggregations.md)	 - Manage route-aggregations resources

