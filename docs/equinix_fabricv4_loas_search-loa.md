## equinix fabricv4 loas search-loa

Search Loas

### Synopsis

The API provides capability to get list of user's Loa using search criteria, including optional filtering, pagination and sorting

Use --request flag to provide optional JSON payload fields.

```
equinix fabricv4 loas search-loa [flags]
```

### Options

```
  -h, --help                                                         help for search-loa
      --loa-search-request-additional-properties string              loa-search-request-additional-properties (JSON)
      --loa-search-request-filter-additional-properties string       loa-search-request-filter-additional-properties (JSON)
      --loa-search-request-filter-and string                         loa-search-request-filter-and (JSON array)
      --loa-search-request-pagination-additional-properties string   loa-search-request-pagination-additional-properties (JSON)
      --loa-search-request-pagination-limit int                      Number of elements to be requested per page. Number must be between 1 and 100, and the default is 20.
      --loa-search-request-pagination-offset int                     Index of the first element.
      --loa-search-request-sort string                               loa-search-request-sort (JSON array)
      --request string                                               JSON payload for additional optional fields not exposed as flags
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug           Enable debug logging for HTTP requests
  -f, --format string   Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix fabricv4 loas](equinix_fabricv4_loas.md)	 - Manage loas resources

