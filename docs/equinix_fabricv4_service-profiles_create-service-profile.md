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
      --service-profile-request-additional-properties string                  service-profile-request-additional-properties (required) (JSON)
      --service-profile-request-allowed-emails string                         service-profile-request-allowed-emails (JSON array)
      --service-profile-request-custom-fields string                          service-profile-request-custom-fields (JSON array)
      --service-profile-request-description string                            service-profile-request-description (required)
      --service-profile-request-href string                                   service-profile-request-href
      --service-profile-request-marketing-info-additional-properties string   service-profile-request-marketing-info-additional-properties (JSON)
      --service-profile-request-marketing-info-logo string                    service-profile-request-marketing-info-logo
      --service-profile-request-marketing-info-process-steps string           service-profile-request-marketing-info-process-steps (JSON array)
      --service-profile-request-marketing-info-promotion                      service-profile-request-marketing-info-promotion
      --service-profile-request-metros string                                 service-profile-request-metros (JSON array)
      --service-profile-request-name string                                   service-profile-request-name (required)
      --service-profile-request-notifications string                          service-profile-request-notifications (JSON array)
      --service-profile-request-ports string                                  service-profile-request-ports (JSON array)
      --service-profile-request-project-additional-properties string          service-profile-request-project-additional-properties (JSON)
      --service-profile-request-project-id string                             service-profile-request-project-id
      --service-profile-request-project-project-id string                     service-profile-request-project-project-id
      --service-profile-request-self-profile                                  service-profile-request-self-profile
      --service-profile-request-tags string                                   service-profile-request-tags (JSON array)
      --service-profile-request-type string                                   service-profile-request-type (required)
      --service-profile-request-uuid string                                   service-profile-request-uuid
      --service-profile-request-virtual-devices string                        service-profile-request-virtual-devices (JSON array)
      --service-profile-request-visibility string                             service-profile-request-visibility
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/equinix/equinix.yaml)
      --debug           Enable debug mode to show HTTP requests and responses
```

### SEE ALSO

* [equinix fabricv4 service-profiles](equinix_fabricv4_service-profiles.md)	 - Manage service-profiles resources

