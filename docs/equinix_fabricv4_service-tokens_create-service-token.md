## equinix fabricv4 service-tokens create-service-token

Create Service Token

### Synopsis

Create Service Tokens generates Equinix Fabric™ service tokens. These tokens authorize users to access protected resources and services.

Use --request flag to provide optional JSON payload fields.

```
equinix fabricv4 service-tokens create-service-token [flags]
```

### Options

```
      --dry-run                                                 dry-run field (required)
  -h, --help                                                    help for create-service-token
      --request string                                          JSON payload for additional optional fields not exposed as flags
      --service-token-account-account-name string               service-token-account-account-name
      --service-token-account-account-number int                service-token-account-account-number
      --service-token-account-additional-properties string      service-token-account-additional-properties (JSON)
      --service-token-account-global-cust-id string             service-token-account-global-cust-id
      --service-token-account-global-org-id string              service-token-account-global-org-id
      --service-token-account-global-organization-name string   service-token-account-global-organization-name
      --service-token-account-org-id int                        service-token-account-org-id
      --service-token-account-organization-name string          service-token-account-organization-name
      --service-token-account-reseller-account-name string      service-token-account-reseller-account-name
      --service-token-account-reseller-account-number int       service-token-account-reseller-account-number
      --service-token-account-reseller-org-id int               service-token-account-reseller-org-id
      --service-token-account-reseller-ucm-id string            service-token-account-reseller-ucm-id
      --service-token-account-ucm-id string                     service-token-account-ucm-id
      --service-token-additional-properties string              service-token-additional-properties (required) (JSON)
      --service-token-changelog-additional-properties string    service-token-changelog-additional-properties (JSON)
      --service-token-changelog-created-by string               service-token-changelog-created-by
      --service-token-changelog-created-by-email string         service-token-changelog-created-by-email
      --service-token-changelog-created-by-full-name string     service-token-changelog-created-by-full-name
      --service-token-changelog-created-date-time string        service-token-changelog-created-date-time (JSON)
      --service-token-changelog-deleted-by string               service-token-changelog-deleted-by
      --service-token-changelog-deleted-by-email string         service-token-changelog-deleted-by-email
      --service-token-changelog-deleted-by-full-name string     service-token-changelog-deleted-by-full-name
      --service-token-changelog-deleted-date-time string        service-token-changelog-deleted-date-time (JSON)
      --service-token-changelog-updated-by string               service-token-changelog-updated-by
      --service-token-changelog-updated-by-email string         service-token-changelog-updated-by-email
      --service-token-changelog-updated-by-full-name string     service-token-changelog-updated-by-full-name
      --service-token-changelog-updated-date-time string        service-token-changelog-updated-date-time (JSON)
      --service-token-connection-a-side string                  service-token-connection-a-side (JSON)
      --service-token-connection-additional-properties string   service-token-connection-additional-properties (JSON)
      --service-token-connection-allow-custom-bandwidth         service-token-connection-allow-custom-bandwidth
      --service-token-connection-allow-remote-connection        service-token-connection-allow-remote-connection
      --service-token-connection-bandwidth-limit int            service-token-connection-bandwidth-limit
      --service-token-connection-href string                    service-token-connection-href
      --service-token-connection-supported-bandwidths string    service-token-connection-supported-bandwidths (JSON array)
      --service-token-connection-type string                    service-token-connection-type
      --service-token-connection-uuid string                    service-token-connection-uuid
      --service-token-connection-z-side string                  service-token-connection-z-side (JSON)
      --service-token-description string                        service-token-description
      --service-token-expiry int                                service-token-expiry
      --service-token-href string                               service-token-href
      --service-token-issuer-side string                        service-token-issuer-side
      --service-token-name string                               service-token-name
      --service-token-notifications string                      service-token-notifications (JSON array)
      --service-token-project-additional-properties string      service-token-project-additional-properties (JSON)
      --service-token-project-project-id string                 service-token-project-project-id
      --service-token-state string                              service-token-state
      --service-token-type string                               service-token-type
      --service-token-uuid string                               service-token-uuid
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/equinix/equinix.yaml)
      --debug           Enable debug mode to show HTTP requests and responses
```

### SEE ALSO

* [equinix fabricv4 service-tokens](equinix_fabricv4_service-tokens.md)	 - Manage service-tokens resources

