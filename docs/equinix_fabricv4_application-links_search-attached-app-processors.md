## equinix fabricv4 application-links search-attached-app-processors

Execute search-attached-app-processors operation

### Synopsis

Execute the search-attached-app-processors operation on this service.

Use --request flag to provide a JSON payload for the request body.
Example: --request '{"field":"value"}'

The command accepts parameters based on the SDK method signature.

```
equinix fabricv4 application-links search-attached-app-processors [flags]
```

### Options

```
      --app-link-attach-processor-search-request-additional-properties string              app-link-attach-processor-search-request-additional-properties (JSON)
      --app-link-attach-processor-search-request-filter-additional-properties string       app-link-attach-processor-search-request-filter-additional-properties (JSON)
      --app-link-attach-processor-search-request-filter-and string                         app-link-attach-processor-search-request-filter-and (JSON array)
      --app-link-attach-processor-search-request-pagination-additional-properties string   app-link-attach-processor-search-request-pagination-additional-properties (JSON)
      --app-link-attach-processor-search-request-pagination-limit int                      Number of elements to be requested per page. Number must be between 1 and 100, and the default is 20.
      --app-link-attach-processor-search-request-pagination-offset int                     Index of the first element.
      --app-link-attach-processor-search-request-sort string                               app-link-attach-processor-search-request-sort (JSON array)
  -h, --help                                                                               help for search-attached-app-processors
      --param-1 string                                                                     param-1 (required)
      --request string                                                                     JSON payload for additional optional fields not exposed as flags
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug           Enable debug logging for HTTP requests
  -f, --format string   Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix fabricv4 application-links](equinix_fabricv4_application-links.md)	 - Manage application-links resources

