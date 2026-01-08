## equinix metalv1 organizations find-organization-payment-methods

Retrieve all payment methods of an organization

### Synopsis

Returns all payment methods of an organization.

Use --request flag to provide optional JSON payload fields.

```
equinix metalv1 organizations find-organization-payment-methods [flags]
```

### Options

```
  -h, --help             help for find-organization-payment-methods
      --id string        Organization UUID (required)
      --include string   include field (JSON or string)
      --page int         page field
      --per-page int     per-page field
      --request string   JSON payload for additional optional fields not exposed as flags
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug           Enable debug logging for HTTP requests
  -f, --format string   Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix metalv1 organizations](equinix_metalv1_organizations.md)	 - Manage organizations resources

