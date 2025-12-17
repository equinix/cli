## equinix metalv1 ports delete-native-vlan

Remove native VLAN

### Synopsis

Removes the native VLAN from this port

Use --request flag to provide optional JSON payload fields.

```
equinix metalv1 ports delete-native-vlan [flags]
```

### Options

```
  -h, --help             help for delete-native-vlan
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

