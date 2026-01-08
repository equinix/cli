## equinix metalv1 devices find-device-by-id

Retrieve a device

### Synopsis

Type-specific options (such as facility for baremetal devices) will be included as part of the main data structure. State value can be one of: active inactive queued or provisioning

Use --request flag to provide optional JSON payload fields.

```
equinix metalv1 devices find-device-by-id [flags]
```

### Options

```
      --exclude string   exclude field (JSON or string)
  -h, --help             help for find-device-by-id
      --id string        Device UUID (required)
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

* [equinix metalv1 devices](equinix_metalv1_devices.md)	 - Manage devices resources

