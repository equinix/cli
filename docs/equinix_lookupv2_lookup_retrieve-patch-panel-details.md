## equinix lookupv2 lookup retrieve-patch-panel-details

Retrieve patch panel details

### Synopsis

This method retrieves details of a patch panel by its ID to a user with cross connects ordering permission.

Use --request flag to provide optional JSON payload fields.

```
equinix lookupv2 lookup retrieve-patch-panel-details [flags]
```

### Options

```
      --a-side-ibx string                a-side-ibx field
      --account-number string            account-number field
  -h, --help                             help for retrieve-patch-panel-details
      --patch-panel-id string            ID of patch panel (required)
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

