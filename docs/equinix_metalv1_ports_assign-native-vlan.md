## equinix metalv1 ports assign-native-vlan

Assign a native VLAN

### Synopsis

Sets a virtual network on this port as a "native VLAN". The VLAN must have already been assigned using the using the "Assign a port to a virtual network" operation.

Use --request flag to provide optional JSON payload fields.

```
equinix metalv1 ports assign-native-vlan [flags]
```

### Options

```
  -h, --help             help for assign-native-vlan
      --id string        Port UUID (required)
      --include string   include field (JSON or string)
      --request string   JSON payload for additional optional fields not exposed as flags
      --vnid string      vnid field
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug           Enable debug logging for HTTP requests
  -f, --format string   Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix metalv1 ports](equinix_metalv1_ports.md)	 - Manage ports resources

