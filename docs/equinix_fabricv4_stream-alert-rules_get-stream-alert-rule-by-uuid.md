## equinix fabricv4 stream-alert-rules get-stream-alert-rule-by-uuid

Get Stream Alert Rules

### Synopsis

This API provides capability to get user's stream alert rules

Use --request flag to provide optional JSON payload fields.

```
equinix fabricv4 stream-alert-rules get-stream-alert-rule-by-uuid [flags]
```

### Options

```
      --alert-rule-id string   alert rule UUID (required)
  -h, --help                   help for get-stream-alert-rule-by-uuid
      --request string         Raw JSON payload for request body fields
      --stream-id string       Stream UUID (required)
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/equinix/equinix.yaml)
      --debug           Enable debug mode to show HTTP requests and responses
```

### SEE ALSO

* [equinix fabricv4 stream-alert-rules](equinix_fabricv4_stream-alert-rules.md)	 - Manage stream-alert-rules resources

