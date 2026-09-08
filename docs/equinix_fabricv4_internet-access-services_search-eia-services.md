## equinix fabricv4 internet-access-services search-eia-services

Search for Internet Access Services

### Synopsis

Search for Internet Access Services

Use --request flag to provide optional JSON payload fields.

```
equinix fabricv4 internet-access-services search-eia-services [flags]
```

### Options

```
  -h, --help                                                                     help for search-eia-services
      --internet-access-search-request-additional-properties string              internet-access-search-request-additional-properties (JSON)
      --internet-access-search-request-filter-additional-properties string       internet-access-search-request-filter-additional-properties (JSON)
      --internet-access-search-request-filter-and string                         internet-access-search-request-filter-and (JSON array)
      --internet-access-search-request-filter-operator string                    internet-access-search-request-filter-operator
      --internet-access-search-request-filter-or string                          internet-access-search-request-filter-or (JSON array)
      --internet-access-search-request-filter-property string                    internet-access-search-request-filter-property
      --internet-access-search-request-filter-values string                      internet-access-search-request-filter-values (JSON array)
      --internet-access-search-request-pagination-additional-properties string   internet-access-search-request-pagination-additional-properties (JSON)
      --internet-access-search-request-pagination-limit int                      Number of elements to be requested per page. Number must be between 1 and 100, and the default is 20.
      --internet-access-search-request-pagination-offset int                     Index of the first element.
      --internet-access-search-request-sort string                               internet-access-search-request-sort (JSON array)
      --request string                                                           JSON payload for additional optional fields not exposed as flags
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug           Enable debug logging for HTTP requests
  -f, --format string   Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix fabricv4 internet-access-services](equinix_fabricv4_internet-access-services.md)	 - Manage internet-access-services resources

