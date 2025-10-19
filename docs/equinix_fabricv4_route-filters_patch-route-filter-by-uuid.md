## equinix fabricv4 route-filters patch-route-filter-by-uuid

Patch Route Filter

### Synopsis

This API provides capability to partially update a Route Filter

Use --request flag to provide optional JSON payload fields.

```
equinix fabricv4 route-filters patch-route-filter-by-uuid [flags]
```

### Options

```
  -h, --help                     help for patch-route-filter-by-uuid
      --request string           JSON payload for request body. Available fields: route-filters-patch-request-item (RouteFiltersPatchRequestItem)
      --route-filter-id string   Route Filters Id (required)
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/equinix/equinix.yaml)
      --debug           Enable debug mode to show HTTP requests and responses
```

### SEE ALSO

* [equinix fabricv4 route-filters](equinix_fabricv4_route-filters.md)	 - Manage route-filters resources

