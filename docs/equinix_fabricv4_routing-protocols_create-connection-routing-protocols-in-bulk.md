## equinix fabricv4 routing-protocols create-connection-routing-protocols-in-bulk

Bulk Create Protocol

### Synopsis

This API provides capability to create Routing Protocol for connections

Use --request flag to provide optional JSON payload fields.

```
equinix fabricv4 routing-protocols create-connection-routing-protocols-in-bulk [flags]
```

### Options

```
      --connection-id string                                                    Connection Id (required)
      --connection-routing-protocol-post-request-additional-properties string   connection-routing-protocol-post-request-additional-properties (JSON)
      --connection-routing-protocol-post-request-data string                    Connection routing protocol configuration (JSON array)
  -h, --help                                                                    help for create-connection-routing-protocols-in-bulk
      --request string                                                          JSON payload for additional optional fields not exposed as flags
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug           Enable debug logging for HTTP requests
  -f, --format string   Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix fabricv4 routing-protocols](equinix_fabricv4_routing-protocols.md)	 - Manage routing-protocols resources

