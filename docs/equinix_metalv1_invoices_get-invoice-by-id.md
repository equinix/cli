## equinix metalv1 invoices get-invoice-by-id

Retrieve an invoice

### Synopsis

Returns the invoice identified by the provided id

Use --request flag to provide optional JSON payload fields.

```
equinix metalv1 invoices get-invoice-by-id [flags]
```

### Options

```
  -h, --help             help for get-invoice-by-id
      --id string        Invoice UUID (required)
      --request string   JSON payload for request body fields
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug           Enable debug logging for HTTP requests
  -f, --format string   Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix metalv1 invoices](equinix_metalv1_invoices.md)	 - Manage invoices resources

