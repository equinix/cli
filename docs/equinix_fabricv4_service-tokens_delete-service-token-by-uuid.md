## equinix fabricv4 service-tokens delete-service-token-by-uuid

Delete Token by uuid

### Synopsis

Delete Service Tokens removes an Equinix Fabric service token corresponding to the specified uuid which are in INACTIVE state.

Use --request flag to provide optional JSON payload fields.

```
equinix fabricv4 service-tokens delete-service-token-by-uuid [flags]
```

### Options

```
  -h, --help                      help for delete-service-token-by-uuid
      --request string            JSON payload for request body fields
      --service-token-id string   Service Token UUID (required)
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug           Enable debug logging for HTTP requests
  -f, --format string   Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix fabricv4 service-tokens](equinix_fabricv4_service-tokens.md)	 - Manage service-tokens resources

