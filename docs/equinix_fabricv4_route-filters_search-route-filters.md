## equinix fabricv4 route-filters search-route-filters

Search Route Filters

### Synopsis

This API provides capability to search Route Filters

Use --request flag to provide optional JSON payload fields.

```
equinix fabricv4 route-filters search-route-filters [flags]
```

### Options

```
  -h, --help                                                                help for search-route-filters
      --request string                                                      JSON payload for additional optional fields not exposed as flags
      --route-filters-search-base-additional-properties string              route-filters-search-base-additional-properties (required) (JSON)
      --route-filters-search-base-filter-additional-properties string       route-filters-search-base-filter-additional-properties (JSON)
      --route-filters-search-base-filter-and string                         route-filters-search-base-filter-and (JSON array)
      --route-filters-search-base-pagination-additional-properties string   route-filters-search-base-pagination-additional-properties (JSON)
      --route-filters-search-base-pagination-limit int                      Maximum number of search results returned per page. Number must be between 1 and 100, and the default is 20.
      --route-filters-search-base-pagination-next string                    URL relative to the next item in the response.
      --route-filters-search-base-pagination-offset int                     Index of the first item returned in the response. The default is 0.
      --route-filters-search-base-pagination-previous string                URL relative to the previous item in the response.
      --route-filters-search-base-pagination-total int                      Total number of elements returned.
      --route-filters-search-base-sort string                               route-filters-search-base-sort (JSON array)
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/equinix/equinix.yaml)
      --debug           Enable debug mode to show HTTP requests and responses
```

### SEE ALSO

* [equinix fabricv4 route-filters](equinix_fabricv4_route-filters.md)	 - Manage route-filters resources

