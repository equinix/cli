## equinix fabricv4 ip-blocks search-ip-blocks

Search for IP Blocks

### Synopsis

Search for IP Blocks based on criteria

Use --request flag to provide optional JSON payload fields.

```
equinix fabricv4 ip-blocks search-ip-blocks [flags]
```

### Options

```
  -h, --help                                                                    help for search-ip-blocks
      --ip-blocks-search-request-body-additional-properties string              ip-blocks-search-request-body-additional-properties (JSON)
      --ip-blocks-search-request-body-filter-additional-properties string       ip-blocks-search-request-body-filter-additional-properties (JSON)
      --ip-blocks-search-request-body-filter-and string                         ip-blocks-search-request-body-filter-and (JSON array)
      --ip-blocks-search-request-body-pagination-additional-properties string   ip-blocks-search-request-body-pagination-additional-properties (JSON)
      --ip-blocks-search-request-body-pagination-limit int                      Number of elements to be requested per page. Number must be between 1 and 100, and the default is 20.
      --ip-blocks-search-request-body-pagination-offset int                     Index of the first element.
      --ip-blocks-search-request-body-sort string                               ip-blocks-search-request-body-sort (JSON array)
      --request string                                                          JSON payload for additional optional fields not exposed as flags
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug           Enable debug logging for HTTP requests
  -f, --format string   Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix fabricv4 ip-blocks](equinix_fabricv4_ip-blocks.md)	 - Manage ip-blocks resources

