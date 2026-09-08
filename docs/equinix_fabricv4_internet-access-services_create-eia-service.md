## equinix fabricv4 internet-access-services create-eia-service

Creates Internet Access Service

### Synopsis

Creates Internet Access Service

Use --request flag to provide optional JSON payload fields.

```
equinix fabricv4 internet-access-services create-eia-service [flags]
```

### Options

```
  -h, --help                                                                         help for create-eia-service
      --internet-access-post-request-account-account-number string                   Account number
      --internet-access-post-request-account-additional-properties string            internet-access-post-request-account-additional-properties (JSON)
      --internet-access-post-request-account-href string                             Account URL path
      --internet-access-post-request-additional-properties string                    internet-access-post-request-additional-properties (JSON)
      --internet-access-post-request-bandwidth int                                   Bandwidth of the service
      --internet-access-post-request-bandwidth-commit int                            Minimum bandwidth commit for burst billing variant of the service
      --internet-access-post-request-billing-additional-properties string            internet-access-post-request-billing-additional-properties (JSON)
      --internet-access-post-request-billing-type string                             internet-access-post-request-billing-type
      --internet-access-post-request-name string                                     The name of the EIA Service
      --internet-access-post-request-order-additional-properties string              internet-access-post-request-order-additional-properties (JSON)
      --internet-access-post-request-order-order-line string                         internet-access-post-request-order-order-line (JSON)
      --internet-access-post-request-order-order-number string                       internet-access-post-request-order-order-number (JSON)
      --internet-access-post-request-order-purchase-order-number string              internet-access-post-request-order-purchase-order-number (JSON)
      --internet-access-post-request-project-additional-properties string            internet-access-post-request-project-additional-properties (JSON)
      --internet-access-post-request-project-project-id string                       Subscriber-assigned project ID
      --internet-access-post-request-routing-protocol-additional-properties string   internet-access-post-request-routing-protocol-additional-properties (JSON)
      --internet-access-post-request-routing-protocol-customer-routes string         internet-access-post-request-routing-protocol-customer-routes (JSON array)
      --internet-access-post-request-routing-protocol-type string                    internet-access-post-request-routing-protocol-type
      --internet-access-post-request-type string                                     internet-access-post-request-type
      --request string                                                               JSON payload for additional optional fields not exposed as flags
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug           Enable debug logging for HTTP requests
  -f, --format string   Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix fabricv4 internet-access-services](equinix_fabricv4_internet-access-services.md)	 - Manage internet-access-services resources

