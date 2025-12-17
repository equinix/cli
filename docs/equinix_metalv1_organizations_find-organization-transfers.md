## equinix metalv1 organizations find-organization-transfers

Retrieve all project transfer requests from or to an organization

### Synopsis

Provides a collection of project transfer requests from or to the organization.

Use --request flag to provide optional JSON payload fields.

```
equinix metalv1 organizations find-organization-transfers [flags]
```

### Options

```
  -h, --help             help for find-organization-transfers
      --id string        Organization UUID (required)
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

* [equinix metalv1 organizations](equinix_metalv1_organizations.md)	 - Manage organizations resources

