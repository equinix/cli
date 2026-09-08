## equinix fabricv4 application-services search-app-services

Search App Services

### Synopsis

The API provides capability to get list of user's App Services using search criteria, including optional filtering, pagination and sorting

Use --request flag to provide optional JSON payload fields.

```
equinix fabricv4 application-services search-app-services [flags]
```

### Options

```
      --app-service-search-request-additional-properties string              app-service-search-request-additional-properties (JSON)
      --app-service-search-request-filter-additional-properties string       app-service-search-request-filter-additional-properties (JSON)
      --app-service-search-request-filter-and string                         app-service-search-request-filter-and (JSON array)
      --app-service-search-request-pagination-additional-properties string   app-service-search-request-pagination-additional-properties (JSON)
      --app-service-search-request-pagination-limit int                      Number of elements to be requested per page. Number must be between 1 and 100, and the default is 20.
      --app-service-search-request-pagination-offset int                     Index of the first element.
      --app-service-search-request-sort string                               app-service-search-request-sort (JSON array)
  -h, --help                                                                 help for search-app-services
      --request string                                                       JSON payload for additional optional fields not exposed as flags
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug           Enable debug logging for HTTP requests
  -f, --format string   Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix fabricv4 application-services](equinix_fabricv4_application-services.md)	 - Manage application-services resources

