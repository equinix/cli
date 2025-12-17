## equinix fabricv4 routing-protocols get-connection-routing-protocols-change-by-uuid

Get Change By ID

### Synopsis

This API provides capability to retrieve specific Routing Protocol Change

Use --request flag to provide optional JSON payload fields.

```
equinix fabricv4 routing-protocols get-connection-routing-protocols-change-by-uuid [flags]
```

### Options

```
      --change-id string             Routing Protocol Change UUID (required)
      --connection-id string         Connection Id (required)
  -h, --help                         help for get-connection-routing-protocols-change-by-uuid
      --request string               JSON payload for request body fields
      --routing-protocol-id string   Routing Protocol Id (required)
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug           Enable debug logging for HTTP requests
  -f, --format string   Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix fabricv4 routing-protocols](equinix_fabricv4_routing-protocols.md)	 - Manage routing-protocols resources

