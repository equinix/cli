## equinix fabricv4 ports search-ports

Search ports

### Synopsis

The API provides capability to get list of user's virtual ports using search criteria, including optional filtering, pagination and sorting

Use --request flag to provide optional JSON payload fields.

```
equinix fabricv4 ports search-ports [flags]
```

### Options

```
  -h, --help                                                             help for search-ports
      --port-v4-search-request-additional-properties string              port-v4-search-request-additional-properties (JSON)
      --port-v4-search-request-filter-additional-properties string       port-v4-search-request-filter-additional-properties (JSON)
      --port-v4-search-request-filter-and string                         port-v4-search-request-filter-and (JSON array)
      --port-v4-search-request-filter-operator string                    port-v4-search-request-filter-operator
      --port-v4-search-request-filter-or string                          port-v4-search-request-filter-or (JSON array)
      --port-v4-search-request-filter-property string                    port-v4-search-request-filter-property
      --port-v4-search-request-filter-values string                      port-v4-search-request-filter-values (JSON array)
      --port-v4-search-request-pagination-additional-properties string   port-v4-search-request-pagination-additional-properties (JSON)
      --port-v4-search-request-pagination-limit int                      Number of elements to be requested per page. Number must be between 1 and 100, and the default is 20.
      --port-v4-search-request-pagination-offset int                     Index of the first element.
      --port-v4-search-request-sort string                               port-v4-search-request-sort (JSON array)
      --request string                                                   JSON payload for additional optional fields not exposed as flags
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug           Enable debug logging for HTTP requests
  -f, --format string   Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix fabricv4 ports](equinix_fabricv4_ports.md)	 - Manage ports resources

