## equinix fabricv4 cloud-routers search-connection-received-routes

Search Received Routes

### Synopsis

The API provides capability to get list of received routes using search criteria, including optional filtering, pagination and sorting

Use --request flag to provide optional JSON payload fields.

```
equinix fabricv4 cloud-routers search-connection-received-routes [flags]
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
  -h, --help                                                                      help for search-connection-received-routes
      --request string                                                            JSON payload for additional optional fields not exposed as flags
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug           Enable debug logging for HTTP requests
  -f, --format string   Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix fabricv4 cloud-routers](equinix_fabricv4_cloud-routers.md)	 - Manage cloud-routers resources

