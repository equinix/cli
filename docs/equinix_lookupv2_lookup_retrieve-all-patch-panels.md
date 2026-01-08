## equinix lookupv2 lookup retrieve-all-patch-panels

Retrieve all patch panels

### Synopsis

This method retrieves all available patch panels associated with a cabinet for a user with cross connects ordering permission.

Use --request flag to provide optional JSON payload fields.

```
equinix lookupv2 lookup retrieve-all-patch-panels [flags]
```

### Options

```
      --a-side-ibx string                a-side-ibx field
      --account-number string            account-number field
      --cabinet-id string                cabinet-id field
  -h, --help                             help for retrieve-all-patch-panels
      --provider-account-number string   provider-account-number field
      --request string                   JSON payload for additional optional fields not exposed as flags
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug           Enable debug logging for HTTP requests
  -f, --format string   Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix lookupv2 lookup](equinix_lookupv2_lookup.md)	 - Manage lookup resources

