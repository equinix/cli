## equinix fabricv4 ports update-port-by-uuid

Update by UUID

### Synopsis

Update Port by UUID

Use --request flag to provide optional JSON payload fields.

```
equinix fabricv4 ports update-port-by-uuid [flags]
```

### Options

```
      --dry-run                        dry-run field (required)
  -h, --help                           help for update-port-by-uuid
      --port-change-operation string   port-change-operation field (required) (JSON or string)
      --port-id string                 Port UUID (required)
      --request string                 JSON payload for additional optional fields not exposed as flags
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/equinix/equinix.yaml)
      --debug           Enable debug mode to show HTTP requests and responses
```

### SEE ALSO

* [equinix fabricv4 ports](equinix_fabricv4_ports.md)	 - Manage ports resources

