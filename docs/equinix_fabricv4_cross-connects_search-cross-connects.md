## equinix fabricv4 cross-connects search-cross-connects

Execute search-cross-connects operation

### Synopsis

Execute the search-cross-connects operation on this service.

Use --request flag to provide a JSON payload for the request body.
Example: --request '{"field":"value"}'

The command accepts parameters based on the SDK method signature.

```
equinix fabricv4 cross-connects search-cross-connects [flags]
```

### Options

```
      --cross-connect-search-request-additional-properties string              cross-connect-search-request-additional-properties (JSON)
      --cross-connect-search-request-filter-additional-properties string       cross-connect-search-request-filter-additional-properties (JSON)
      --cross-connect-search-request-filter-and string                         cross-connect-search-request-filter-and (JSON array)
      --cross-connect-search-request-pagination-additional-properties string   cross-connect-search-request-pagination-additional-properties (JSON)
      --cross-connect-search-request-pagination-limit int                      Number of elements to be requested per page. Number must be between 1 and 100, and the default is 20.
      --cross-connect-search-request-pagination-offset int                     Index of the first element.
      --cross-connect-search-request-sort string                               cross-connect-search-request-sort (JSON array)
  -h, --help                                                                   help for search-cross-connects
      --request string                                                         JSON payload for additional optional fields not exposed as flags
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug           Enable debug logging for HTTP requests
  -f, --format string   Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix fabricv4 cross-connects](equinix_fabricv4_cross-connects.md)	 - Manage cross-connects resources

