## equinix fabricv4 stream-alert-rules delete-stream-alert-rule-by-uuid

Update Stream Alert Rules

### Synopsis

This API provides capability to delete a user's stream alert rule

Use --request flag to provide optional JSON payload fields.

```
equinix fabricv4 stream-alert-rules delete-stream-alert-rule-by-uuid [flags]
```

### Options

```
      --alert-rule-id string   alert rule UUID (required)
  -h, --help                   help for delete-stream-alert-rule-by-uuid
      --request string         JSON payload for request body fields
      --stream-id string       Stream UUID (required)
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug           Enable debug logging for HTTP requests
  -f, --format string   Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix fabricv4 stream-alert-rules](equinix_fabricv4_stream-alert-rules.md)	 - Manage stream-alert-rules resources

