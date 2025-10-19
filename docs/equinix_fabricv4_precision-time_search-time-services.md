## equinix fabricv4 precision-time search-time-services

Search Time Services

### Synopsis

The API provides capability to get list of user's Time Services using search criteria, including optional filtering, pagination and sorting

Use --request flag to provide optional JSON payload fields.

```
equinix fabricv4 precision-time search-time-services [flags]
```

### Options

```
  -h, --help                                                                   help for search-time-services
      --request string                                                         JSON payload for additional optional fields not exposed as flags
      --time-services-search-request-additional-properties string              time-services-search-request-additional-properties (JSON)
      --time-services-search-request-filter-additional-properties string       time-services-search-request-filter-additional-properties (JSON)
      --time-services-search-request-filter-and string                         time-services-search-request-filter-and (JSON array)
      --time-services-search-request-pagination-additional-properties string   time-services-search-request-pagination-additional-properties (JSON)
      --time-services-search-request-pagination-limit int                      Number of elements to be requested per page. Number must be between 1 and 100, and the default is 20.
      --time-services-search-request-pagination-offset int                     Index of the first element.
      --time-services-search-request-sort string                               time-services-search-request-sort (JSON array)
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/equinix/equinix.yaml)
      --debug           Enable debug mode to show HTTP requests and responses
```

### SEE ALSO

* [equinix fabricv4 precision-time](equinix_fabricv4_precision-time.md)	 - Manage precision-time resources

