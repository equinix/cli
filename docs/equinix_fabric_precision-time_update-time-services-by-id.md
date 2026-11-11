## equinix fabric precision-time update-time-services-by-id

Update By ID.

### Synopsis

The API provides capability to update Precision Time Service by service id.

Use --request flag to provide optional JSON payload fields.

```
equinix fabric precision-time update-time-services-by-id [flags]
```

### Options

```
  -h, --help                                     help for update-time-services-by-id
      --precision-time-change-operation string   precision-time-change-operation field (JSON or string)
      --request string                           JSON payload for additional optional fields not exposed as flags
      --service-id string                        Service UUID (required)
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug           Enable debug logging for HTTP requests
  -f, --format string   Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix fabric precision-time](equinix_fabric_precision-time.md)	 - Manage precision-time resources

