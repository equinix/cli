## equinix fabricv4 service-tokens update-service-token-by-uuid

Update Token By ID

### Synopsis

This API provides capability to update user's Service Token

Use --request flag to provide optional JSON payload fields.

```
equinix fabricv4 service-tokens update-service-token-by-uuid [flags]
```

### Options

```
      --dry-run                                 dry-run field (required)
  -h, --help                                    help for update-service-token-by-uuid
      --request string                          JSON payload for additional optional fields not exposed as flags
      --service-token-change-operation string   service-token-change-operation field (required) (JSON or string)
      --service-token-id string                 Service Token UUID (required)
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/equinix/equinix.yaml)
      --debug           Enable debug mode to show HTTP requests and responses
```

### SEE ALSO

* [equinix fabricv4 service-tokens](equinix_fabricv4_service-tokens.md)	 - Manage service-tokens resources

