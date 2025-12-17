## equinix fabricv4 service-profiles get-service-profile-by-uuid

Get Profile

### Synopsis

Get service profile by UUID. View Point parameter if set to zSide will give seller's view of the profile otherwise buyer's view.

Use --request flag to provide optional JSON payload fields.

```
equinix fabricv4 service-profiles get-service-profile-by-uuid [flags]
```

### Options

```
  -h, --help                        help for get-service-profile-by-uuid
      --request string              JSON payload for additional optional fields not exposed as flags
      --service-profile-id string   Service Profile UUID (required)
      --view-point string           view-point field
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug           Enable debug logging for HTTP requests
  -f, --format string   Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix fabricv4 service-profiles](equinix_fabricv4_service-profiles.md)	 - Manage service-profiles resources

