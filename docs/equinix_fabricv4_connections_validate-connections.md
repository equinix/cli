## equinix fabricv4 connections validate-connections

Validate Connection

### Synopsis

This API provides capability to validate by auth key

Use --request flag to provide optional JSON payload fields.

```
equinix fabricv4 connections validate-connections [flags]
```

### Options

```
  -h, --help                                                   help for validate-connections
      --request string                                         JSON payload for additional optional fields not exposed as flags
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

* [equinix fabricv4 connections](equinix_fabricv4_connections.md)	 - Manage connections resources

