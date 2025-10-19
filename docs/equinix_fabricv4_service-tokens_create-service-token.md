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
      --dry-run                                                 dry-run field
  -h, --help                                                    help for create-service-token
      --request string                                          JSON payload for additional optional fields not exposed as flags
      --service-token-account-account-name string               Account name
      --service-token-account-account-number int                Account number
      --service-token-account-additional-properties string      service-token-account-additional-properties (JSON)
      --service-token-account-global-cust-id string             Account name
      --service-token-account-global-org-id string              Global organization identifier
      --service-token-account-global-organization-name string   Global organization name
      --service-token-account-org-id int                        Customer organization identifier
      --service-token-account-organization-name string          Customer organization name
      --service-token-account-reseller-account-name string      Reseller account name
      --service-token-account-reseller-account-number int       Reseller account number
      --service-token-account-reseller-org-id int               Reseller customer organization identifier
      --service-token-account-reseller-ucm-id string            Reseller account ucmId
      --service-token-account-ucm-id string                     Account ucmId
      --service-token-additional-properties string              service-token-additional-properties (JSON)
      --service-token-changelog-additional-properties string    service-token-changelog-additional-properties (JSON)
      --service-token-changelog-created-by string               Created by User Key
      --service-token-changelog-created-by-email string         Created by User Email Address
      --service-token-changelog-created-by-full-name string     Created by User Full Name
      --service-token-changelog-created-date-time string        service-token-changelog-created-date-time (JSON)
      --service-token-changelog-deleted-by string               Deleted by User Key
      --service-token-changelog-deleted-by-email string         Deleted by User Email Address
      --service-token-changelog-deleted-by-full-name string     Deleted by User Full Name
      --service-token-changelog-deleted-date-time string        service-token-changelog-deleted-date-time (JSON)
      --service-token-changelog-updated-by string               Updated by User Key
      --service-token-changelog-updated-by-email string         Updated by User Email Address
      --service-token-changelog-updated-by-full-name string     Updated by User Full Name
      --service-token-changelog-updated-date-time string        service-token-changelog-updated-date-time (JSON)
      --service-token-connection-a-side string                  service-token-connection-a-side (JSON)
      --service-token-connection-additional-properties string   service-token-connection-additional-properties (JSON)
      --service-token-connection-allow-custom-bandwidth         Allow custom bandwidth value
      --service-token-connection-allow-remote-connection        Authorization to connect remotely
      --service-token-connection-bandwidth-limit int            Connection bandwidth limit in Mbps
      --service-token-connection-href string                    An absolute URL that is the subject of the link's context.
      --service-token-connection-supported-bandwidths string    List of permitted bandwidths. (JSON array)
      --service-token-connection-type string                    service-token-connection-type
      --service-token-connection-uuid string                    Equinix-assigned connection identifier
      --service-token-connection-z-side string                  service-token-connection-z-side (JSON)
      --service-token-description string                        Customer-provided service token description
      --service-token-expiry int                                Deprecated
      --service-token-href string                               An absolute URL that is the subject of the link's context.
      --service-token-issuer-side string                        information about token side
                                                                Deprecated
      --service-token-name string                               Customer-provided service token name
      --service-token-notifications string                      Service token related notifications (JSON array)
      --service-token-project-additional-properties string      service-token-project-additional-properties (JSON)
      --service-token-project-project-id string                 Subscriber-assigned project ID
      --service-token-state string                              service-token-state
      --service-token-type string                               service-token-type
      --service-token-uuid string                               Equinix-assigned service token identifier
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/equinix/equinix.yaml)
      --debug           Enable debug mode to show HTTP requests and responses
```

### SEE ALSO

* [equinix fabricv4 service-tokens](equinix_fabricv4_service-tokens.md)	 - Manage service-tokens resources

