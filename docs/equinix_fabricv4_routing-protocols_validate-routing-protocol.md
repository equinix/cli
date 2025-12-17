## equinix fabricv4 routing-protocols validate-routing-protocol

Validate Subnet

### Synopsis

This API provides capability to validate all subnets associated with any connection in the given FCR

Use --request flag to provide optional JSON payload fields.

```
equinix fabricv4 routing-protocols validate-routing-protocol [flags]
```

### Options

```
  -h, --help                                                   help for validate-routing-protocol
      --request string                                         JSON payload for additional optional fields not exposed as flags
      --router-id string                                       Cloud Router UUID (required)
      --validate-request-additional-properties string          validate-request-additional-properties (JSON)
      --validate-request-filter-additional-properties string   validate-request-filter-additional-properties (JSON)
      --validate-request-filter-and string                     validate-request-filter-and (JSON array)
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug           Enable debug logging for HTTP requests
  -f, --format string   Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix fabricv4 routing-protocols](equinix_fabricv4_routing-protocols.md)	 - Manage routing-protocols resources

