## equinix eiav2 e-i-a-service create-equinix-internet-accessv2

Execute create-equinix-internet-accessv2 operation

### Synopsis

Execute the create-equinix-internet-accessv2 operation on this service.

Use --request flag to provide a JSON payload for the request body.
Example: --request '{"field":"value"}'

The command accepts parameters based on the SDK method signature.

```
equinix eiav2 e-i-a-service create-equinix-internet-accessv2 [flags]
```

### Options

```
  -h, --help                                                                      help for create-equinix-internet-accessv2
      --request string                                                            JSON payload for additional optional fields not exposed as flags
      --service-request-additional-properties string                              service-request-additional-properties (JSON)
      --service-request-connections string                                        Collection of service connections uuids (JSON array)
      --service-request-description string                                        service-request-description
      --service-request-name string                                               service-request-name
      --service-request-order-additional-properties string                        service-request-order-additional-properties (JSON)
      --service-request-order-contacts string                                     service-request-order-contacts (JSON array)
      --service-request-order-draft                                               service-request-order-draft
      --service-request-order-purchase-order string                               service-request-order-purchase-order (JSON)
      --service-request-order-reference-number string                             service-request-order-reference-number
      --service-request-order-signature string                                    service-request-order-signature (JSON)
      --service-request-order-tags string                                         service-request-order-tags (JSON array)
      --service-request-routing-protocol-bgp-routing-protocol-request string      service-request-routing-protocol-bgp-routing-protocol-request (JSON)
      --service-request-routing-protocol-direct-routing-protocol-request string   service-request-routing-protocol-direct-routing-protocol-request (JSON)
      --service-request-routing-protocol-static-routing-protocol-request string   service-request-routing-protocol-static-routing-protocol-request (JSON)
      --service-request-tags string                                               service-request-tags (JSON array)
      --service-request-type string                                               service-request-type
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug           Enable debug logging for HTTP requests
  -f, --format string   Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix eiav2 e-i-a-service](equinix_eiav2_e-i-a-service.md)	 - Manage e-i-a-service resources

