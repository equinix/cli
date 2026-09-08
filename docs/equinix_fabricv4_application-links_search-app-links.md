## equinix fabricv4 application-links search-app-links

Search App Links

### Synopsis

The API provides capability to get list of user's App Links using search criteria, including optional filtering, pagination and sorting

Use --request flag to provide optional JSON payload fields.

```
equinix fabricv4 application-links search-app-links [flags]
```

### Options

```
      --app-link-search-request-additional-properties string              app-link-search-request-additional-properties (JSON)
      --app-link-search-request-filter-additional-properties string       app-link-search-request-filter-additional-properties (JSON)
      --app-link-search-request-filter-and string                         app-link-search-request-filter-and (JSON array)
      --app-link-search-request-pagination-additional-properties string   app-link-search-request-pagination-additional-properties (JSON)
      --app-link-search-request-pagination-limit int                      Number of elements to be requested per page. Number must be between 1 and 100, and the default is 20.
      --app-link-search-request-pagination-offset int                     Index of the first element.
      --app-link-search-request-sort string                               app-link-search-request-sort (JSON array)
  -h, --help                                                              help for search-app-links
      --request string                                                    JSON payload for additional optional fields not exposed as flags
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug           Enable debug logging for HTTP requests
  -f, --format string   Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix fabricv4 application-links](equinix_fabricv4_application-links.md)	 - Manage application-links resources

