## equinix metalv1 interconnections list-interconnection-virtual-circuits

List a interconnection's virtual circuits

### Synopsis

List the virtual circuit record(s) associated with a particular interconnection id.

Use --request flag to provide optional JSON payload fields.

```
equinix metalv1 interconnections list-interconnection-virtual-circuits [flags]
```

### Options

```
      --connection-id string   UUID of the interconnection (required)
  -h, --help                   help for list-interconnection-virtual-circuits
      --request string         JSON payload for request body fields
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug           Enable debug logging for HTTP requests
  -f, --format string   Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix metalv1 interconnections](equinix_metalv1_interconnections.md)	 - Manage interconnections resources

