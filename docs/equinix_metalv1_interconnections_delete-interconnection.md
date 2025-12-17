## equinix metalv1 interconnections delete-interconnection

Delete interconnection

### Synopsis

Delete a interconnection, its associated ports and virtual circuits.

Use --request flag to provide optional JSON payload fields.

```
equinix metalv1 interconnections delete-interconnection [flags]
```

### Options

```
      --connection-id string   Interconnection UUID (required)
      --exclude string         exclude field (JSON or string)
  -h, --help                   help for delete-interconnection
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

* [equinix metalv1 interconnections](equinix_metalv1_interconnections.md)	 - Manage interconnections resources

