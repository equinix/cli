## equinix fabricv4 internet-exchange-services create-exchange-service

Execute create-exchange-service operation

### Synopsis

Execute the create-exchange-service operation on this service.

Use --request flag to provide a JSON payload for the request body.
Example: --request '{"field":"value"}'

The command accepts parameters based on the SDK method signature.

```
equinix fabricv4 internet-exchange-services create-exchange-service [flags]
```

### Options

```
      --exchange-service-create-request-account-account-name string                              Account name
      --exchange-service-create-request-account-account-number int                               Account number
      --exchange-service-create-request-account-additional-properties string                     exchange-service-create-request-account-additional-properties (JSON)
      --exchange-service-create-request-account-global-cust-id string                            Account name
      --exchange-service-create-request-account-global-org-id string                             Global organization identifier
      --exchange-service-create-request-account-global-organization-name string                  Global organization name
      --exchange-service-create-request-account-org-id int                                       Customer organization identifier
      --exchange-service-create-request-account-organization-name string                         Customer organization name
      --exchange-service-create-request-account-reseller-account-name string                     Reseller account name
      --exchange-service-create-request-account-reseller-account-number int                      Reseller account number
      --exchange-service-create-request-account-reseller-org-id int                              Reseller customer organization identifier
      --exchange-service-create-request-account-reseller-ucm-id string                           Reseller account ucmId
      --exchange-service-create-request-account-ucm-id string                                    Account ucmId
      --exchange-service-create-request-additional-properties string                             exchange-service-create-request-additional-properties (JSON)
      --exchange-service-create-request-description string                                       exchange-service-create-request-description
      --exchange-service-create-request-name string                                              exchange-service-create-request-name
      --exchange-service-create-request-notifications string                                     exchange-service-create-request-notifications (JSON array)
      --exchange-service-create-request-order-additional-properties string                       exchange-service-create-request-order-additional-properties (JSON)
      --exchange-service-create-request-order-customer-reference-id string                       exchange-service-create-request-order-customer-reference-id
      --exchange-service-create-request-order-domain-order-number string                         exchange-service-create-request-order-domain-order-number
      --exchange-service-create-request-order-order-id string                                    exchange-service-create-request-order-order-id
      --exchange-service-create-request-order-order-line-item-number string                      exchange-service-create-request-order-order-line-item-number
      --exchange-service-create-request-order-order-number string                                exchange-service-create-request-order-order-number
      --exchange-service-create-request-order-purchase-order-number string                       exchange-service-create-request-order-purchase-order-number
      --exchange-service-create-request-order-uuid string                                        exchange-service-create-request-order-uuid
      --exchange-service-create-request-project-additional-properties string                     exchange-service-create-request-project-additional-properties (JSON)
      --exchange-service-create-request-project-project-id string                                Subscriber-assigned project ID
      --exchange-service-create-request-public-peering-connection-additional-properties string   exchange-service-create-request-public-peering-connection-additional-properties (JSON)
      --exchange-service-create-request-public-peering-connection-uuid string                    exchange-service-create-request-public-peering-connection-uuid
      --exchange-service-create-request-routing-protocol-additional-properties string            exchange-service-create-request-routing-protocol-additional-properties (JSON)
      --exchange-service-create-request-routing-protocol-bgp-ipv4 string                         exchange-service-create-request-routing-protocol-bgp-ipv4 (JSON)
      --exchange-service-create-request-routing-protocol-bgp-ipv6 string                         exchange-service-create-request-routing-protocol-bgp-ipv6 (JSON)
      --exchange-service-create-request-routing-protocol-customer-asn int                        exchange-service-create-request-routing-protocol-customer-asn
      --exchange-service-create-request-routing-protocol-mac-address string                      exchange-service-create-request-routing-protocol-mac-address
      --exchange-service-create-request-type string                                              exchange-service-create-request-type
  -h, --help                                                                                     help for create-exchange-service
      --request string                                                                           JSON payload for additional optional fields not exposed as flags
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug           Enable debug logging for HTTP requests
  -f, --format string   Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix fabricv4 internet-exchange-services](equinix_fabricv4_internet-exchange-services.md)	 - Manage internet-exchange-services resources

