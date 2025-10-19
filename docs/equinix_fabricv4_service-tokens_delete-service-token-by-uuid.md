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
      --request string            Raw JSON payload for optional request fields
      --service-token-id string   Service Token UUID (required)
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/equinix/equinix.yaml)
      --debug           Enable debug mode to show HTTP requests and responses
```

### SEE ALSO

* [equinix fabricv4 service-tokens](equinix_fabricv4_service-tokens.md)	 - Manage service-tokens resources

