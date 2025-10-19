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
  -h, --help                          help for update-service-profile-by-uuid
      --if-match string               if-match field (required)
      --json-patch-operation string   json-patch-operation field (required) (JSON or string)
      --request string                JSON payload for additional optional fields not exposed as flags
      --service-profile-id string     Service Profile UUID (required)
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/equinix/equinix.yaml)
      --debug           Enable debug mode to show HTTP requests and responses
```

### SEE ALSO

* [equinix fabricv4 service-profiles](equinix_fabricv4_service-profiles.md)	 - Manage service-profiles resources

