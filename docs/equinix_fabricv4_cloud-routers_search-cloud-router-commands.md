## equinix fabricv4 cloud-routers search-cloud-router-commands

Search Commands

### Synopsis

This API provides capability to search commands

Use --request flag to provide optional JSON payload fields.

```
equinix fabricv4 cloud-routers search-cloud-router-commands [flags]
```

### Options

```
      --cloud-router-command-search-request-additional-properties string              cloud-router-command-search-request-additional-properties (JSON)
      --cloud-router-command-search-request-filter-additional-properties string       cloud-router-command-search-request-filter-additional-properties (JSON)
      --cloud-router-command-search-request-filter-and string                         cloud-router-command-search-request-filter-and (JSON array)
      --cloud-router-command-search-request-pagination-additional-properties string   cloud-router-command-search-request-pagination-additional-properties (JSON)
      --cloud-router-command-search-request-pagination-limit int                      Number of elements to be requested per page. Number must be between 1 and 100, and the default is 20.
      --cloud-router-command-search-request-pagination-offset int                     Index of the first element.
      --cloud-router-command-search-request-sort string                               cloud-router-command-search-request-sort (JSON array)
  -h, --help                                                                          help for search-cloud-router-commands
      --request string                                                                JSON payload for additional optional fields not exposed as flags
      --router-id string                                                              Router UUID (required)
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug           Enable debug logging for HTTP requests
  -f, --format string   Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix fabricv4 cloud-routers](equinix_fabricv4_cloud-routers.md)	 - Manage cloud-routers resources

