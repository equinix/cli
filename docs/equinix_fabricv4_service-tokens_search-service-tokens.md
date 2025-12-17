## equinix fabricv4 service-tokens search-service-tokens

Search servicetokens

### Synopsis

The API provides capability to get list of user's servicetokens using search criteria, including optional filtering, pagination and sorting

Use --request flag to provide optional JSON payload fields.

```
equinix fabricv4 service-tokens search-service-tokens [flags]
```

### Options

```
  -h, --help                                                                   help for search-service-tokens
      --limit float                                                            limit field
      --offset float                                                           offset field
      --request string                                                         JSON payload for additional optional fields not exposed as flags
      --service-token-search-request-additional-properties string              service-token-search-request-additional-properties (JSON)
      --service-token-search-request-filter-additional-properties string       service-token-search-request-filter-additional-properties (JSON)
      --service-token-search-request-filter-and string                         service-token-search-request-filter-and (JSON array)
      --service-token-search-request-filter-operator string                    service-token-search-request-filter-operator
      --service-token-search-request-filter-property string                    service-token-search-request-filter-property
      --service-token-search-request-filter-values string                      service-token-search-request-filter-values (JSON array)
      --service-token-search-request-pagination-additional-properties string   service-token-search-request-pagination-additional-properties (JSON)
      --service-token-search-request-pagination-limit int                      Number of elements to be requested per page. Number must be between 1 and 100, and the default is 20.
      --service-token-search-request-pagination-offset int                     Index of the first element.
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug           Enable debug logging for HTTP requests
  -f, --format string   Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix fabricv4 service-tokens](equinix_fabricv4_service-tokens.md)	 - Manage service-tokens resources

