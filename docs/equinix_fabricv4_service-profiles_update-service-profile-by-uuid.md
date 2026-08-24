## equinix fabricv4 service-profiles update-service-profile-by-uuid

Update Profile

### Synopsis

Update Service Profile by UUID

Use --request flag to provide optional JSON payload fields.

```
equinix fabricv4 service-profiles update-service-profile-by-uuid [flags]
```

### Options

```
  -h, --help                                      help for update-service-profile-by-uuid
      --request string                            JSON payload for additional optional fields not exposed as flags
      --service-profile-id string                 Service Profile UUID (required)
      --service-profile-update-operation string   service-profile-update-operation field (JSON or string)
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug           Enable debug logging for HTTP requests
  -f, --format string   Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix fabricv4 service-profiles](equinix_fabricv4_service-profiles.md)	 - Manage service-profiles resources

