## equinix fabricv4 routing-protocols get-connection-routing-protocol-all-bgp-actions

Get BGP Actions

### Synopsis

This API provides capability to get all BGP actions status

Use --request flag to provide optional JSON payload fields.

```
equinix fabricv4 routing-protocols get-connection-routing-protocol-all-bgp-actions [flags]
```

### Options

```
      --connection-id string         Connection Id (required)
  -h, --help                         help for get-connection-routing-protocol-all-bgp-actions
      --limit int                    limit field
      --offset int                   offset field
      --request string               JSON payload for additional optional fields not exposed as flags
      --routing-protocol-id string   Routing Protocol Id (required)
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/equinix/equinix.yaml)
      --debug           Enable debug mode to show HTTP requests and responses
```

### SEE ALSO

* [equinix fabricv4 routing-protocols](equinix_fabricv4_routing-protocols.md)	 - Manage routing-protocols resources

