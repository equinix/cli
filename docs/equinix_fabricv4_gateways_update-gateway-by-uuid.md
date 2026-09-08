## equinix fabricv4 gateways update-gateway-by-uuid

Update Gateway by ID

### Synopsis

Update Gateway by Uuid

Use --request flag to provide optional JSON payload fields.

```
equinix fabricv4 gateways update-gateway-by-uuid [flags]
```

### Options

```
      --gateway-change-operation string   gateway-change-operation field (JSON or string)
      --gateway-id string                 Gateway UUID (required)
  -h, --help                              help for update-gateway-by-uuid
      --request string                    JSON payload for additional optional fields not exposed as flags
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug           Enable debug logging for HTTP requests
  -f, --format string   Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix fabricv4 gateways](equinix_fabricv4_gateways.md)	 - Manage gateways resources

