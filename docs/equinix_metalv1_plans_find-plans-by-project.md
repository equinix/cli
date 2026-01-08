## equinix metalv1 plans find-plans-by-project

Retrieve all plans visible by the project

### Synopsis

Returns a listing of available plans for the given project

Use --request flag to provide optional JSON payload fields.

```
equinix metalv1 plans find-plans-by-project [flags]
```

### Options

```
      --exclude string   exclude field (JSON or string)
  -h, --help             help for find-plans-by-project
      --id string        Project UUID (required)
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

* [equinix metalv1 plans](equinix_metalv1_plans.md)	 - Manage plans resources

