## equinix fabric route-filters get-route-filter-change-by-uuid

Get Change By ID

### Synopsis

This API provides capability to retrieve a specific Route Filter's Changes

Use --request flag to provide optional JSON payload fields.

```
equinix fabric route-filters get-route-filter-change-by-uuid [flags]
```

### Options

```
      --change-id string         Routing Protocol Change UUID (required)
  -h, --help                     help for get-route-filter-change-by-uuid
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

