## equinix fabricv4 fabric-one create-interconnect

Create Interconnect

### Synopsis

This API provides capability to create user's Interconnect <font color="red"> <sup color='red'>Beta</sup></font>

Use --request flag to provide optional JSON payload fields.

```
equinix fabricv4 fabric-one create-interconnect [flags]
```

### Options

```
  -h, --help                                                                help for create-interconnect
      --interconnect-post-request-account-account-name string               Account name
      --interconnect-post-request-account-account-number int                Account number
      --interconnect-post-request-account-additional-properties string      interconnect-post-request-account-additional-properties (JSON)
      --interconnect-post-request-account-global-cust-id string             Account name
      --interconnect-post-request-account-global-org-id string              Global organization identifier
      --interconnect-post-request-account-global-organization-name string   Global organization name
      --interconnect-post-request-account-org-id int                        Customer organization identifier
      --interconnect-post-request-account-organization-name string          Customer organization name
      --interconnect-post-request-account-reseller-account-name string      Reseller account name
      --interconnect-post-request-account-reseller-account-number int       Reseller account number
      --interconnect-post-request-account-reseller-org-id int               Reseller customer organization identifier
      --interconnect-post-request-account-reseller-ucm-id string            Reseller account ucmId
      --interconnect-post-request-account-ucm-id string                     Account ucmId
      --interconnect-post-request-additional-properties string              interconnect-post-request-additional-properties (JSON)
      --interconnect-post-request-description string                        Customer-provided interconnect description
      --interconnect-post-request-location-additional-properties string     interconnect-post-request-location-additional-properties (JSON)
      --interconnect-post-request-location-metro-code string                Metro code where the interconnect is created
      --interconnect-post-request-name string                               Customer-provided interconnect name
      --interconnect-post-request-notifications string                      Preferences for notifications on interconnect configuration or status changes (JSON array)
      --interconnect-post-request-order-additional-properties string        interconnect-post-request-order-additional-properties (JSON)
      --interconnect-post-request-order-billing-tier string                 Billing tier for connection bandwidth
      --interconnect-post-request-order-contracted-bandwidth int            Contracted bandwidth
      --interconnect-post-request-order-customer-reference-number string    Customer reference number
      --interconnect-post-request-order-order-id string                     Order Identification
      --interconnect-post-request-order-order-number string                 Order Reference Number
      --interconnect-post-request-order-purchase-order-number string        Purchase order number
      --interconnect-post-request-order-term-length int                     Term length in months, valid values are 1, 12, 24, 36 where 1 is the default value (for on-demand case).
      --interconnect-post-request-package-additional-properties string      interconnect-post-request-package-additional-properties (JSON)
      --interconnect-post-request-package-bandwidth-max int                 Maximum bandwidth in Mbps
      --interconnect-post-request-package-code string                       Interconnect Package code (e.g. LAB, BASIC, STANDARD, PREMIUM)
      --interconnect-post-request-package-description string                Interconnect Package description
      --interconnect-post-request-package-href string                       Interconnect Package URI
      --interconnect-post-request-package-is-remote                         Authorization to connect remotely
      --interconnect-post-request-package-routes-max int                    Maximum number of routes
      --interconnect-post-request-package-type string                       Interconnect Package Type
      --interconnect-post-request-package-uuid string                       Equinix-assigned Interconnect Package identifier
      --interconnect-post-request-project-additional-properties string      interconnect-post-request-project-additional-properties (JSON)
      --interconnect-post-request-project-project-id string                 Subscriber-assigned project ID
      --interconnect-post-request-type string                               interconnect-post-request-type
      --request string                                                      JSON payload for additional optional fields not exposed as flags
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug           Enable debug logging for HTTP requests
  -f, --format string   Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix fabricv4 fabric-one](equinix_fabricv4_fabric-one.md)	 - Manage fabric-one resources

