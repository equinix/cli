## equinix metalv1 ports find-port-vlan-assignments

List Current VLAN assignments for a port

### Synopsis

Show the port's current VLAN assignments, including if this VLAN is set as native, and the current state of the assignment (ex. 'assigned' or 'unassigning')

Use --request flag to provide optional JSON payload fields.

```
equinix metalv1 ports find-port-vlan-assignments [flags]
```

### Options

```
  -h, --help             help for find-port-vlan-assignments
      --id string        Port UUID (required)
      --include string   include field (JSON or string)
      --request string   JSON payload for additional optional fields not exposed as flags
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug           Enable debug logging for HTTP requests
  -f, --format string   Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix metalv1 ports](equinix_metalv1_ports.md)	 - Manage ports resources

