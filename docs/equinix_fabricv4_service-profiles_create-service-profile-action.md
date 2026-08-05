## equinix fabricv4 service-profiles create-service-profile-action

Execute create-service-profile-action operation

### Synopsis

Execute the create-service-profile-action operation on this service.

Use --request flag to provide a JSON payload for the request body.
Example: --request '{"field":"value"}'

The command accepts parameters based on the SDK method signature.

```
equinix fabricv4 service-profiles create-service-profile-action [flags]
```

### Options

```
  -h, --help                                                          help for create-service-profile-action
      --profile-id string                                             profile-id (required)
      --request string                                                JSON payload for additional optional fields not exposed as flags
      --service-profile-action-request-additional-properties string   service-profile-action-request-additional-properties (JSON)
      --service-profile-action-request-description string             service-profile-action-request-description
      --service-profile-action-request-type string                    service-profile-action-request-type
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug           Enable debug logging for HTTP requests
  -f, --format string   Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix fabricv4 service-profiles](equinix_fabricv4_service-profiles.md)	 - Manage service-profiles resources

