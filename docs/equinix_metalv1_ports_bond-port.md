## equinix metalv1 ports bond-port

Enabling bonding

### Synopsis

Enabling bonding for one or all ports

Use --request flag to provide optional JSON payload fields.

```
equinix metalv1 ports bond-port [flags]
```

### Options

```
      --bulk-enable      bulk-enable field
  -h, --help             help for bond-port
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

