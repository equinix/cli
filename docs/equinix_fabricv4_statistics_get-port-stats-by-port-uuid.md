## equinix fabricv4 statistics get-port-stats-by-port-uuid

Get Stats by uuid **(DEPRECATED)**

### Synopsis

This API provides service-level traffic metrics so that you can view access and gather key information required to manage service subscription sizing and capacity.**(DEPRECATED)** Deprecated

Use --request flag to provide optional JSON payload fields.

```
equinix fabricv4 statistics get-port-stats-by-port-uuid [flags]
```

### Options

```
  -h, --help             help for get-port-stats-by-port-uuid
      --port-id string   Port UUID (required)
      --request string   JSON payload for additional optional fields not exposed as flags
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug           Enable debug logging for HTTP requests
  -f, --format string   Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix fabricv4 statistics](equinix_fabricv4_statistics.md)	 - Manage statistics resources

