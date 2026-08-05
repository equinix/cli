## equinix fabricv4 internet-access-services create-eia-service

Execute create-eia-service operation

### Synopsis

Execute the create-eia-service operation on this service.

Use --request flag to provide a JSON payload for the request body.
Example: --request '{"field":"value"}'

The command accepts parameters based on the SDK method signature.

```
equinix fabricv4 internet-access-services create-eia-service [flags]
```

### Options

```
  -h, --help                                                                         help for create-eia-service
      --internet-access-post-request-account-account-number string                   internet-access-post-request-account-account-number
      --internet-access-post-request-account-additional-properties string            internet-access-post-request-account-additional-properties (JSON)
      --internet-access-post-request-account-href string                             internet-access-post-request-account-href
      --internet-access-post-request-additional-properties string                    internet-access-post-request-additional-properties (JSON)
      --internet-access-post-request-bandwidth int                                   internet-access-post-request-bandwidth
      --internet-access-post-request-bandwidth-commit int                            internet-access-post-request-bandwidth-commit
      --internet-access-post-request-billing-additional-properties string            internet-access-post-request-billing-additional-properties (JSON)
      --internet-access-post-request-billing-type string                             internet-access-post-request-billing-type
      --internet-access-post-request-name string                                     internet-access-post-request-name
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

