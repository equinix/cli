## equinix fabricv4 cloud-events search-last-op-events

Search Last Operational Cloud Events

### Synopsis

This API provides capability to search last operational cloud events from a filtered query

Use --request flag to provide optional JSON payload fields.

```
equinix fabricv4 cloud-events search-last-op-events [flags]
```

### Options

```
  -h, --help                                                                       help for search-last-op-events
      --operational-event-search-request-additional-properties string              operational-event-search-request-additional-properties (JSON)
      --operational-event-search-request-filter-additional-properties string       operational-event-search-request-filter-additional-properties (JSON)
      --operational-event-search-request-filter-and string                         operational-event-search-request-filter-and (JSON array)
      --operational-event-search-request-pagination-additional-properties string   operational-event-search-request-pagination-additional-properties (JSON)
      --operational-event-search-request-pagination-limit int                      Number of elements to be requested per page. Number must be between 1 and 100, and the default is 20.
      --operational-event-search-request-pagination-offset int                     Index of the first element.
      --request string                                                             JSON payload for additional optional fields not exposed as flags
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug           Enable debug logging for HTTP requests
  -f, --format string   Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix fabricv4 cloud-events](equinix_fabricv4_cloud-events.md)	 - Manage cloud-events resources

