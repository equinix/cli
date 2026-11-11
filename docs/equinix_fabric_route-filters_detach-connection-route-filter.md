## equinix fabric route-filters detach-connection-route-filter

Detach Route Filter

### Synopsis

This API provides capability to detach a Route Filter from a Connection

Use --request flag to provide optional JSON payload fields.

```
equinix fabric route-filters detach-connection-route-filter [flags]
```

### Options

```
      --connection-id string     Connection Id (required)
  -h, --help                     help for detach-connection-route-filter
      --request string           JSON payload for request body fields
      --route-filter-id string   Route Filters Id (required)
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug           Enable debug logging for HTTP requests
  -f, --format string   Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix fabric route-filters](equinix_fabric_route-filters.md)	 - Manage route-filters resources

