## equinix fabricv4 streams update-stream-asset-by-uuid

Attach Asset

### Synopsis

This API provides capability to attach an asset to a stream

Use --request flag to provide optional JSON payload fields.

```
equinix fabricv4 streams update-stream-asset-by-uuid [flags]
```

### Options

```
      --asset string                                            asset (required)
      --asset-id string                                         asset UUID (required)
  -h, --help[=true|false]                                       help for update-stream-asset-by-uuid
      --request string                                          JSON payload for additional optional fields not exposed as flags
      --stream-asset-put-request-additional-properties string   stream-asset-put-request-additional-properties (JSON)
      --stream-asset-put-request-metrics-enabled[=true|false]   enable metric
      --stream-id string                                        Stream UUID (required)
```

### Options inherited from parent commands

```
      --config string        config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug[=true|false]   Enable debug logging for HTTP requests
  -f, --format string        Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix fabricv4 streams](equinix_fabricv4_streams.md)	 - Manage streams resources

