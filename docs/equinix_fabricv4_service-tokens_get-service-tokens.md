## equinix fabricv4 service-tokens get-service-tokens

Get All Tokens

### Synopsis

Get All ServiceTokens creates a list of all Equinix Fabric service tokens associated with the subscriber's account.

Use --request flag to provide optional JSON payload fields.

```
equinix fabricv4 service-tokens get-service-tokens [flags]
```

### Options

```
  -h, --help             help for get-service-tokens
      --limit float      limit field
      --offset float     offset field
      --request string   JSON payload for additional optional fields not exposed as flags
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug           Enable debug logging for HTTP requests
  -f, --format string   Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix fabricv4 service-tokens](equinix_fabricv4_service-tokens.md)	 - Manage service-tokens resources

