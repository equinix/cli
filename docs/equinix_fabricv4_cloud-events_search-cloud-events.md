## equinix fabricv4 cloud-events search-cloud-events

Search Cloud Events

### Synopsis

This API provides capability to search cloud events from a filtered query

Use --request flag to provide optional JSON payload fields.

```
equinix fabricv4 cloud-events search-cloud-events [flags]
```

### Options

```
      --cloud-event-search-request-additional-properties string              cloud-event-search-request-additional-properties (JSON)
      --cloud-event-search-request-filter-additional-properties string       cloud-event-search-request-filter-additional-properties (JSON)
      --cloud-event-search-request-filter-and string                         cloud-event-search-request-filter-and (JSON array)
      --cloud-event-search-request-pagination-additional-properties string   cloud-event-search-request-pagination-additional-properties (JSON)
      --cloud-event-search-request-pagination-limit int                      Number of elements to be requested per page. Number must be between 1 and 100, and the default is 20.
      --cloud-event-search-request-pagination-offset int                     Index of the first element.
  -h, --help                                                                 help for search-cloud-events
      --request string                                                       JSON payload for additional optional fields not exposed as flags
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/equinix/equinix.yaml)
      --debug           Enable debug mode to show HTTP requests and responses
```

### SEE ALSO

* [equinix fabricv4 cloud-events](equinix_fabricv4_cloud-events.md)	 - Manage cloud-events resources

