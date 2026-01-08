## equinix metalv1 payment-methods update-payment-method

Update the payment method

### Synopsis

Updates the payment method.

Use --request flag to provide optional JSON payload fields.

```
equinix metalv1 payment-methods update-payment-method [flags]
```

### Options

```
  -h, --help                                                       help for update-payment-method
      --id string                                                  Payment Method UUID (required)
      --include string                                             include field (JSON or string)
      --payment-method-update-input-additional-properties string   payment-method-update-input-additional-properties (JSON)
      --payment-method-update-input-billing_address string         payment-method-update-input-billing_address (JSON)
      --payment-method-update-input-cardholder_name string         payment-method-update-input-cardholder_name
      --payment-method-update-input-default                        payment-method-update-input-default
      --payment-method-update-input-expiration_month string        payment-method-update-input-expiration_month
      --payment-method-update-input-expiration_year int            payment-method-update-input-expiration_year
      --payment-method-update-input-name string                    payment-method-update-input-name
      --request string                                             JSON payload for additional optional fields not exposed as flags
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug           Enable debug logging for HTTP requests
  -f, --format string   Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix metalv1 payment-methods](equinix_metalv1_payment-methods.md)	 - Manage payment-methods resources

