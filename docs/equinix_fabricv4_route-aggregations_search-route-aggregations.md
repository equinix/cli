## equinix fabricv4 route-aggregations search-route-aggregations

Search Aggregations

### Synopsis

This API provides capability to search Route Aggregations

Use --request flag to provide optional JSON payload fields.

```
equinix fabricv4 route-aggregations search-route-aggregations [flags]
```

### Options

```
  -h, --help                                                                        help for search-route-aggregations
      --request string                                                              JSON payload for additional optional fields not exposed as flags
      --route-aggregations-search-request-additional-properties string              route-aggregations-search-request-additional-properties (JSON)
      --route-aggregations-search-request-filter-search-and-expression string       route-aggregations-search-request-filter-search-and-expression (JSON)
      --route-aggregations-search-request-filter-search-or-expression string        route-aggregations-search-request-filter-search-or-expression (JSON)
      --route-aggregations-search-request-pagination-additional-properties string   route-aggregations-search-request-pagination-additional-properties (JSON)
      --route-aggregations-search-request-pagination-limit int                      Number of elements to be requested per page. Number must be between 1 and 100, and the default is 20.
      --route-aggregations-search-request-pagination-offset int                     Index of the first element.
      --route-aggregations-search-request-sort string                               route-aggregations-search-request-sort (JSON array)
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug           Enable debug logging for HTTP requests
  -f, --format string   Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix fabricv4 route-aggregations](equinix_fabricv4_route-aggregations.md)	 - Manage route-aggregations resources

