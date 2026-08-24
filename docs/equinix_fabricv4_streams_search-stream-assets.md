## equinix fabricv4 streams search-stream-assets

Search Stream Assets

### Synopsis

This API provides capability to search stream assets

Use --request flag to provide optional JSON payload fields.

```
equinix fabricv4 streams search-stream-assets [flags]
```

### Options

```
  -h, --help                                                                  help for search-stream-assets
      --request string                                                        JSON payload for additional optional fields not exposed as flags
      --stream-asset-search-request-additional-properties string              stream-asset-search-request-additional-properties (JSON)
      --stream-asset-search-request-filter-additional-properties string       stream-asset-search-request-filter-additional-properties (JSON)
      --stream-asset-search-request-filter-and string                         stream-asset-search-request-filter-and (JSON array)
      --stream-asset-search-request-pagination-additional-properties string   stream-asset-search-request-pagination-additional-properties (JSON)
      --stream-asset-search-request-pagination-limit int                      Number of elements to be requested per page. Number must be between 1 and 100, and the default is 20.
      --stream-asset-search-request-pagination-offset int                     Index of the first element.
      --stream-asset-search-request-sort string                               stream-asset-search-request-sort (JSON array)
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug           Enable debug logging for HTTP requests
  -f, --format string   Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix fabricv4 streams](equinix_fabricv4_streams.md)	 - Manage streams resources

