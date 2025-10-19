## equinix fabricv4 cloud-routers create-cloud-router

Create Routers

### Synopsis

This API provides capability to create user's Cloud Routers

Use --request flag to provide optional JSON payload fields.

```
equinix fabricv4 cloud-routers create-cloud-router [flags]
```

### Options

```
      --cloud-router-post-request-account-account-name string                             Account name
      --cloud-router-post-request-account-account-number int                              Account number
      --cloud-router-post-request-account-additional-properties string                    cloud-router-post-request-account-additional-properties (JSON)
      --cloud-router-post-request-account-global-cust-id string                           Account name
      --cloud-router-post-request-account-global-org-id string                            Global organization identifier
      --cloud-router-post-request-account-global-organization-name string                 Global organization name
      --cloud-router-post-request-account-org-id int                                      Customer organization identifier
      --cloud-router-post-request-account-organization-name string                        Customer organization name
      --cloud-router-post-request-account-reseller-account-name string                    Reseller account name
      --cloud-router-post-request-account-reseller-account-number int                     Reseller account number
      --cloud-router-post-request-account-reseller-org-id int                             Reseller customer organization identifier
      --cloud-router-post-request-account-reseller-ucm-id string                          Reseller account ucmId
      --cloud-router-post-request-account-ucm-id string                                   Account ucmId
      --cloud-router-post-request-additional-properties string                            cloud-router-post-request-additional-properties (required) (JSON)
      --cloud-router-post-request-location-additional-properties string                   cloud-router-post-request-location-additional-properties (required) (JSON)
      --cloud-router-post-request-location-metro-code string                              cloud-router-post-request-location-metro-code (required)
      --cloud-router-post-request-location-metro-href string                              The Canonical URL at which the resource resides.
      --cloud-router-post-request-location-metro-name string                              cloud-router-post-request-location-metro-name
      --cloud-router-post-request-location-region string                                  cloud-router-post-request-location-region
      --cloud-router-post-request-marketplace-subscription-additional-properties string   cloud-router-post-request-marketplace-subscription-additional-properties (JSON)
      --cloud-router-post-request-marketplace-subscription-href string                    Marketplace Subscription URI
      --cloud-router-post-request-marketplace-subscription-type string                    cloud-router-post-request-marketplace-subscription-type
      --cloud-router-post-request-marketplace-subscription-uuid string                    Equinix-assigned Marketplace Subscription identifier
      --cloud-router-post-request-name string                                             Customer-provided Cloud Router name (required)
      --cloud-router-post-request-notifications string                                    Preferences for notifications on connection configuration or status changes (JSON array)
      --cloud-router-post-request-order-additional-properties string                      cloud-router-post-request-order-additional-properties (JSON)
      --cloud-router-post-request-order-billing-tier string                               Billing tier for connection bandwidth
      --cloud-router-post-request-order-customer-reference-number string                  Customer reference number
      --cloud-router-post-request-order-order-id string                                   Order Identification
      --cloud-router-post-request-order-order-number string                               Order Reference Number
      --cloud-router-post-request-order-purchase-order-number string                      Purchase order number
      --cloud-router-post-request-order-term-length int                                   Term length in months, valid values are 1, 12, 24, 36 where 1 is the default value (for on-demand case).
      --cloud-router-post-request-package-additional-properties string                    cloud-router-post-request-package-additional-properties (required) (JSON)
      --cloud-router-post-request-package-code string                                     cloud-router-post-request-package-code (required)
      --cloud-router-post-request-package-href string                                     Fabric Cloud Router URI
      --cloud-router-post-request-package-type string                                     cloud-router-post-request-package-type
      --cloud-router-post-request-project-additional-properties string                    cloud-router-post-request-project-additional-properties (JSON)
      --cloud-router-post-request-project-project-id string                               Subscriber-assigned project ID
      --cloud-router-post-request-type string                                             cloud-router-post-request-type (required)
      --dry-run                                                                           dry-run field (required)
  -h, --help                                                                              help for create-cloud-router
      --request string                                                                    JSON payload for additional optional fields not exposed as flags
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/equinix/equinix.yaml)
      --debug           Enable debug mode to show HTTP requests and responses
```

### SEE ALSO

* [equinix fabricv4 cloud-routers](equinix_fabricv4_cloud-routers.md)	 - Manage cloud-routers resources

