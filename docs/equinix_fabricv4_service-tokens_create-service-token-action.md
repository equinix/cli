## equinix fabricv4 service-tokens create-service-token-action

ServiceToken Actions

### Synopsis

This API provides capability to accept/reject user's servicetokens

Use --request flag to provide optional JSON payload fields.

```
equinix fabricv4 service-tokens create-service-token-action [flags]
```

### Options

```
  -h, --help                                                        help for create-service-token-action
      --request string                                              JSON payload for additional optional fields not exposed as flags
      --service-token-action-request-additional-properties string   service-token-action-request-additional-properties (JSON)
      --service-token-action-request-type string                    service-token-action-request-type
      --service-token-id string                                     Service Token UUID (required)
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug           Enable debug logging for HTTP requests
  -f, --format string   Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix fabricv4 service-tokens](equinix_fabricv4_service-tokens.md)	 - Manage service-tokens resources

