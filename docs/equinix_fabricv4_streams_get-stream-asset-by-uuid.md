## equinix fabricv4 streams get-stream-asset-by-uuid

Get Asset

### Synopsis

This API provides capability to get user's assets attached to a stream

Use --request flag to provide optional JSON payload fields.

```
equinix fabricv4 streams get-stream-asset-by-uuid [flags]
```

### Options

```
      --asset string       asset (required)
      --asset-id string    asset UUID (required)
  -h, --help               help for get-stream-asset-by-uuid
      --request string     JSON payload for request body fields
      --stream-id string   Stream UUID (required)
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug           Enable debug logging for HTTP requests
  -f, --format string   Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix fabricv4 streams](equinix_fabricv4_streams.md)	 - Manage streams resources

