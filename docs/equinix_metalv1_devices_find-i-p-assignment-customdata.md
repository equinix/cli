## equinix metalv1 devices find-i-p-assignment-customdata

Retrieve the custom metadata of an IP Assignment

### Synopsis

Provides the custom metadata stored for this IP Assignment in json format

Use --request flag to provide optional JSON payload fields.

```
equinix metalv1 devices find-i-p-assignment-customdata [flags]
```

### Options

```
  -h, --help                 help for find-i-p-assignment-customdata
      --id string            Ip Assignment UUID (required)
      --instance-id string   Instance UUID (required)
      --request string       JSON payload for request body fields
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug           Enable debug logging for HTTP requests
  -f, --format string   Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix metalv1 devices](equinix_metalv1_devices.md)	 - Manage devices resources

