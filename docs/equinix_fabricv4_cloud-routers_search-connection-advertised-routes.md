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
      --connection-id string   Connection Id (required)
  -h, --help                   help for search-connection-advertised-routes
      --request string         JSON payload for request body. Available fields: connection-route-search-request (ConnectionRouteSearchRequest)
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/equinix/equinix.yaml)
      --debug           Enable debug mode to show HTTP requests and responses
```

### SEE ALSO

* [equinix fabricv4 cloud-routers](equinix_fabricv4_cloud-routers.md)	 - Manage cloud-routers resources

