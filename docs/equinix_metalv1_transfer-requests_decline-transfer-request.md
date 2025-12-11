## equinix metalv1 transfer-requests decline-transfer-request

Decline a transfer request

### Synopsis

Decline a transfer request.

Use --request flag to provide optional JSON payload fields.

```
equinix metalv1 transfer-requests decline-transfer-request [flags]
```

### Options

```
  -h, --help             help for decline-transfer-request
      --id string        Transfer request UUID (required)
      --request string   JSON payload for request body fields
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug           Enable debug logging for HTTP requests
```

### SEE ALSO

* [equinix metalv1 transfer-requests](equinix_metalv1_transfer-requests.md)	 - Manage transfer-requests resources

