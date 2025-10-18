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
  -h, --help              help for get-service-token-by-uuid
      --request string    Raw JSON payload for optional request fields
      --token-id string   token-id (required)
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/equinix/equinix.yaml)
```

### SEE ALSO

* [equinix fabricv4 service-tokens](equinix_fabricv4_service-tokens.md)	 - Manage service-tokens resources

