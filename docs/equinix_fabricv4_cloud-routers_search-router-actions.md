## equinix fabricv4 cloud-routers search-router-actions

Search Route Table Actions

### Synopsis

This API provides capability to refresh route table and bgp session summary information

Use --request flag to provide optional JSON payload fields.

```
equinix fabricv4 cloud-routers search-router-actions [flags]
```

### Options

```
      --cloud-router-actions-search-request-additional-properties string              cloud-router-actions-search-request-additional-properties (required) (JSON)
      --cloud-router-actions-search-request-filter-additional-properties string       cloud-router-actions-search-request-filter-additional-properties (JSON)
      --cloud-router-actions-search-request-filter-and string                         cloud-router-actions-search-request-filter-and (JSON array)
      --cloud-router-actions-search-request-pagination-additional-properties string   cloud-router-actions-search-request-pagination-additional-properties (JSON)
      --cloud-router-actions-search-request-pagination-limit int                      cloud-router-actions-search-request-pagination-limit
      --cloud-router-actions-search-request-pagination-offset int                     cloud-router-actions-search-request-pagination-offset
      --cloud-router-actions-search-request-sort string                               cloud-router-actions-search-request-sort (JSON array)
  -h, --help                                                                          help for search-router-actions
      --request string                                                                JSON payload for additional optional fields not exposed as flags
      --router-id string                                                              Router UUID (required)
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/equinix/equinix.yaml)
      --debug           Enable debug mode to show HTTP requests and responses
```

### SEE ALSO

* [equinix fabricv4 cloud-routers](equinix_fabricv4_cloud-routers.md)	 - Manage cloud-routers resources

