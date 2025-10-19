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
  -h, --help                                                    help for update-stream-asset-by-uuid
      --request string                                          JSON payload for additional optional fields not exposed as flags
      --stream-asset-put-request-additional-properties string   stream-asset-put-request-additional-properties (JSON)
      --stream-asset-put-request-metrics-enabled                enable metric
      --stream-id string                                        Stream UUID (required)
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/equinix/equinix.yaml)
      --debug           Enable debug mode to show HTTP requests and responses
```

### SEE ALSO

* [equinix fabricv4 streams](equinix_fabricv4_streams.md)	 - Manage streams resources

