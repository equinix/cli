## equinix metalv1 ports create-port-vlan-assignment-batch

Create a new Port-VLAN Assignment management batch

### Synopsis

Create a new asynchronous batch request which handles adding and/or removing the VLANs to which the port is assigned.

Use --request flag to provide optional JSON payload fields.

```
equinix metalv1 ports create-port-vlan-assignment-batch [flags]
```

### Options

```
  -h, --help                                                                   help for create-port-vlan-assignment-batch
      --id string                                                              Port UUID (required)
      --include string                                                         include field (JSON or string)
      --port-vlan-assignment-batch-create-input-additional-properties string   port-vlan-assignment-batch-create-input-additional-properties (JSON)
      --port-vlan-assignment-batch-create-input-vlan_assignments string        port-vlan-assignment-batch-create-input-vlan_assignments (JSON array)
      --request string                                                         JSON payload for additional optional fields not exposed as flags
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug           Enable debug logging for HTTP requests
  -f, --format string   Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix metalv1 ports](equinix_metalv1_ports.md)	 - Manage ports resources

