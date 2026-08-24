## equinix fabricv4 ip-blocks submit-ip-block

Submits new Equinix owned or customer owned IP Block request

### Synopsis

Submits new Equinix owned or customer owned IP Block request

Use --request flag to provide optional JSON payload fields.

```
equinix fabricv4 ip-blocks submit-ip-block [flags]
```

### Options

```
  -h, --help                                                                    help for submit-ip-block
      --request string                                                          JSON payload for additional optional fields not exposed as flags
      --submit-ip-block-request-body-account-account-number string              account number
      --submit-ip-block-request-body-account-additional-properties string       submit-ip-block-request-body-account-additional-properties (JSON)
      --submit-ip-block-request-body-additional-properties string               submit-ip-block-request-body-additional-properties (JSON)
      --submit-ip-block-request-body-location-additional-properties string      submit-ip-block-request-body-location-additional-properties (JSON)
      --submit-ip-block-request-body-location-metro-code string                 submit-ip-block-request-body-location-metro-code
      --submit-ip-block-request-body-location-metro-href string                 Metro URL path for the linked resource
      --submit-ip-block-request-body-order-additional-properties string         submit-ip-block-request-body-order-additional-properties (JSON)
      --submit-ip-block-request-body-order-order-line string                    submit-ip-block-request-body-order-order-line (JSON)
      --submit-ip-block-request-body-order-order-number string                  submit-ip-block-request-body-order-order-number (JSON)
      --submit-ip-block-request-body-order-purchase-order-number string         submit-ip-block-request-body-order-purchase-order-number (JSON)
      --submit-ip-block-request-body-prefix string                              CIDR prefix
      --submit-ip-block-request-body-prefix-length int                          IpBlockPrefix length
      --submit-ip-block-request-body-project-additional-properties string       submit-ip-block-request-body-project-additional-properties (JSON)
      --submit-ip-block-request-body-project-project-id string                  project id
      --submit-ip-block-request-body-regulations-additional-properties string   submit-ip-block-request-body-regulations-additional-properties (JSON)
      --submit-ip-block-request-body-regulations-addressing-plans string        List of addressing plans (JSON array)
      --submit-ip-block-request-body-regulations-questions string               submit-ip-block-request-body-regulations-questions (JSON)
      --submit-ip-block-request-body-type string                                submit-ip-block-request-body-type
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug           Enable debug logging for HTTP requests
  -f, --format string   Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix fabricv4 ip-blocks](equinix_fabricv4_ip-blocks.md)	 - Manage ip-blocks resources

