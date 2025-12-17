## equinix fabricv4 routing-protocols delete-connection-routing-protocol-by-uuid

Delete Protocol

### Synopsis

This API provides capability to delete Routing Protocols on virtual connection

Use --request flag to provide optional JSON payload fields.

```
equinix fabricv4 routing-protocols delete-connection-routing-protocol-by-uuid [flags]
```

### Options

```
      --connection-id string         Connection Id (required)
  -h, --help                         help for delete-connection-routing-protocol-by-uuid
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

