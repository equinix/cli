## equinix lookupv2 lookup retrieve-list-of-providers

Retrieve list of providers

### Synopsis

This method retrieves all available cross connect service providers to a user with cross connects ordering permission.

Use --request flag to provide optional JSON payload fields.

```
equinix lookupv2 lookup retrieve-list-of-providers [flags]
```

### Options

```
      --account-number string   account-number field
      --cage-id string          cage-id field
  -h, --help                    help for retrieve-list-of-providers
      --request string          JSON payload for additional optional fields not exposed as flags
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug           Enable debug logging for HTTP requests
  -f, --format string   Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix lookupv2 lookup](equinix_lookupv2_lookup.md)	 - Manage lookup resources

