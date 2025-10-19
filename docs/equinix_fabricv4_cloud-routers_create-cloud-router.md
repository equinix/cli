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
      --cloud-router-post-request-account-account-name string                             cloud-router-post-request-account-account-name
      --cloud-router-post-request-account-account-number int                              cloud-router-post-request-account-account-number
      --cloud-router-post-request-account-additional-properties string                    cloud-router-post-request-account-additional-properties (JSON)
      --cloud-router-post-request-account-global-cust-id string                           cloud-router-post-request-account-global-cust-id
      --cloud-router-post-request-account-global-org-id string                            cloud-router-post-request-account-global-org-id
      --cloud-router-post-request-account-global-organization-name string                 cloud-router-post-request-account-global-organization-name
      --cloud-router-post-request-account-org-id int                                      cloud-router-post-request-account-org-id
      --cloud-router-post-request-account-organization-name string                        cloud-router-post-request-account-organization-name
      --cloud-router-post-request-account-reseller-account-name string                    cloud-router-post-request-account-reseller-account-name
      --cloud-router-post-request-account-reseller-account-number int                     cloud-router-post-request-account-reseller-account-number
      --cloud-router-post-request-account-reseller-org-id int                             cloud-router-post-request-account-reseller-org-id
      --cloud-router-post-request-account-reseller-ucm-id string                          cloud-router-post-request-account-reseller-ucm-id
      --cloud-router-post-request-account-ucm-id string                                   cloud-router-post-request-account-ucm-id
      --cloud-router-post-request-additional-properties string                            cloud-router-post-request-additional-properties (required) (JSON)
      --cloud-router-post-request-location-additional-properties string                   cloud-router-post-request-location-additional-properties (required) (JSON)
      --cloud-router-post-request-location-metro-code string                              cloud-router-post-request-location-metro-code (required)
      --cloud-router-post-request-location-metro-href string                              cloud-router-post-request-location-metro-href
      --cloud-router-post-request-location-metro-name string                              cloud-router-post-request-location-metro-name
      --cloud-router-post-request-location-region string                                  cloud-router-post-request-location-region
      --cloud-router-post-request-marketplace-subscription-additional-properties string   cloud-router-post-request-marketplace-subscription-additional-properties (JSON)
      --cloud-router-post-request-marketplace-subscription-href string                    cloud-router-post-request-marketplace-subscription-href
      --cloud-router-post-request-marketplace-subscription-type string                    cloud-router-post-request-marketplace-subscription-type
      --cloud-router-post-request-marketplace-subscription-uuid string                    cloud-router-post-request-marketplace-subscription-uuid
      --cloud-router-post-request-name string                                             cloud-router-post-request-name (required)
      --cloud-router-post-request-notifications string                                    cloud-router-post-request-notifications (JSON array)
      --cloud-router-post-request-order-additional-properties string                      cloud-router-post-request-order-additional-properties (JSON)
      --cloud-router-post-request-order-billing-tier string                               cloud-router-post-request-order-billing-tier
      --cloud-router-post-request-order-customer-reference-number string                  cloud-router-post-request-order-customer-reference-number
      --cloud-router-post-request-order-order-id string                                   cloud-router-post-request-order-order-id
      --cloud-router-post-request-order-order-number string                               cloud-router-post-request-order-order-number
      --cloud-router-post-request-order-purchase-order-number string                      cloud-router-post-request-order-purchase-order-number
      --cloud-router-post-request-order-term-length int                                   cloud-router-post-request-order-term-length
      --cloud-router-post-request-package-additional-properties string                    cloud-router-post-request-package-additional-properties (required) (JSON)
      --cloud-router-post-request-package-code string                                     cloud-router-post-request-package-code (required)
      --cloud-router-post-request-package-href string                                     cloud-router-post-request-package-href
      --cloud-router-post-request-package-type string                                     cloud-router-post-request-package-type
      --cloud-router-post-request-project-additional-properties string                    cloud-router-post-request-project-additional-properties (JSON)
      --cloud-router-post-request-project-project-id string                               cloud-router-post-request-project-project-id
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

