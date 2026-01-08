## equinix metalv1 metal-gateways find-metal-gateways-by-project

Returns all metal gateways for a project

### Synopsis

Return all metal gateways for a project

Use --request flag to provide optional JSON payload fields.

```
equinix metalv1 metal-gateways find-metal-gateways-by-project [flags]
```

### Options

```
      --exclude string      exclude field (JSON or string)
  -h, --help                help for find-metal-gateways-by-project
      --include string      include field (JSON or string)
      --page int            page field
      --per-page int        per-page field
      --project-id string   Project UUID (required)
      --request string      JSON payload for additional optional fields not exposed as flags
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug           Enable debug logging for HTTP requests
  -f, --format string   Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix metalv1 metal-gateways](equinix_metalv1_metal-gateways.md)	 - Manage metal-gateways resources

