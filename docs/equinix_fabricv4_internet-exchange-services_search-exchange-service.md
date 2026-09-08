## equinix fabricv4 internet-exchange-services search-exchange-service

Search Internet Exchange Service

### Synopsis

The API provides capability to get list of user's Internet Exchange Service using search criteria, including optional filtering, pagination and sorting.

Use --request flag to provide optional JSON payload fields.

```
equinix fabricv4 internet-exchange-services search-exchange-service [flags]
```

### Options

```
      --exchange-service-search-request-additional-properties string                         exchange-service-search-request-additional-properties (JSON)
      --exchange-service-search-request-filter-exchange-service-and-expression string        exchange-service-search-request-filter-exchange-service-and-expression (JSON)
      --exchange-service-search-request-filter-exchange-service-or-expression string         exchange-service-search-request-filter-exchange-service-or-expression (JSON)
      --exchange-service-search-request-filter-exchange-service-property-expression string   exchange-service-search-request-filter-exchange-service-property-expression (JSON)
      --exchange-service-search-request-pagination-additional-properties string              exchange-service-search-request-pagination-additional-properties (JSON)
      --exchange-service-search-request-pagination-limit int                                 Number of elements to be requested per page. Number must be between 1 and 100, and the default is 20.
      --exchange-service-search-request-pagination-offset int                                Index of the first element.
      --exchange-service-search-request-sort string                                          exchange-service-search-request-sort (JSON array)
  -h, --help                                                                                 help for search-exchange-service
      --request string                                                                       JSON payload for additional optional fields not exposed as flags
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug           Enable debug logging for HTTP requests
  -f, --format string   Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix fabricv4 internet-exchange-services](equinix_fabricv4_internet-exchange-services.md)	 - Manage internet-exchange-services resources

