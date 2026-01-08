## equinix fabricv4 service-profiles create-service-profile

Create Profile

### Synopsis

Create Service Profile creates Equinix Fabric™ Service Profile.

Use --request flag to provide optional JSON payload fields.

```
equinix fabricv4 service-profiles create-service-profile [flags]
```

### Options

```
  -h, --help                                                                  help for create-service-profile
      --request string                                                        JSON payload for additional optional fields not exposed as flags
      --service-profile-request-access-point-type-configs string              service-profile-request-access-point-type-configs (JSON array)
      --service-profile-request-additional-properties string                  service-profile-request-additional-properties (JSON)
      --service-profile-request-allowed-emails string                         service-profile-request-allowed-emails (JSON array)
      --service-profile-request-custom-fields string                          service-profile-request-custom-fields (JSON array)
      --service-profile-request-description string                            User-provided service description should be of maximum length 375
      --service-profile-request-href string                                   Service Profile URI response attribute
      --service-profile-request-marketing-info-additional-properties string   service-profile-request-marketing-info-additional-properties (JSON)
      --service-profile-request-marketing-info-logo string                    Logo file name
      --service-profile-request-marketing-info-process-steps string           service-profile-request-marketing-info-process-steps (JSON array)
      --service-profile-request-marketing-info-promotion                      Profile promotion on marketplace
      --service-profile-request-metros string                                 Derived response attribute. (JSON array)
      --service-profile-request-name string                                   Customer-assigned service profile name
      --service-profile-request-notifications string                          Recipients of notifications on service profile change (JSON array)
      --service-profile-request-ports string                                  service-profile-request-ports (JSON array)
      --service-profile-request-project-additional-properties string          service-profile-request-project-additional-properties (JSON)
      --service-profile-request-project-id string                             service-profile-request-project-id
      --service-profile-request-project-project-id string                     Subscriber-assigned project ID
      --service-profile-request-self-profile                                  response attribute indicates whether the profile belongs to the same organization as the api-invoker.
      --service-profile-request-tags string                                   service-profile-request-tags (JSON array)
      --service-profile-request-type string                                   service-profile-request-type
      --service-profile-request-uuid string                                   Equinix-assigned service profile identifier
      --service-profile-request-virtual-devices string                        service-profile-request-virtual-devices (JSON array)
      --service-profile-request-visibility string                             service-profile-request-visibility
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug           Enable debug logging for HTTP requests
  -f, --format string   Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix fabricv4 service-profiles](equinix_fabricv4_service-profiles.md)	 - Manage service-profiles resources

