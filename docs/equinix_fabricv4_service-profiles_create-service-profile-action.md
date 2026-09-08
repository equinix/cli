## equinix fabricv4 service-profiles create-service-profile-action

Profile Actions

### Synopsis

This API provides capability to accept/reject service profile update requests

Use --request flag to provide optional JSON payload fields.

```
equinix fabricv4 service-profiles create-service-profile-action [flags]
```

### Options

```
  -h, --help                                                          help for create-service-profile-action
      --request string                                                JSON payload for additional optional fields not exposed as flags
      --service-profile-action-request-additional-properties string   service-profile-action-request-additional-properties (JSON)
      --service-profile-action-request-description string             Action description
      --service-profile-action-request-type string                    Action type. Example values: PROFILE_UPDATE_ACCEPTANCE, PROFILE_UPDATE_REJECTION
      --service-profile-id string                                     Service Profile UUID (required)
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug           Enable debug logging for HTTP requests
  -f, --format string   Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix fabricv4 service-profiles](equinix_fabricv4_service-profiles.md)	 - Manage service-profiles resources

