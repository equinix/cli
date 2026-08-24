## equinix fabricv4 route-filter-rules search-route-filter-rules

Search Route Filter Rules

### Synopsis

This API provides capability to search Route Filter Rules

Use --request flag to provide optional JSON payload fields.

```
equinix fabricv4 route-filter-rules search-route-filter-rules [flags]
```

### Options

```
  -h, --help                                                                               help for search-route-filter-rules
      --request string                                                                     JSON payload for additional optional fields not exposed as flags
      --route-filter-id string                                                             Route Filters Id (required)
      --route-filter-rules-search-request-additional-properties string                     route-filter-rules-search-request-additional-properties (JSON)
      --route-filter-rules-search-request-filter-route-filter-rule-and-expression string   route-filter-rules-search-request-filter-route-filter-rule-and-expression (JSON)
      --route-filter-rules-search-request-filter-route-filter-rule-or-expression string    route-filter-rules-search-request-filter-route-filter-rule-or-expression (JSON)
      --route-filter-rules-search-request-pagination-additional-properties string          route-filter-rules-search-request-pagination-additional-properties (JSON)
      --route-filter-rules-search-request-pagination-limit int                             Number of elements to be requested per page. Number must be between 1 and 100, and the default is 20.
      --route-filter-rules-search-request-pagination-offset int                            Index of the first element.
      --route-filter-rules-search-request-sort string                                      route-filter-rules-search-request-sort (JSON array)
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug           Enable debug logging for HTTP requests
  -f, --format string   Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix fabricv4 route-filter-rules](equinix_fabricv4_route-filter-rules.md)	 - Manage route-filter-rules resources

