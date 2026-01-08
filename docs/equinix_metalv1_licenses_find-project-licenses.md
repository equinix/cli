## equinix metalv1 licenses find-project-licenses

Retrieve all licenses

### Synopsis

Provides a collection of licenses for a given project.

Use --request flag to provide optional JSON payload fields.

```
equinix metalv1 licenses find-project-licenses [flags]
```

### Options

```
      --exclude string   exclude field (JSON or string)
  -h, --help             help for find-project-licenses
      --id string        Project UUID (required)
      --include string   include field (JSON or string)
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

* [equinix metalv1 licenses](equinix_metalv1_licenses.md)	 - Manage licenses resources

