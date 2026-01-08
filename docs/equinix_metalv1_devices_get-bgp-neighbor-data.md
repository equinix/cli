## equinix metalv1 devices get-bgp-neighbor-data

Retrieve BGP neighbor data for this device

### Synopsis

Provides a summary of the BGP neighbor data associated to the BGP sessions for this device.

Use --request flag to provide optional JSON payload fields.

```
equinix metalv1 devices get-bgp-neighbor-data [flags]
```

### Options

```
  -h, --help             help for get-bgp-neighbor-data
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

