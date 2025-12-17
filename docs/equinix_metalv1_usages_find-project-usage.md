## equinix metalv1 usages find-project-usage

Retrieve all usages for project

### Synopsis

Returns all usages for a project.

Use --request flag to provide optional JSON payload fields.

```
equinix metalv1 usages find-project-usage [flags]
```

### Options

```
      --created-after string    created-after field
      --created-before string   created-before field
  -h, --help                    help for find-project-usage
      --id string               Project UUID (required)
      --request string          JSON payload for additional optional fields not exposed as flags
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug           Enable debug logging for HTTP requests
  -f, --format string   Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix metalv1 usages](equinix_metalv1_usages.md)	 - Manage usages resources

