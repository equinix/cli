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
      --connection-id string                                                              Connection Id (required)
  -h, --help                                                                              help for replace-connection-routing-protocol-by-uuid
      --request string                                                                    JSON payload for additional optional fields not exposed as flags
      --routing-protocol-base-routing-protocol-b-g-p-type-additional-properties string    routing-protocol-base-routing-protocol-b-g-p-type-additional-properties (JSON)
      --routing-protocol-base-routing-protocol-b-g-p-type-as-override-enabled             routing-protocol-base-routing-protocol-b-g-p-type-as-override-enabled
      --routing-protocol-base-routing-protocol-b-g-p-type-bfd string                      routing-protocol-base-routing-protocol-b-g-p-type-bfd (JSON)
      --routing-protocol-base-routing-protocol-b-g-p-type-bgp-auth-key string             routing-protocol-base-routing-protocol-b-g-p-type-bgp-auth-key
      --routing-protocol-base-routing-protocol-b-g-p-type-bgp-ipv4 string                 routing-protocol-base-routing-protocol-b-g-p-type-bgp-ipv4 (JSON)
      --routing-protocol-base-routing-protocol-b-g-p-type-bgp-ipv6 string                 routing-protocol-base-routing-protocol-b-g-p-type-bgp-ipv6 (JSON)
      --routing-protocol-base-routing-protocol-b-g-p-type-customer-asn int                routing-protocol-base-routing-protocol-b-g-p-type-customer-asn
      --routing-protocol-base-routing-protocol-b-g-p-type-equinix-asn int                 routing-protocol-base-routing-protocol-b-g-p-type-equinix-asn
      --routing-protocol-base-routing-protocol-b-g-p-type-name string                     routing-protocol-base-routing-protocol-b-g-p-type-name
      --routing-protocol-base-routing-protocol-b-g-p-type-type string                     routing-protocol-base-routing-protocol-b-g-p-type-type
      --routing-protocol-base-routing-protocol-direct-type-additional-properties string   routing-protocol-base-routing-protocol-direct-type-additional-properties (JSON)
      --routing-protocol-base-routing-protocol-direct-type-direct-ipv4 string             routing-protocol-base-routing-protocol-direct-type-direct-ipv4 (JSON)
      --routing-protocol-base-routing-protocol-direct-type-direct-ipv6 string             routing-protocol-base-routing-protocol-direct-type-direct-ipv6 (JSON)
      --routing-protocol-base-routing-protocol-direct-type-name string                    routing-protocol-base-routing-protocol-direct-type-name
      --routing-protocol-base-routing-protocol-direct-type-type string                    routing-protocol-base-routing-protocol-direct-type-type
      --routing-protocol-id string                                                        Routing Protocol Id (required)
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/equinix/equinix.yaml)
      --debug           Enable debug mode to show HTTP requests and responses
```

### SEE ALSO

* [equinix fabricv4 routing-protocols](equinix_fabricv4_routing-protocols.md)	 - Manage routing-protocols resources

