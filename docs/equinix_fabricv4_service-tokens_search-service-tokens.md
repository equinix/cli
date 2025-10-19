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
      --limit float                                                            limit field (required)
      --offset float                                                           offset field (required)
      --request string                                                         JSON payload for additional optional fields not exposed as flags
      --service-token-search-request-additional-properties string              service-token-search-request-additional-properties (required) (JSON)
      --service-token-search-request-filter-additional-properties string       service-token-search-request-filter-additional-properties (JSON)
      --service-token-search-request-filter-and string                         service-token-search-request-filter-and (JSON array)
      --service-token-search-request-filter-operator string                    service-token-search-request-filter-operator
      --service-token-search-request-filter-property string                    service-token-search-request-filter-property
      --service-token-search-request-filter-values string                      service-token-search-request-filter-values (JSON array)
      --service-token-search-request-pagination-additional-properties string   service-token-search-request-pagination-additional-properties (JSON)
      --service-token-search-request-pagination-limit int                      service-token-search-request-pagination-limit
      --service-token-search-request-pagination-offset int                     service-token-search-request-pagination-offset
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/equinix/equinix.yaml)
      --debug           Enable debug mode to show HTTP requests and responses
```

### SEE ALSO

* [equinix fabricv4 service-tokens](equinix_fabricv4_service-tokens.md)	 - Manage service-tokens resources

