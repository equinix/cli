## equinix fabricv4 route-filter-rules get-route-filter-rules

Get Route Filter Rules

### Synopsis

This API provides capability to get all Route Filters Rules for Fabric

Use --request flag to provide optional JSON payload fields.

```
equinix fabricv4 route-filter-rules get-route-filter-rules [flags]
```

### Options

```
  -h, --help                     help for get-route-filter-rules
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

* [equinix fabricv4 route-filter-rules](equinix_fabricv4_route-filter-rules.md)	 - Manage route-filter-rules resources

