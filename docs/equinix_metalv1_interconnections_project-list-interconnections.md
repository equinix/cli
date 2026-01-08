## equinix metalv1 interconnections project-list-interconnections

List project connections

### Synopsis

List the connections belonging to the project

Use --request flag to provide optional JSON payload fields.

```
equinix metalv1 interconnections project-list-interconnections [flags]
```

### Options

```
      --exclude string      exclude field (JSON or string)
  -h, --help                help for project-list-interconnections
      --include string      include field (JSON or string)
      --page int            page field
      --per-page int        per-page field
      --project-id string   UUID of the project (required)
      --request string      JSON payload for additional optional fields not exposed as flags
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug           Enable debug logging for HTTP requests
  -f, --format string   Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix metalv1 interconnections](equinix_metalv1_interconnections.md)	 - Manage interconnections resources

