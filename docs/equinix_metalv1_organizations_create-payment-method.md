## equinix metalv1 organizations create-payment-method

Create a payment method for the given organization

### Synopsis

Creates a payment method.

Use --request flag to provide optional JSON payload fields.

```
equinix metalv1 organizations create-payment-method [flags]
```

### Options

```
  -h, --help                                                       help for create-payment-method
      --id string                                                  Organization UUID (required)
      --include string                                             include field (JSON or string)
      --payment-method-create-input-additional-properties string   payment-method-create-input-additional-properties (JSON)
      --payment-method-create-input-default                        payment-method-create-input-default
      --payment-method-create-input-name string                    payment-method-create-input-name
      --payment-method-create-input-nonce string                   payment-method-create-input-nonce
      --request string                                             JSON payload for additional optional fields not exposed as flags
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug           Enable debug logging for HTTP requests
  -f, --format string   Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix metalv1 organizations](equinix_metalv1_organizations.md)	 - Manage organizations resources

