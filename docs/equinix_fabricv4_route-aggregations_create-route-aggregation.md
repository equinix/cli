## equinix fabricv4 route-aggregations create-route-aggregation

Create Aggregations

### Synopsis

This API provides capability to create a Route Aggregation

Use --request flag to provide optional JSON payload fields.

```
equinix fabricv4 route-aggregations create-route-aggregation [flags]
```

### Options

```
  -h, --help                                                           help for create-route-aggregation
      --request string                                                 JSON payload for additional optional fields not exposed as flags
      --route-aggregations-base-additional-properties string           route-aggregations-base-additional-properties (required) (JSON)
      --route-aggregations-base-description string                     route-aggregations-base-description
      --route-aggregations-base-name string                            route-aggregations-base-name (required)
      --route-aggregations-base-project-additional-properties string   route-aggregations-base-project-additional-properties (required) (JSON)
      --route-aggregations-base-project-project-id string              route-aggregations-base-project-project-id (required)
      --route-aggregations-base-type string                            route-aggregations-base-type (required)
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/equinix/equinix.yaml)
      --debug           Enable debug mode to show HTTP requests and responses
```

### SEE ALSO

* [equinix fabricv4 route-aggregations](equinix_fabricv4_route-aggregations.md)	 - Manage route-aggregations resources

