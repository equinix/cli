## equinix metalv1 devices create-i-p-assignment

Create an ip assignment

### Synopsis

Creates an ip assignment for a device.

Use --request flag to provide optional JSON payload fields.

```
equinix metalv1 devices create-i-p-assignment [flags]
```

### Options

```
      --exclude string                                      exclude field (JSON or string)
  -h, --help                                                help for create-i-p-assignment
      --i-p-assignment-input-additional-properties string   i-p-assignment-input-additional-properties (JSON)
      --i-p-assignment-input-address string                 i-p-assignment-input-address
      --i-p-assignment-input-customdata string              i-p-assignment-input-customdata (JSON)
      --id string                                           Device UUID (required)
      --include string                                      include field (JSON or string)
      --request string                                      JSON payload for additional optional fields not exposed as flags
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug           Enable debug logging for HTTP requests
  -f, --format string   Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix metalv1 devices](equinix_metalv1_devices.md)	 - Manage devices resources

