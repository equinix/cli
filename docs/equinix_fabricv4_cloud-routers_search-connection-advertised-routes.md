## equinix fabricv4 cloud-routers search-connection-advertised-routes

Search Advertised Routes

### Synopsis

The API provides capability to get list of user's advertised routes using search criteria, including optional filtering, pagination and sorting

Use --request flag to provide optional JSON payload fields.

```
equinix fabricv4 cloud-routers search-connection-advertised-routes [flags]
```

### Options

```
      --connection-id string                                                      Connection Id (required)
      --connection-route-search-request-additional-properties string              connection-route-search-request-additional-properties (JSON)
      --connection-route-search-request-filter-additional-properties string       connection-route-search-request-filter-additional-properties (JSON)
      --connection-route-search-request-filter-and string                         connection-route-search-request-filter-and (JSON array)
      --connection-route-search-request-pagination-additional-properties string   connection-route-search-request-pagination-additional-properties (JSON)
      --connection-route-search-request-pagination-limit int                      Number of elements to be requested per page. Number must be between 1 and 100, and the default is 20.
      --connection-route-search-request-pagination-offset int                     Index of the first element.
      --connection-route-search-request-sort string                               connection-route-search-request-sort (JSON array)
  -h, --help                                                                      help for search-connection-advertised-routes
      --request string                                                            JSON payload for additional optional fields not exposed as flags
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/equinix/equinix.yaml)
      --debug           Enable debug mode to show HTTP requests and responses
```

### SEE ALSO

* [equinix fabricv4 cloud-routers](equinix_fabricv4_cloud-routers.md)	 - Manage cloud-routers resources

