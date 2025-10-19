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
      --route-filters-search-base-pagination-limit int                      route-filters-search-base-pagination-limit
      --route-filters-search-base-pagination-next string                    route-filters-search-base-pagination-next
      --route-filters-search-base-pagination-offset int                     route-filters-search-base-pagination-offset
      --route-filters-search-base-pagination-previous string                route-filters-search-base-pagination-previous
      --route-filters-search-base-pagination-total int                      route-filters-search-base-pagination-total
      --route-filters-search-base-sort string                               route-filters-search-base-sort (JSON array)
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/equinix/equinix.yaml)
      --debug           Enable debug mode to show HTTP requests and responses
```

### SEE ALSO

* [equinix fabricv4 route-filters](equinix_fabricv4_route-filters.md)	 - Manage route-filters resources

