## equinix metalv1 devices find-bgp-sessions

Retrieve all BGP sessions

### Synopsis

Provides a listing of available BGP sessions for the device.

Use --request flag to provide optional JSON payload fields.

```
equinix metalv1 devices find-bgp-sessions [flags]
```

### Options

```
  -h, --help             help for find-bgp-sessions
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

