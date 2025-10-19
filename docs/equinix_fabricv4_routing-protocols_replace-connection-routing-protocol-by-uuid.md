## equinix fabricv4 routing-protocols replace-connection-routing-protocol-by-uuid

Replace Protocol

### Synopsis

This API provides capability to replace complete Routing Protocols on a virtual connection

Use --request flag to provide optional JSON payload fields.

```
equinix fabricv4 routing-protocols replace-connection-routing-protocol-by-uuid [flags]
```

### Options

```
      --connection-id string         Connection Id (required)
  -h, --help                         help for replace-connection-routing-protocol-by-uuid
      --request string               JSON payload for request body. Available fields: routing-protocol-base (RoutingProtocolBase)
      --routing-protocol-id string   Routing Protocol Id (required)
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/equinix/equinix.yaml)
      --debug           Enable debug mode to show HTTP requests and responses
```

### SEE ALSO

* [equinix fabricv4 routing-protocols](equinix_fabricv4_routing-protocols.md)	 - Manage routing-protocols resources

