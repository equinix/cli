## equinix fabricv4 streams update-stream-by-uuid

Update Stream

### Synopsis

This API provides capability to update user's stream

Use --request flag to provide optional JSON payload fields.

```
equinix fabricv4 streams update-stream-by-uuid [flags]
```

### Options

```
  -h, --help                                              help for update-stream-by-uuid
      --request string                                    JSON payload for additional optional fields not exposed as flags
      --stream-id string                                  Stream UUID (required)
      --stream-put-request-additional-properties string   stream-put-request-additional-properties (JSON)
      --stream-put-request-description string             Customer-provided stream description
      --stream-put-request-name string                    Customer-provided stream name
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/equinix/equinix.yaml)
      --debug           Enable debug mode to show HTTP requests and responses
```

### SEE ALSO

* [equinix fabricv4 streams](equinix_fabricv4_streams.md)	 - Manage streams resources

