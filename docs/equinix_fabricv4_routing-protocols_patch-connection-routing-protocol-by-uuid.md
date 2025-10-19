## equinix fabricv4 routing-protocols patch-connection-routing-protocol-by-uuid

Patch Protocol

### Synopsis

This API provides capability to partially update Routing Protocols on a virtual connection

Use --request flag to provide optional JSON payload fields.

```
equinix fabricv4 routing-protocols patch-connection-routing-protocol-by-uuid [flags]
```

### Options

```
      --connection-id string         Connection Id (required)
  -h, --help                         help for patch-connection-routing-protocol-by-uuid
      --request string               Raw JSON payload for optional request fields
      --routing-protocol-id string   Routing Protocol Id (required)
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/equinix/equinix.yaml)
      --debug           Enable debug mode to show HTTP requests and responses
```

### SEE ALSO

* [equinix fabricv4 routing-protocols](equinix_fabricv4_routing-protocols.md)	 - Manage routing-protocols resources

