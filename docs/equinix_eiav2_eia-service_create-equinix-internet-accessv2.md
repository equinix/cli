## equinix eiav2 eia-service create-equinix-internet-accessv2

Creates Equinix Internet Access Service

### Synopsis

By passing in the appropriate options, you can create Equinix Internet Access Service product. The entire request either succeeds or fails. In case of failure all the changes in the system are rolled back, so the system gets back to its stated before submitting the request

Use --request flag to provide optional JSON payload fields.

```
equinix eiav2 eia-service create-equinix-internet-accessv2 [flags]
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

* [equinix eiav2 eia-service](equinix_eiav2_eia-service.md)	 - Manage eia-service resources

