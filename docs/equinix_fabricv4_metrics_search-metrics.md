## equinix fabricv4 metrics search-metrics

Search Metrics

### Synopsis

This API provides capability to search metrics from a filtered query

Use --request flag to provide optional JSON payload fields.

```
equinix fabricv4 metrics search-metrics [flags]
```

### Options

```
  -h, --help                                                             help for search-metrics
      --metrics-search-request-additional-properties string              metrics-search-request-additional-properties (JSON)
      --metrics-search-request-filter-additional-properties string       metrics-search-request-filter-additional-properties (JSON)
      --metrics-search-request-filter-and string                         metrics-search-request-filter-and (JSON array)
      --metrics-search-request-pagination-additional-properties string   metrics-search-request-pagination-additional-properties (JSON)
      --metrics-search-request-pagination-limit int                      Number of elements to be requested per page. Number must be between 1 and 100, and the default is 20.
      --metrics-search-request-pagination-offset int                     Index of the first element.
      --request string                                                   JSON payload for additional optional fields not exposed as flags
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug           Enable debug logging for HTTP requests
  -f, --format string   Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix fabricv4 metrics](equinix_fabricv4_metrics.md)	 - Manage metrics resources

