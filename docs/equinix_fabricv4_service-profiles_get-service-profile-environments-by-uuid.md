## equinix fabricv4 service-profiles get-service-profile-environments-by-uuid

Get Profile Environments

### Synopsis

Get service profile provider environments by UUID <sup color='red'>Beta</sup></font>.

Use --request flag to provide optional JSON payload fields.

```
equinix fabricv4 service-profiles get-service-profile-environments-by-uuid [flags]
```

### Options

```
  -h, --help                        help for get-service-profile-environments-by-uuid
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

