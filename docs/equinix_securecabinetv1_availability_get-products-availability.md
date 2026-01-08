## equinix securecabinetv1 availability get-products-availability

Secure Cabinet availability.

### Synopsis

Returns the availability of Secure Cabinets and recommended additional products for provided billing account.

Use --request flag to provide optional JSON payload fields.

```
equinix securecabinetv1 availability get-products-availability [flags]
```

### Options

```
      --account-number string   Billing Account Number. (required)
  -h, --help                    help for get-products-availability
      --request string          JSON payload for request body fields
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug           Enable debug logging for HTTP requests
  -f, --format string   Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix securecabinetv1 availability](equinix_securecabinetv1_availability.md)	 - Manage availability resources

