## equinix metalv1 payment-methods find-payment-method-by-id

Retrieve a payment method

### Synopsis

Returns a payment method

Use --request flag to provide optional JSON payload fields.

```
equinix metalv1 payment-methods find-payment-method-by-id [flags]
```

### Options

```
  -h, --help             help for find-payment-method-by-id
      --id string        Payment Method UUID (required)
      --include string   include field (JSON or string)
      --request string   JSON payload for additional optional fields not exposed as flags
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug           Enable debug logging for HTTP requests
  -f, --format string   Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix metalv1 payment-methods](equinix_metalv1_payment-methods.md)	 - Manage payment-methods resources

