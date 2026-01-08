## equinix fabricv4 route-filters get-route-filter-changes

Get All Changes

### Synopsis

This API provides capability to retrieve all of a Route Filter's Changes

Use --request flag to provide optional JSON payload fields.

```
equinix fabricv4 route-filters get-route-filter-changes [flags]
```

### Options

```
  -h, --help                     help for get-route-filter-changes
      --limit int                limit field
      --offset int               offset field
      --request string           JSON payload for additional optional fields not exposed as flags
      --route-filter-id string   Route Filters Id (required)
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug           Enable debug logging for HTTP requests
  -f, --format string   Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix fabricv4 route-filters](equinix_fabricv4_route-filters.md)	 - Manage route-filters resources

