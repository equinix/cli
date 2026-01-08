## equinix fabricv4 stream-alert-rules get-stream-alert-rules

Get Stream Alert Rules

### Synopsis

This API provides capability to retrieve stream alert rules

Use --request flag to provide optional JSON payload fields.

```
equinix fabricv4 stream-alert-rules get-stream-alert-rules [flags]
```

### Options

```
  -h, --help               help for get-stream-alert-rules
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

* [equinix fabricv4 stream-alert-rules](equinix_fabricv4_stream-alert-rules.md)	 - Manage stream-alert-rules resources

