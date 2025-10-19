## equinix fabricv4 statistics get-connection-stats-by-port-uuid

Get Stats by uuid **(DEPRECATED)**

### Synopsis

This API provides service-level metrics so that you can view access and gather key information required to manage service subscription sizing and capacity **(DEPRECATED)** Deprecated

Use --request flag to provide optional JSON payload fields.

```
equinix fabricv4 statistics get-connection-stats-by-port-uuid [flags]
```

### Options

```
      --connection-id string   Connection UUID (required)
  -h, --help                   help for get-connection-stats-by-port-uuid
      --request string         JSON payload for request body. Available fields: end-date-time (Time), start-date-time (Time), view-point (ViewPoint)
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/equinix/equinix.yaml)
      --debug           Enable debug mode to show HTTP requests and responses
```

### SEE ALSO

* [equinix fabricv4 statistics](equinix_fabricv4_statistics.md)	 - Manage statistics resources

