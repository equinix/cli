## equinix fabricv4 application-processors search-app-processors

Execute search-app-processors operation

### Synopsis

Execute the search-app-processors operation on this service.

Use --request flag to provide a JSON payload for the request body.
Example: --request '{"field":"value"}'

The command accepts parameters based on the SDK method signature.

```
equinix fabricv4 application-processors search-app-processors [flags]
```

### Options

```
      --app-processor-search-request-additional-properties string              app-processor-search-request-additional-properties (JSON)
      --app-processor-search-request-filter-additional-properties string       app-processor-search-request-filter-additional-properties (JSON)
      --app-processor-search-request-filter-and string                         app-processor-search-request-filter-and (JSON array)
      --app-processor-search-request-pagination-additional-properties string   app-processor-search-request-pagination-additional-properties (JSON)
      --app-processor-search-request-pagination-limit int                      Number of elements to be requested per page. Number must be between 1 and 100, and the default is 20.
      --app-processor-search-request-pagination-offset int                     Index of the first element.
      --app-processor-search-request-sort string                               app-processor-search-request-sort (JSON array)
  -h, --help[=true|false]                                                      help for search-app-processors
      --request string                                                         JSON payload for additional optional fields not exposed as flags
```

### Options inherited from parent commands

```
      --config string        config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug[=true|false]   Enable debug logging for HTTP requests
  -f, --format string        Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix fabricv4 application-processors](equinix_fabricv4_application-processors.md)	 - Manage application-processors resources

