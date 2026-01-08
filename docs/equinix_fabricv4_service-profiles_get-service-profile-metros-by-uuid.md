## equinix fabricv4 service-profiles get-service-profile-metros-by-uuid

Get Profile Metros

### Synopsis

Get service profile metros by UUID.

Use --request flag to provide optional JSON payload fields.

```
equinix fabricv4 service-profiles get-service-profile-metros-by-uuid [flags]
```

### Options

```
  -h, --help                        help for get-service-profile-metros-by-uuid
      --limit int                   limit field
      --offset int                  offset field
      --request string              JSON payload for additional optional fields not exposed as flags
      --service-profile-id string   Service Profile UUID (required)
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug           Enable debug logging for HTTP requests
  -f, --format string   Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix fabricv4 service-profiles](equinix_fabricv4_service-profiles.md)	 - Manage service-profiles resources

