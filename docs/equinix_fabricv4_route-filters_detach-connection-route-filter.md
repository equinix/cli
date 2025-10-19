## equinix fabricv4 route-filters detach-connection-route-filter

Detach Route Filter

### Synopsis

This API provides capability to detach a Route Filter from a Connection

Use --request flag to provide optional JSON payload fields.

```
equinix fabricv4 route-filters detach-connection-route-filter [flags]
```

### Options

```
      --connection-id string     Connection Id (required)
  -h, --help                     help for detach-connection-route-filter
      --request string           Raw JSON payload for request body fields
      --route-filter-id string   Route Filters Id (required)
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/equinix/equinix.yaml)
      --debug           Enable debug mode to show HTTP requests and responses
```

### SEE ALSO

* [equinix fabricv4 route-filters](equinix_fabricv4_route-filters.md)	 - Manage route-filters resources

