## equinix fabricv4 route-aggregation-rules search-route-aggregation-rules

Search Route Aggregation Rules

### Synopsis

This API provides capability to search Route Aggregation Rules

Use --request flag to provide optional JSON payload fields.

```
equinix fabricv4 route-aggregation-rules search-route-aggregation-rules [flags]
```

### Options

```
  -h, --help                                                                                         help for search-route-aggregation-rules
      --request string                                                                               JSON payload for additional optional fields not exposed as flags
      --route-aggregation-id string                                                                  Route Aggregations Id (required)
      --route-aggregation-rules-search-request-additional-properties string                          route-aggregation-rules-search-request-additional-properties (JSON)
      --route-aggregation-rules-search-request-filter-route-aggregation-rule-and-expression string   route-aggregation-rules-search-request-filter-route-aggregation-rule-and-expression (JSON)
      --route-aggregation-rules-search-request-filter-route-aggregation-rule-or-expression string    route-aggregation-rules-search-request-filter-route-aggregation-rule-or-expression (JSON)
      --route-aggregation-rules-search-request-pagination-additional-properties string               route-aggregation-rules-search-request-pagination-additional-properties (JSON)
      --route-aggregation-rules-search-request-pagination-limit int                                  Number of elements to be requested per page. Number must be between 1 and 100, and the default is 20.
      --route-aggregation-rules-search-request-pagination-offset int                                 Index of the first element.
      --route-aggregation-rules-search-request-sort string                                           route-aggregation-rules-search-request-sort (JSON array)
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug           Enable debug logging for HTTP requests
  -f, --format string   Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix fabricv4 route-aggregation-rules](equinix_fabricv4_route-aggregation-rules.md)	 - Manage route-aggregation-rules resources

