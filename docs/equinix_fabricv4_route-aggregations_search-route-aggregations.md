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
      --route-aggregations-search-base-pagination-limit int                      Maximum number of search results returned per page. Number must be between 1 and 100, and the default is 20.
      --route-aggregations-search-base-pagination-next string                    URL relative to the next item in the response.
      --route-aggregations-search-base-pagination-offset int                     Index of the first item returned in the response. The default is 0.
      --route-aggregations-search-base-pagination-previous string                URL relative to the previous item in the response.
      --route-aggregations-search-base-pagination-total int                      Total number of elements returned.
      --route-aggregations-search-base-sort string                               route-aggregations-search-base-sort (JSON array)
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/equinix/equinix.yaml)
      --debug           Enable debug mode to show HTTP requests and responses
```

### SEE ALSO

* [equinix fabricv4 route-aggregations](equinix_fabricv4_route-aggregations.md)	 - Manage route-aggregations resources

