## equinix fabricv4 stream-subscriptions get-stream-subscriptions

Get Subscriptions

### Synopsis

This API provides capability to retrieve stream subscriptions

Use --request flag to provide optional JSON payload fields.

```
equinix fabricv4 stream-subscriptions get-stream-subscriptions [flags]
```

### Options

```
  -h, --help               help for get-stream-subscriptions
      --limit int          limit field
      --offset int         offset field
      --request string     JSON payload for additional optional fields not exposed as flags
      --stream-id string   Stream UUID (required)
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug           Enable debug logging for HTTP requests
  -f, --format string   Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix fabricv4 stream-subscriptions](equinix_fabricv4_stream-subscriptions.md)	 - Manage stream-subscriptions resources

