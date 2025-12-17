## equinix metalv1 devices find-instance-bandwidth

Retrieve an instance bandwidth

### Synopsis

Retrieve an instance bandwidth for a given period of time.

Use --request flag to provide optional JSON payload fields.

```
equinix metalv1 devices find-instance-bandwidth [flags]
```

### Options

```
      --from string      from field
  -h, --help             help for find-instance-bandwidth
      --id string        Device UUID (required)
      --request string   JSON payload for additional optional fields not exposed as flags
      --until string     until field
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug           Enable debug logging for HTTP requests
  -f, --format string   Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix metalv1 devices](equinix_metalv1_devices.md)	 - Manage devices resources

