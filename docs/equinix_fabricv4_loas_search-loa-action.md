## equinix fabricv4 loas search-loa-action

Search Loa Actions

### Synopsis

The API provides capability to get list of user's Loa Actions using search criteria, including optional filtering, pagination and sorting

Use --request flag to provide optional JSON payload fields.

```
equinix fabricv4 loas search-loa-action [flags]
```

### Options

```
  -h, --help                                                                help for search-loa-action
      --loa-action-search-request-additional-properties string              loa-action-search-request-additional-properties (JSON)
      --loa-action-search-request-filter-additional-properties string       loa-action-search-request-filter-additional-properties (JSON)
      --loa-action-search-request-filter-and string                         loa-action-search-request-filter-and (JSON array)
      --loa-action-search-request-pagination-additional-properties string   loa-action-search-request-pagination-additional-properties (JSON)
      --loa-action-search-request-pagination-limit int                      Number of elements to be requested per page. Number must be between 1 and 100, and the default is 20.
      --loa-action-search-request-pagination-offset int                     Index of the first element.
      --loa-action-search-request-sort string                               loa-action-search-request-sort (JSON array)
      --loa-id string                                                       Loa UUID (required)
      --request string                                                      JSON payload for additional optional fields not exposed as flags
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug           Enable debug logging for HTTP requests
  -f, --format string   Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix fabricv4 loas](equinix_fabricv4_loas.md)	 - Manage loas resources

