## equinix fabricv4 events search-last-op-events

Execute search-last-op-events operation

### Synopsis

Execute the search-last-op-events operation on this service.

Use --request flag to provide a JSON payload for the request body.
Example: --request '{"field":"value"}'

The command accepts parameters based on the SDK method signature.

```
equinix fabricv4 events search-last-op-events [flags]
```

### Options

```
  -h, --help                                                                help for search-last-op-events
      --live-event-search-request-additional-properties string              live-event-search-request-additional-properties (JSON)
      --live-event-search-request-filter-additional-properties string       live-event-search-request-filter-additional-properties (JSON)
      --live-event-search-request-filter-and string                         live-event-search-request-filter-and (JSON array)
      --live-event-search-request-pagination-additional-properties string   live-event-search-request-pagination-additional-properties (JSON)
      --live-event-search-request-pagination-limit int                      Number of elements to be requested per page. Number must be between 1 and 100, and the default is 20.
      --live-event-search-request-pagination-offset int                     Index of the first element.
      --request string                                                      JSON payload for additional optional fields not exposed as flags
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug           Enable debug logging for HTTP requests
  -f, --format string   Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix fabricv4 events](equinix_fabricv4_events.md)	 - Manage events resources

