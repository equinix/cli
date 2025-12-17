## equinix metalv1 organizations find-organization-customdata

Retrieve the custom metadata of an organization

### Synopsis

Provides the custom metadata stored for this organization in json format

Use --request flag to provide optional JSON payload fields.

```
equinix metalv1 organizations find-organization-customdata [flags]
```

### Options

```
  -h, --help             help for find-organization-customdata
      --id string        Organization UUID (required)
      --request string   JSON payload for request body fields
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug           Enable debug logging for HTTP requests
  -f, --format string   Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix metalv1 organizations](equinix_metalv1_organizations.md)	 - Manage organizations resources

