## equinix fabricv4 route-filter-rules patch-route-filter-rule-by-uuid

Patch Route Filter Rule

### Synopsis

This API provides capability to partially update a Route Filter Rule

Use --request flag to provide optional JSON payload fields.

```
equinix fabricv4 route-filter-rules patch-route-filter-rule-by-uuid [flags]
```

### Options

```
  -h, --help                                           help for patch-route-filter-rule-by-uuid
      --request string                                 JSON payload for additional optional fields not exposed as flags
      --route-filter-id string                         Route Filters Id (required)
      --route-filter-rule-id string                    Route  Filter  Rules Id (required)
      --route-filter-rules-patch-request-item string   route-filter-rules-patch-request-item field (JSON or string)
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug           Enable debug logging for HTTP requests
  -f, --format string   Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix fabricv4 route-filter-rules](equinix_fabricv4_route-filter-rules.md)	 - Manage route-filter-rules resources

