## equinix fabricv4 service-profiles service-profile-environment-action

Service Profile Environment Actions

### Synopsis

This API provides capability to perform actions on a service profile environment, such as validating an activation key.<sup color='red'>Beta</sup></font>

Use --request flag to provide optional JSON payload fields.

```
equinix fabricv4 service-profiles service-profile-environment-action [flags]
```

### Options

```
      --environment-action-request-additional-properties string               environment-action-request-additional-properties (JSON)
      --environment-action-request-key-details-account-id string              Account identifier
      --environment-action-request-key-details-additional-properties string   environment-action-request-key-details-additional-properties (JSON)
      --environment-action-request-key-details-bandwidth int                  Bandwidth in Mbps
      --environment-action-request-key-details-provider-id string             AWS Connection identifier
      --environment-action-request-key-details-region string                  Cloud provider region identifier
      --environment-action-request-key-details-value string                   Provider Encoded activation key
      --environment-action-request-type string                                environment-action-request-type
      --environment-id string                                                 Provider Environment Reference (required)
  -h, --help                                                                  help for service-profile-environment-action
      --request string                                                        JSON payload for additional optional fields not exposed as flags
      --service-profile-id string                                             Service Profile UUID (required)
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug           Enable debug logging for HTTP requests
  -f, --format string   Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix fabricv4 service-profiles](equinix_fabricv4_service-profiles.md)	 - Manage service-profiles resources

