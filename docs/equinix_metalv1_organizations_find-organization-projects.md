## equinix metalv1 organizations find-organization-projects

Retrieve all projects of an organization

### Synopsis

Returns a collection of projects that belong to the organization.

Use --request flag to provide optional JSON payload fields.

```
equinix metalv1 organizations find-organization-projects [flags]
```

### Options

```
      --exclude string   exclude field (JSON or string)
  -h, --help             help for find-organization-projects
      --id string        Organization UUID (required)
      --include string   include field (JSON or string)
      --name string      name field
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

