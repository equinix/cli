## equinix metalv1 capacity check-capacity-for-metro

Check capacity for a metro

### Synopsis

Validates if a deploy can be fulfilled in a metro.

Use --request flag to provide optional JSON payload fields.

```
equinix metalv1 capacity check-capacity-for-metro [flags]
```

### Options

```
      --capacity-input-additional-properties string   capacity-input-additional-properties (JSON)
      --capacity-input-servers string                 capacity-input-servers (JSON array)
  -h, --help                                          help for check-capacity-for-metro
      --request string                                JSON payload for additional optional fields not exposed as flags
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug           Enable debug logging for HTTP requests
  -f, --format string   Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix metalv1 capacity](equinix_metalv1_capacity.md)	 - Manage capacity resources

