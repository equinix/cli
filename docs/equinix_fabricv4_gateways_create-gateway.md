## equinix fabricv4 gateways create-gateway

Create Gateway

### Synopsis

This API provides capability to create user's Gateways

Use --request flag to provide optional JSON payload fields.

```
equinix fabricv4 gateways create-gateway [flags]
```

### Options

```
      --gateway-post-request-account-account-name string               Account name
      --gateway-post-request-account-account-number int                Account number
      --gateway-post-request-account-additional-properties string      gateway-post-request-account-additional-properties (JSON)
      --gateway-post-request-account-global-cust-id string             Account name
      --gateway-post-request-account-global-org-id string              Global organization identifier
      --gateway-post-request-account-global-organization-name string   Global organization name
      --gateway-post-request-account-org-id int                        Customer organization identifier
      --gateway-post-request-account-organization-name string          Customer organization name
      --gateway-post-request-account-reseller-account-name string      Reseller account name
      --gateway-post-request-account-reseller-account-number int       Reseller account number
      --gateway-post-request-account-reseller-org-id int               Reseller customer organization identifier
      --gateway-post-request-account-reseller-ucm-id string            Reseller account ucmId
      --gateway-post-request-account-ucm-id string                     Account ucmId
      --gateway-post-request-additional-properties string              gateway-post-request-additional-properties (JSON)
      --gateway-post-request-bandwidth int                             Gateway bandwidth in Mbps
      --gateway-post-request-description string                        Customer-provided Gateway description
      --gateway-post-request-ha-enabled                                High availability enabled
      --gateway-post-request-local-asn int                             Gateway local Autonomous System Number
      --gateway-post-request-name string                               Customer-provided Gateway name
      --gateway-post-request-order-additional-properties string        gateway-post-request-order-additional-properties (JSON)
      --gateway-post-request-order-billing-tier string                 Billing tier for connection bandwidth
      --gateway-post-request-order-contracted-bandwidth int            Contracted bandwidth
      --gateway-post-request-order-customer-reference-number string    Customer reference number
      --gateway-post-request-order-order-id string                     Order Identification
      --gateway-post-request-order-order-number string                 Order Reference Number
      --gateway-post-request-order-purchase-order-number string        Purchase order number
      --gateway-post-request-order-term-length int                     Term length in months, valid values are 1, 12, 24, 36 where 1 is the default value (for on-demand case).
      --gateway-post-request-project-additional-properties string      gateway-post-request-project-additional-properties (JSON)
      --gateway-post-request-project-project-id string                 Subscriber-assigned project ID
      --gateway-post-request-router-additional-properties string       gateway-post-request-router-additional-properties (JSON)
      --gateway-post-request-router-href string                        Fabric Cloud Router URI
      --gateway-post-request-router-type string                        gateway-post-request-router-type
      --gateway-post-request-router-uuid string                        Cloud Router UUID
      --gateway-post-request-type string                               gateway-post-request-type
  -h, --help                                                           help for create-gateway
      --request string                                                 JSON payload for additional optional fields not exposed as flags
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug           Enable debug logging for HTTP requests
  -f, --format string   Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix fabricv4 gateways](equinix_fabricv4_gateways.md)	 - Manage gateways resources

