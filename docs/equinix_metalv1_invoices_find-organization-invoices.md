## equinix metalv1 invoices find-organization-invoices

Retrieve all invoices for an organization

### Synopsis

Returns all invoices for an organization

Use --request flag to provide optional JSON payload fields.

```
equinix metalv1 invoices find-organization-invoices [flags]
```

### Options

```
  -h, --help             help for find-organization-invoices
      --id string        Organization UUID (required)
      --page int         page field
      --per-page int     per-page field
      --request string   JSON payload for additional optional fields not exposed as flags
      --status string    status field
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug           Enable debug logging for HTTP requests
  -f, --format string   Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix metalv1 invoices](equinix_metalv1_invoices.md)	 - Manage invoices resources

