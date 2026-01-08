## equinix metalv1 transfer-requests find-transfer-request-by-id

View a transfer request

### Synopsis

Returns a single transfer request.

Use --request flag to provide optional JSON payload fields.

```
equinix metalv1 transfer-requests find-transfer-request-by-id [flags]
```

### Options

```
  -h, --help             help for find-transfer-request-by-id
      --id string        Transfer request UUID (required)
      --include string   include field (JSON or string)
      --request string   JSON payload for additional optional fields not exposed as flags
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug           Enable debug logging for HTTP requests
  -f, --format string   Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix metalv1 transfer-requests](equinix_metalv1_transfer-requests.md)	 - Manage transfer-requests resources

