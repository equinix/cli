## equinix metalv1 devices delete-device

Delete the device

### Synopsis

Deletes a device and deprovisions it in our datacenter.

Use --request flag to provide optional JSON payload fields.

```
equinix metalv1 devices delete-device [flags]
```

### Options

```
      --force-delete     force-delete field
  -h, --help             help for delete-device
      --id string        Device UUID (required)
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

