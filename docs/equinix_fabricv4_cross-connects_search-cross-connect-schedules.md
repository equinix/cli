## equinix fabricv4 cross-connects search-cross-connect-schedules

Execute search-cross-connect-schedules operation

### Synopsis

Execute the search-cross-connect-schedules operation on this service.

Use --request flag to provide a JSON payload for the request body.
Example: --request '{"field":"value"}'

The command accepts parameters based on the SDK method signature.

```
equinix fabricv4 cross-connects search-cross-connect-schedules [flags]
```

### Options

```
      --cross-connect-schedule-search-request-additional-properties string              cross-connect-schedule-search-request-additional-properties (JSON)
      --cross-connect-schedule-search-request-filter-additional-properties string       cross-connect-schedule-search-request-filter-additional-properties (JSON)
      --cross-connect-schedule-search-request-filter-and string                         cross-connect-schedule-search-request-filter-and (JSON array)
      --cross-connect-schedule-search-request-pagination-additional-properties string   cross-connect-schedule-search-request-pagination-additional-properties (JSON)
      --cross-connect-schedule-search-request-pagination-limit int                      Number of elements to be requested per page. Number must be between 1 and 100, and the default is 20.
      --cross-connect-schedule-search-request-pagination-offset int                     Index of the first element.
      --cross-connect-schedule-search-request-sort string                               cross-connect-schedule-search-request-sort (JSON array)
  -h, --help                                                                            help for search-cross-connect-schedules
      --request string                                                                  JSON payload for additional optional fields not exposed as flags
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug           Enable debug logging for HTTP requests
  -f, --format string   Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix fabricv4 cross-connects](equinix_fabricv4_cross-connects.md)	 - Manage cross-connects resources

