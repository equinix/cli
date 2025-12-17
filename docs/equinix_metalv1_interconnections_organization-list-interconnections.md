## equinix metalv1 interconnections organization-list-interconnections

List organization connections

### Synopsis

List the connections belonging to the organization

Use --request flag to provide optional JSON payload fields.

```
equinix metalv1 interconnections organization-list-interconnections [flags]
```

### Options

```
      --exclude string           exclude field (JSON or string)
  -h, --help                     help for organization-list-interconnections
      --include string           include field (JSON or string)
      --organization-id string   UUID of the organization (required)
      --request string           JSON payload for additional optional fields not exposed as flags
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug           Enable debug logging for HTTP requests
  -f, --format string   Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix metalv1 interconnections](equinix_metalv1_interconnections.md)	 - Manage interconnections resources

