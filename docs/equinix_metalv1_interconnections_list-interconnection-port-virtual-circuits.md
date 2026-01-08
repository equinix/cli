## equinix metalv1 interconnections list-interconnection-port-virtual-circuits

List a interconnection port's virtual circuits

### Synopsis

List the virtual circuit record(s) associatiated with a particular interconnection port.

Use --request flag to provide optional JSON payload fields.

```
equinix metalv1 interconnections list-interconnection-port-virtual-circuits [flags]
```

### Options

```
      --connection-id string   UUID of the interconnection (required)
      --exclude string         exclude field (JSON or string)
  -h, --help                   help for list-interconnection-port-virtual-circuits
      --include string         include field (JSON or string)
      --port-id string         UUID of the interconnection port (required)
      --request string         JSON payload for additional optional fields not exposed as flags
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug           Enable debug logging for HTTP requests
  -f, --format string   Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix metalv1 interconnections](equinix_metalv1_interconnections.md)	 - Manage interconnections resources

