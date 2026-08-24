## equinix fabricv4 application-domains search-app-domains

Search App Domains

### Synopsis

The API provides capability to get list of user's App Domains using search criteria, including optional filtering, pagination and sorting

Use --request flag to provide optional JSON payload fields.

```
equinix fabricv4 application-domains search-app-domains [flags]
```

### Options

```
      --app-domain-search-request-additional-properties string              app-domain-search-request-additional-properties (JSON)
      --app-domain-search-request-filter-additional-properties string       app-domain-search-request-filter-additional-properties (JSON)
      --app-domain-search-request-filter-and string                         app-domain-search-request-filter-and (JSON array)
      --app-domain-search-request-pagination-additional-properties string   app-domain-search-request-pagination-additional-properties (JSON)
      --app-domain-search-request-pagination-limit int                      Number of elements to be requested per page. Number must be between 1 and 100, and the default is 20.
      --app-domain-search-request-pagination-offset int                     Index of the first element.
      --app-domain-search-request-sort string                               app-domain-search-request-sort (JSON array)
  -h, --help                                                                help for search-app-domains
      --request string                                                      JSON payload for additional optional fields not exposed as flags
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug           Enable debug logging for HTTP requests
  -f, --format string   Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix fabricv4 application-domains](equinix_fabricv4_application-domains.md)	 - Manage application-domains resources

