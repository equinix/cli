## equinix fabricv4 stream-subscriptions search-stream-subscriptions

Search Stream Subscriptions

### Synopsis

This API provides capability to search stream subscriptions

Use --request flag to provide optional JSON payload fields.

```
equinix fabricv4 stream-subscriptions search-stream-subscriptions [flags]
```

### Options

```
  -h, --help                                                                         help for search-stream-subscriptions
      --request string                                                               JSON payload for additional optional fields not exposed as flags
      --stream-subscription-search-request-additional-properties string              stream-subscription-search-request-additional-properties (JSON)
      --stream-subscription-search-request-filter-additional-properties string       stream-subscription-search-request-filter-additional-properties (JSON)
      --stream-subscription-search-request-filter-and string                         stream-subscription-search-request-filter-and (JSON array)
      --stream-subscription-search-request-pagination-additional-properties string   stream-subscription-search-request-pagination-additional-properties (JSON)
      --stream-subscription-search-request-pagination-limit int                      Number of elements to be requested per page. Number must be between 1 and 100, and the default is 20.
      --stream-subscription-search-request-pagination-offset int                     Index of the first element.
      --stream-subscription-search-request-sort string                               stream-subscription-search-request-sort (JSON array)
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug           Enable debug logging for HTTP requests
  -f, --format string   Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix fabricv4 stream-subscriptions](equinix_fabricv4_stream-subscriptions.md)	 - Manage stream-subscriptions resources

