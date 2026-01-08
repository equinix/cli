## equinix fabricv4 route-filter-rules get-route-filter-rule-change-by-uuid

Get Change By ID

### Synopsis

This API provides capability to retrieve a specific Route Filter Rule's Changes

Use --request flag to provide optional JSON payload fields.

```
equinix fabricv4 route-filter-rules get-route-filter-rule-change-by-uuid [flags]
```

### Options

```
      --change-id string              Route Filter Rule Change UUID (required)
  -h, --help                          help for get-route-filter-rule-change-by-uuid
      --request string                JSON payload for request body fields
      --route-filter-id string        Route Filters Id (required)
      --route-filter-rule-id string   Route  Filter  Rules Id (required)
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug           Enable debug logging for HTTP requests
  -f, --format string   Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix fabricv4 route-filter-rules](equinix_fabricv4_route-filter-rules.md)	 - Manage route-filter-rules resources

