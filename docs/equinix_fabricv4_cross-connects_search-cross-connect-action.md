## equinix fabricv4 cross-connects search-cross-connect-action

Execute search-cross-connect-action operation

### Synopsis

Execute the search-cross-connect-action operation on this service.

Use --request flag to provide a JSON payload for the request body.
Example: --request '{"field":"value"}'

The command accepts parameters based on the SDK method signature.

```
equinix fabricv4 cross-connects search-cross-connect-action [flags]
```

### Options

```
      --cross-connect-action-search-request-additional-properties string              cross-connect-action-search-request-additional-properties (JSON)
      --cross-connect-action-search-request-filter-additional-properties string       cross-connect-action-search-request-filter-additional-properties (JSON)
      --cross-connect-action-search-request-filter-and string                         cross-connect-action-search-request-filter-and (JSON array)
      --cross-connect-action-search-request-pagination-additional-properties string   cross-connect-action-search-request-pagination-additional-properties (JSON)
      --cross-connect-action-search-request-pagination-limit int                      Number of elements to be requested per page. Number must be between 1 and 100, and the default is 20.
      --cross-connect-action-search-request-pagination-offset int                     Index of the first element.
      --cross-connect-action-search-request-sort string                               cross-connect-action-search-request-sort (JSON array)
  -h, --help[=true|false]                                                             help for search-cross-connect-action
      --param-1 string                                                                param-1 (required)
      --request string                                                                JSON payload for additional optional fields not exposed as flags
```

### Options inherited from parent commands

```
      --config string        config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug[=true|false]   Enable debug logging for HTTP requests
  -f, --format string        Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix fabricv4 cross-connects](equinix_fabricv4_cross-connects.md)	 - Manage cross-connects resources

