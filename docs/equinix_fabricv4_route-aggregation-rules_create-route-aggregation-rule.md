## equinix fabricv4 route-aggregation-rules create-route-aggregation-rule

Create RARule

### Synopsis

This API provides capability to create a Route Aggregation Rule

Use --request flag to provide optional JSON payload fields.

```
equinix fabricv4 route-aggregation-rules create-route-aggregation-rule [flags]
```

### Options

```
  -h, --help                                                        help for create-route-aggregation-rule
      --request string                                              JSON payload for additional optional fields not exposed as flags
      --route-aggregation-id string                                 Route Aggregations Id (required)
      --route-aggregation-rules-base-additional-properties string   route-aggregation-rules-base-additional-properties (required) (JSON)
      --route-aggregation-rules-base-description string             route-aggregation-rules-base-description
      --route-aggregation-rules-base-name string                    route-aggregation-rules-base-name
      --route-aggregation-rules-base-prefix string                  route-aggregation-rules-base-prefix (required)
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/equinix/equinix.yaml)
      --debug           Enable debug mode to show HTTP requests and responses
```

### SEE ALSO

* [equinix fabricv4 route-aggregation-rules](equinix_fabricv4_route-aggregation-rules.md)	 - Manage route-aggregation-rules resources

