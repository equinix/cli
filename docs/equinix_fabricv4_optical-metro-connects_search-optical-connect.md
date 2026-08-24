## equinix fabricv4 optical-metro-connects search-optical-connect

Search Optical Metro Connect Services

### Synopsis

Get Optical Metro Connects matching the supplied criteria, with optional filtering, pagination and sorting.

Use --request flag to provide optional JSON payload fields.

```
equinix fabricv4 optical-metro-connects search-optical-connect [flags]
```

### Options

```
  -h, --help                                                                     help for search-optical-connect
      --optical-connect-search-request-additional-properties string              optical-connect-search-request-additional-properties (JSON)
      --optical-connect-search-request-filter-additional-properties string       optical-connect-search-request-filter-additional-properties (JSON)
      --optical-connect-search-request-filter-and string                         optical-connect-search-request-filter-and (JSON array)
      --optical-connect-search-request-pagination-additional-properties string   optical-connect-search-request-pagination-additional-properties (JSON)
      --optical-connect-search-request-pagination-limit int                      Number of elements to be requested per page. Number must be between 1 and 100, and the default is 20.
      --optical-connect-search-request-pagination-offset int                     Index of the first element.
      --optical-connect-search-request-sort string                               optical-connect-search-request-sort (JSON array)
      --request string                                                           JSON payload for additional optional fields not exposed as flags
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug           Enable debug logging for HTTP requests
  -f, --format string   Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix fabricv4 optical-metro-connects](equinix_fabricv4_optical-metro-connects.md)	 - Manage optical-metro-connects resources

