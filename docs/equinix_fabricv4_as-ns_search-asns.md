## equinix fabricv4 as-ns search-asns

Execute search-asns operation

### Synopsis

Execute the search-asns operation on this service.

Use --request flag to provide a JSON payload for the request body.
Example: --request '{"field":"value"}'

The command accepts parameters based on the SDK method signature.

```
equinix fabricv4 as-ns search-asns [flags]
```

### Options

```
      --asn-search-request-body-additional-properties string              asn-search-request-body-additional-properties (JSON)
      --asn-search-request-body-filter-additional-properties string       asn-search-request-body-filter-additional-properties (JSON)
      --asn-search-request-body-filter-and string                         asn-search-request-body-filter-and (JSON array)
      --asn-search-request-body-filter-operator string                    asn-search-request-body-filter-operator
      --asn-search-request-body-filter-or string                          asn-search-request-body-filter-or (JSON array)
      --asn-search-request-body-filter-property string                    asn-search-request-body-filter-property
      --asn-search-request-body-filter-values string                      asn-search-request-body-filter-values (JSON array)
      --asn-search-request-body-pagination-additional-properties string   asn-search-request-body-pagination-additional-properties (JSON)
      --asn-search-request-body-pagination-limit int                      Number of elements to be requested per page. Number must be between 1 and 100, and the default is 20.
      --asn-search-request-body-pagination-offset int                     Index of the first element.
      --asn-search-request-body-sort string                               asn-search-request-body-sort (JSON array)
  -h, --help[=true|false]                                                 help for search-asns
      --request string                                                    JSON payload for additional optional fields not exposed as flags
```

### Options inherited from parent commands

```
      --config string        config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug[=true|false]   Enable debug logging for HTTP requests
  -f, --format string        Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix fabricv4 as-ns](equinix_fabricv4_as-ns.md)	 - Manage as-ns resources

