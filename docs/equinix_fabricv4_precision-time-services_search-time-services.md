## equinix fabricv4 precision-time-services search-time-services

Execute search-time-services operation

### Synopsis

Execute the search-time-services operation on this service.

Use --request flag to provide a JSON payload for the request body.
Example: --request '{"field":"value"}'

The command accepts parameters based on the SDK method signature.

```
equinix fabricv4 precision-time-services search-time-services [flags]
```

### Options

```
  -h, --help[=true|false]                                                      help for search-time-services
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
      --config string        config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug[=true|false]   Enable debug logging for HTTP requests
  -f, --format string        Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix fabricv4 precision-time-services](equinix_fabricv4_precision-time-services.md)	 - Manage precision-time-services resources

