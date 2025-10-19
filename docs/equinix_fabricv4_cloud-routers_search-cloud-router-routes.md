## equinix fabricv4 cloud-routers search-cloud-router-routes

Search Route Table

### Synopsis

The API provides capability to get list of user's Fabric Cloud Router route table entries using search criteria, including optional filtering, pagination and sorting

Use --request flag to provide optional JSON payload fields.

```
equinix fabricv4 cloud-routers search-cloud-router-routes [flags]
```

### Options

```
  -h, --help                                                                       help for search-cloud-router-routes
      --request string                                                             JSON payload for additional optional fields not exposed as flags
      --route-table-entry-search-request-additional-properties string              route-table-entry-search-request-additional-properties (required) (JSON)
      --route-table-entry-search-request-filter-additional-properties string       route-table-entry-search-request-filter-additional-properties (JSON)
      --route-table-entry-search-request-filter-and string                         route-table-entry-search-request-filter-and (JSON array)
      --route-table-entry-search-request-pagination-additional-properties string   route-table-entry-search-request-pagination-additional-properties (JSON)
      --route-table-entry-search-request-pagination-limit int                      route-table-entry-search-request-pagination-limit
      --route-table-entry-search-request-pagination-offset int                     route-table-entry-search-request-pagination-offset
      --route-table-entry-search-request-sort string                               route-table-entry-search-request-sort (JSON array)
      --router-id string                                                           Router UUID (required)
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/equinix/equinix.yaml)
      --debug           Enable debug mode to show HTTP requests and responses
```

### SEE ALSO

* [equinix fabricv4 cloud-routers](equinix_fabricv4_cloud-routers.md)	 - Manage cloud-routers resources

