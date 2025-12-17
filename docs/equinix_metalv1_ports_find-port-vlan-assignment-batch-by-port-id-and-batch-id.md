## equinix metalv1 ports find-port-vlan-assignment-batch-by-port-id-and-batch-id

Retrieve a VLAN Assignment Batch's details

### Synopsis

Returns the details of an existing Port-VLAN Assignment batch, including the list of VLANs to assign or unassign, and the current state of the batch.

Use --request flag to provide optional JSON payload fields.

```
equinix metalv1 ports find-port-vlan-assignment-batch-by-port-id-and-batch-id [flags]
```

### Options

```
      --batch-id string   Batch ID (required)
  -h, --help              help for find-port-vlan-assignment-batch-by-port-id-and-batch-id
      --id string         Port UUID (required)
      --include string    include field (JSON or string)
      --request string    JSON payload for additional optional fields not exposed as flags
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug           Enable debug logging for HTTP requests
  -f, --format string   Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix metalv1 ports](equinix_metalv1_ports.md)	 - Manage ports resources

