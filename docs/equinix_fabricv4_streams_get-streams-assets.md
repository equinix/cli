## equinix fabricv4 streams get-streams-assets

Get Assets

### Synopsis

This API provides capability to retrieve stream assets

Use --request flag to provide optional JSON payload fields.

```
equinix fabricv4 streams get-streams-assets [flags]
```

### Options

```
  -h, --help                                                                  help for get-streams-assets
      --limit int                                                             limit field (required)
      --offset int                                                            offset field (required)
      --request string                                                        JSON payload for additional optional fields not exposed as flags
      --stream-asset-search-request-additional-properties string              stream-asset-search-request-additional-properties (required) (JSON)
      --stream-asset-search-request-filter-additional-properties string       stream-asset-search-request-filter-additional-properties (JSON)
      --stream-asset-search-request-filter-and string                         stream-asset-search-request-filter-and (JSON array)
      --stream-asset-search-request-pagination-additional-properties string   stream-asset-search-request-pagination-additional-properties (JSON)
      --stream-asset-search-request-pagination-limit int                      stream-asset-search-request-pagination-limit
      --stream-asset-search-request-pagination-offset int                     stream-asset-search-request-pagination-offset
      --stream-asset-search-request-sort string                               stream-asset-search-request-sort (JSON array)
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/equinix/equinix.yaml)
      --debug           Enable debug mode to show HTTP requests and responses
```

### SEE ALSO

* [equinix fabricv4 streams](equinix_fabricv4_streams.md)	 - Manage streams resources

