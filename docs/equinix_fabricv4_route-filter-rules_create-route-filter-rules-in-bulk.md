## equinix fabricv4 route-filter-rules create-route-filter-rules-in-bulk

Bulk Create Route Filter Rules

### Synopsis

This API provides capability to create bulk route filter rules

Use --request flag to provide optional JSON payload fields.

```
equinix fabricv4 route-filter-rules create-route-filter-rules-in-bulk [flags]
```

### Options

```
  -h, --help                                                           help for create-route-filter-rules-in-bulk
      --request string                                                 JSON payload for additional optional fields not exposed as flags
      --route-filter-id string                                         Route Filters Id (required)
      --route-filter-rules-post-request-additional-properties string   route-filter-rules-post-request-additional-properties (required) (JSON)
      --route-filter-rules-post-request-data string                    Route Filter Rule configuration (JSON array)
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/equinix/equinix.yaml)
      --debug           Enable debug mode to show HTTP requests and responses
```

### SEE ALSO

* [equinix fabricv4 route-filter-rules](equinix_fabricv4_route-filter-rules.md)	 - Manage route-filter-rules resources

