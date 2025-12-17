## equinix metalv1 ports find-port-vlan-assignment-by-port-id-and-assignment-id

Show a particular Port VLAN assignment's details

### Synopsis

Show the details of a specific Port-VLAN assignment, including the current state and if the VLAN is set as native.

Use --request flag to provide optional JSON payload fields.

```
equinix metalv1 ports find-port-vlan-assignment-by-port-id-and-assignment-id [flags]
```

### Options

```
      --assignment-id string   Assignment ID (required)
  -h, --help                   help for find-port-vlan-assignment-by-port-id-and-assignment-id
      --id string              Port UUID (required)
      --include string         include field (JSON or string)
      --request string         JSON payload for additional optional fields not exposed as flags
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug           Enable debug logging for HTTP requests
  -f, --format string   Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix metalv1 ports](equinix_metalv1_ports.md)	 - Manage ports resources

