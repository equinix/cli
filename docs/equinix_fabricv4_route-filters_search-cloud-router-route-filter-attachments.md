## equinix fabricv4 route-filters search-cloud-router-route-filter-attachments

Execute search-cloud-router-route-filter-attachments operation

### Synopsis

Execute the search-cloud-router-route-filter-attachments operation on this service.

Use --request flag to provide a JSON payload for the request body.
Example: --request '{"field":"value"}'

The command accepts parameters based on the SDK method signature.

```
equinix fabricv4 route-filters search-cloud-router-route-filter-attachments [flags]
```

### Options

```
      --cloud-router-route-filters-search-base-additional-properties string                             cloud-router-route-filters-search-base-additional-properties (JSON)
      --cloud-router-route-filters-search-base-filter-cloud-router-route-filter-and-expression string   cloud-router-route-filters-search-base-filter-cloud-router-route-filter-and-expression (JSON)
      --cloud-router-route-filters-search-base-filter-cloud-router-route-filter-or-expression string    cloud-router-route-filters-search-base-filter-cloud-router-route-filter-or-expression (JSON)
      --cloud-router-route-filters-search-base-pagination-additional-properties string                  cloud-router-route-filters-search-base-pagination-additional-properties (JSON)
      --cloud-router-route-filters-search-base-pagination-limit int                                     Maximum number of search results returned per page. Number must be between 1 and 100, and the default is 20.
      --cloud-router-route-filters-search-base-pagination-next string                                   URL relative to the next item in the response.
      --cloud-router-route-filters-search-base-pagination-offset int                                    Index of the first item returned in the response. The default is 0.
      --cloud-router-route-filters-search-base-pagination-previous string                               URL relative to the previous item in the response.
      --cloud-router-route-filters-search-base-pagination-total int                                     Total number of elements returned.
      --cloud-router-route-filters-search-base-sort string                                              cloud-router-route-filters-search-base-sort (JSON array)
  -h, --help                                                                                            help for search-cloud-router-route-filter-attachments
      --request string                                                                                  JSON payload for additional optional fields not exposed as flags
      --router-id string                                                                                router-id (required)
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug           Enable debug logging for HTTP requests
  -f, --format string   Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix fabricv4 route-filters](equinix_fabricv4_route-filters.md)	 - Manage route-filters resources

