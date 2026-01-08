## equinix fabricv4 service-tokens get-service-token-by-uuid

Get Token by uuid

### Synopsis

Get Specified Service Tokens uses the uuid of an Equinix Fabric service token to return details about the token's type, state, location, bandwidth, and other key properties.

Use --request flag to provide optional JSON payload fields.

```
equinix fabricv4 service-tokens get-service-token-by-uuid [flags]
```

### Options

```
  -h, --help                      help for get-service-token-by-uuid
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

