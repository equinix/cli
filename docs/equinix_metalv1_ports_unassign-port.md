## equinix metalv1 ports unassign-port

Unassign a port

### Synopsis

Unassign a port for a hardware.

Use --request flag to provide optional JSON payload fields.

```
equinix metalv1 ports unassign-port [flags]
```

### Options

```
  -h, --help                                             help for unassign-port
      --id string                                        Port UUID (required)
      --include string                                   include field (JSON or string)
      --port-assign-input-additional-properties string   port-assign-input-additional-properties (JSON)
      --port-assign-input-vnid string                    Virtual Network ID. May be the UUID of the Virtual Network record, or the VLAN value itself.
      --request string                                   JSON payload for additional optional fields not exposed as flags
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug           Enable debug logging for HTTP requests
  -f, --format string   Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix metalv1 ports](equinix_metalv1_ports.md)	 - Manage ports resources

