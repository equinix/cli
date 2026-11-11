## equinix fabricv4 events search-cloud-events

Execute search-cloud-events operation

### Synopsis

Execute the search-cloud-events operation on this service.

Use --request flag to provide a JSON payload for the request body.
Example: --request '{"field":"value"}'

The command accepts parameters based on the SDK method signature.

```
equinix fabricv4 events search-cloud-events [flags]
```

### Options

```
      --cloud-event-search-request-additional-properties string              cloud-event-search-request-additional-properties (JSON)
      --cloud-event-search-request-filter-additional-properties string       cloud-event-search-request-filter-additional-properties (JSON)
      --cloud-event-search-request-filter-and string                         cloud-event-search-request-filter-and (JSON array)
      --cloud-event-search-request-pagination-additional-properties string   cloud-event-search-request-pagination-additional-properties (JSON)
      --cloud-event-search-request-pagination-limit int                      Number of elements to be requested per page. Number must be between 1 and 100, and the default is 20.
      --cloud-event-search-request-pagination-offset int                     Index of the first element.
  -h, --help[=true|false]                                                    help for search-cloud-events
      --request string                                                       JSON payload for additional optional fields not exposed as flags
```

### Options inherited from parent commands

```
      --config string        config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug[=true|false]   Enable debug logging for HTTP requests
  -f, --format string        Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix fabricv4 events](equinix_fabricv4_events.md)	 - Manage events resources

