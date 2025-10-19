## equinix fabricv4 streams delete-stream-asset-by-uuid

Detach Asset

### Synopsis

This API provides capability to detach an asset from a stream

Use --request flag to provide optional JSON payload fields.

```
equinix fabricv4 streams delete-stream-asset-by-uuid [flags]
```

### Options

```
      --asset string       asset (required)
      --asset-id string    asset UUID (required)
  -h, --help               help for delete-stream-asset-by-uuid
      --request string     Raw JSON payload for optional request fields
      --stream-id string   Stream UUID (required)
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/equinix/equinix.yaml)
      --debug           Enable debug mode to show HTTP requests and responses
```

### SEE ALSO

* [equinix fabricv4 streams](equinix_fabricv4_streams.md)	 - Manage streams resources

