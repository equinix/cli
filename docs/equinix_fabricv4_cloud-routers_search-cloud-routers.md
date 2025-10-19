## equinix fabricv4 cloud-routers search-cloud-routers

Search Routers

### Synopsis

The API provides capability to get list of user's Cloud Routers using search criteria, including optional filtering, pagination and sorting

Use --request flag to provide optional JSON payload fields.

```
equinix fabricv4 cloud-routers search-cloud-routers [flags]
```

### Options

```
      --cloud-router-search-request-additional-properties string              cloud-router-search-request-additional-properties (required) (JSON)
      --cloud-router-search-request-filter-additional-properties string       cloud-router-search-request-filter-additional-properties (JSON)
      --cloud-router-search-request-filter-and string                         cloud-router-search-request-filter-and (JSON array)
      --cloud-router-search-request-pagination-additional-properties string   cloud-router-search-request-pagination-additional-properties (JSON)
      --cloud-router-search-request-pagination-limit int                      cloud-router-search-request-pagination-limit
      --cloud-router-search-request-pagination-offset int                     cloud-router-search-request-pagination-offset
      --cloud-router-search-request-sort string                               cloud-router-search-request-sort (JSON array)
  -h, --help                                                                  help for search-cloud-routers
      --request string                                                        JSON payload for additional optional fields not exposed as flags
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/equinix/equinix.yaml)
      --debug           Enable debug mode to show HTTP requests and responses
```

### SEE ALSO

* [equinix fabricv4 cloud-routers](equinix_fabricv4_cloud-routers.md)	 - Manage cloud-routers resources

