## equinix fabricv4 connections create-connection

Create Connection

### Synopsis

This API provides capability to create user's virtual connection

Use --request flag to provide optional JSON payload fields.

```
equinix fabricv4 connections create-connection [flags]
```

### Options

```
      --connection-post-request-a-side-access-point string                              connection-post-request-a-side-access-point (JSON)
      --connection-post-request-a-side-additional-info string                           connection-post-request-a-side-additional-info (JSON array)
      --connection-post-request-a-side-additional-properties string                     connection-post-request-a-side-additional-properties (required) (JSON)
      --connection-post-request-a-side-company-profile string                           connection-post-request-a-side-company-profile (JSON)
      --connection-post-request-a-side-internet-access string                           connection-post-request-a-side-internet-access (JSON)
      --connection-post-request-a-side-invitation string                                connection-post-request-a-side-invitation (JSON)
      --connection-post-request-a-side-service-token string                             connection-post-request-a-side-service-token (JSON)
      --connection-post-request-additional-info string                                  connection-post-request-additional-info (JSON array)
      --connection-post-request-additional-properties string                            connection-post-request-additional-properties (required) (JSON)
      --connection-post-request-bandwidth int                                           connection-post-request-bandwidth (required)
      --connection-post-request-end-customer-additional-properties string               connection-post-request-end-customer-additional-properties (JSON)
      --connection-post-request-end-customer-is-disclosed                               connection-post-request-end-customer-is-disclosed
      --connection-post-request-end-customer-mdm-id string                              connection-post-request-end-customer-mdm-id
      --connection-post-request-end-customer-name string                                connection-post-request-end-customer-name
      --connection-post-request-geo-scope string                                        connection-post-request-geo-scope
      --connection-post-request-marketplace-subscription-additional-properties string   connection-post-request-marketplace-subscription-additional-properties (JSON)
      --connection-post-request-marketplace-subscription-href string                    connection-post-request-marketplace-subscription-href
      --connection-post-request-marketplace-subscription-type string                    connection-post-request-marketplace-subscription-type
      --connection-post-request-marketplace-subscription-uuid string                    connection-post-request-marketplace-subscription-uuid
      --connection-post-request-name string                                             connection-post-request-name (required)
      --connection-post-request-notifications string                                    connection-post-request-notifications (required) (JSON array)
      --connection-post-request-order-additional-properties string                      connection-post-request-order-additional-properties (JSON)
      --connection-post-request-order-billing-tier string                               connection-post-request-order-billing-tier
      --connection-post-request-order-customer-reference-number string                  connection-post-request-order-customer-reference-number
      --connection-post-request-order-order-id string                                   connection-post-request-order-order-id
      --connection-post-request-order-order-number string                               connection-post-request-order-order-number
      --connection-post-request-order-purchase-order-number string                      connection-post-request-order-purchase-order-number
      --connection-post-request-order-term-length int                                   connection-post-request-order-term-length
      --connection-post-request-project-additional-properties string                    connection-post-request-project-additional-properties (JSON)
      --connection-post-request-project-project-id string                               connection-post-request-project-project-id
      --connection-post-request-redundancy-additional-properties string                 connection-post-request-redundancy-additional-properties (JSON)
      --connection-post-request-redundancy-group string                                 connection-post-request-redundancy-group
      --connection-post-request-redundancy-priority string                              connection-post-request-redundancy-priority
      --connection-post-request-type string                                             connection-post-request-type (required)
      --connection-post-request-z-side-access-point string                              connection-post-request-z-side-access-point (JSON)
      --connection-post-request-z-side-additional-info string                           connection-post-request-z-side-additional-info (JSON array)
      --connection-post-request-z-side-additional-properties string                     connection-post-request-z-side-additional-properties (required) (JSON)
      --connection-post-request-z-side-company-profile string                           connection-post-request-z-side-company-profile (JSON)
      --connection-post-request-z-side-internet-access string                           connection-post-request-z-side-internet-access (JSON)
      --connection-post-request-z-side-invitation string                                connection-post-request-z-side-invitation (JSON)
      --connection-post-request-z-side-service-token string                             connection-post-request-z-side-service-token (JSON)
      --dry-run                                                                         dry-run field (required)
  -h, --help                                                                            help for create-connection
      --request string                                                                  JSON payload for additional optional fields not exposed as flags
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/equinix/equinix.yaml)
      --debug           Enable debug mode to show HTTP requests and responses
```

### SEE ALSO

* [equinix fabricv4 connections](equinix_fabricv4_connections.md)	 - Manage connections resources

