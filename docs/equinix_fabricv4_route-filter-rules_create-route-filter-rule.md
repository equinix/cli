## equinix fabricv4 route-filter-rules create-route-filter-rule

Create Route Filter Rule

### Synopsis

This API provides capability to create a Route Filter Rule

Use --request flag to provide optional JSON payload fields.

```
equinix fabricv4 route-filter-rules create-route-filter-rule [flags]
```

### Options

```
  -h, --help                                                   help for create-route-filter-rule
      --request string                                         JSON payload for additional optional fields not exposed as flags
      --route-filter-id string                                 Route Filters Id (required)
      --route-filter-rules-base-additional-properties string   route-filter-rules-base-additional-properties (JSON)
      --route-filter-rules-base-description string             Customer-provided Route Filter Rule description
      --route-filter-rules-base-name string                    route-filter-rules-base-name
      --route-filter-rules-base-prefix string                  route-filter-rules-base-prefix
      --route-filter-rules-base-prefix-match string            route-filter-rules-base-prefix-match
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug           Enable debug logging for HTTP requests
  -f, --format string   Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix fabricv4 route-filter-rules](equinix_fabricv4_route-filter-rules.md)	 - Manage route-filter-rules resources

