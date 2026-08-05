## equinix fabricv4 ip-blocks submit-ip-block

Execute submit-ip-block operation

### Synopsis

Execute the submit-ip-block operation on this service.

Use --request flag to provide a JSON payload for the request body.
Example: --request '{"field":"value"}'

The command accepts parameters based on the SDK method signature.

```
equinix fabricv4 ip-blocks submit-ip-block [flags]
```

### Options

```
  -h, --help                                                                    help for submit-ip-block
      --request string                                                          JSON payload for additional optional fields not exposed as flags
      --submit-ip-block-request-body-account-account-number string              submit-ip-block-request-body-account-account-number
      --submit-ip-block-request-body-account-additional-properties string       submit-ip-block-request-body-account-additional-properties (JSON)
      --submit-ip-block-request-body-additional-properties string               submit-ip-block-request-body-additional-properties (JSON)
      --submit-ip-block-request-body-location-additional-properties string      submit-ip-block-request-body-location-additional-properties (JSON)
      --submit-ip-block-request-body-location-metro-code string                 submit-ip-block-request-body-location-metro-code
      --submit-ip-block-request-body-location-metro-href string                 submit-ip-block-request-body-location-metro-href
      --submit-ip-block-request-body-order-additional-properties string         submit-ip-block-request-body-order-additional-properties (JSON)
      --submit-ip-block-request-body-order-order-line string                    submit-ip-block-request-body-order-order-line (JSON)
      --submit-ip-block-request-body-order-order-number string                  submit-ip-block-request-body-order-order-number (JSON)
      --submit-ip-block-request-body-order-purchase-order-number string         submit-ip-block-request-body-order-purchase-order-number (JSON)
      --submit-ip-block-request-body-prefix string                              submit-ip-block-request-body-prefix
      --submit-ip-block-request-body-prefix-length int                          submit-ip-block-request-body-prefix-length
      --submit-ip-block-request-body-project-additional-properties string       submit-ip-block-request-body-project-additional-properties (JSON)
      --submit-ip-block-request-body-project-project-id string                  submit-ip-block-request-body-project-project-id
      --submit-ip-block-request-body-regulations-additional-properties string   submit-ip-block-request-body-regulations-additional-properties (JSON)
      --submit-ip-block-request-body-regulations-addressing-plans string        submit-ip-block-request-body-regulations-addressing-plans (JSON array)
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

