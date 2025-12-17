## equinix metalv1 ports find-port-vlan-assignment-batches

List the VLAN Assignment Batches for a port

### Synopsis

Show all the VLAN assignment batches that have been created for managing this port's VLAN assignments

Use --request flag to provide optional JSON payload fields.

```
equinix metalv1 ports find-port-vlan-assignment-batches [flags]
```

### Options

```
  -h, --help             help for find-port-vlan-assignment-batches
      --id string        Port UUID (required)
      --request string   JSON payload for request body fields
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug           Enable debug logging for HTTP requests
  -f, --format string   Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix metalv1 ports](equinix_metalv1_ports.md)	 - Manage ports resources

