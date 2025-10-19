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
  -h, --help                                                                     help for search-route-aggregations
      --request string                                                           JSON payload for additional optional fields not exposed as flags
      --route-aggregations-search-base-additional-properties string              route-aggregations-search-base-additional-properties (required) (JSON)
      --route-aggregations-search-base-filter-additional-properties string       route-aggregations-search-base-filter-additional-properties (JSON)
      --route-aggregations-search-base-filter-and string                         route-aggregations-search-base-filter-and (JSON array)
      --route-aggregations-search-base-pagination-additional-properties string   route-aggregations-search-base-pagination-additional-properties (JSON)
      --route-aggregations-search-base-pagination-limit int                      route-aggregations-search-base-pagination-limit
      --route-aggregations-search-base-pagination-next string                    route-aggregations-search-base-pagination-next
      --route-aggregations-search-base-pagination-offset int                     route-aggregations-search-base-pagination-offset
      --route-aggregations-search-base-pagination-previous string                route-aggregations-search-base-pagination-previous
      --route-aggregations-search-base-pagination-total int                      route-aggregations-search-base-pagination-total
      --route-aggregations-search-base-sort string                               route-aggregations-search-base-sort (JSON array)
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/equinix/equinix.yaml)
      --debug           Enable debug mode to show HTTP requests and responses
```

### SEE ALSO

* [equinix fabricv4 route-aggregations](equinix_fabricv4_route-aggregations.md)	 - Manage route-aggregations resources

