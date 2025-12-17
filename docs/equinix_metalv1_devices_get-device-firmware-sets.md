## equinix metalv1 devices get-device-firmware-sets

Get Device's associated Firmware Set

### Synopsis

Returns the firmware set associated with the device. If a custom firmware set is associated with the device, then it is returned. Otherwise, if a default firmware set is available it is returned.

Use --request flag to provide optional JSON payload fields.

```
equinix metalv1 devices get-device-firmware-sets [flags]
```

### Options

```
  -h, --help             help for get-device-firmware-sets
      --id string        Device UUID (required)
      --request string   JSON payload for request body fields
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug           Enable debug logging for HTTP requests
  -f, --format string   Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix metalv1 devices](equinix_metalv1_devices.md)	 - Manage devices resources

