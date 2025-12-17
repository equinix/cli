## equinix metalv1 hardware-reservations find-project-hardware-reservations

Retrieve all hardware reservations for a given project

### Synopsis

Provides a collection of hardware reservations for a given project.

Use --request flag to provide optional JSON payload fields.

```
equinix metalv1 hardware-reservations find-project-hardware-reservations [flags]
```

### Options

```
      --exclude string         exclude field (JSON or string)
  -h, --help                   help for find-project-hardware-reservations
      --id string              Project UUID (required)
      --include string         include field (JSON or string)
      --page int               page field
      --per-page int           per-page field
      --provisionable string   provisionable field
      --query string           query field
      --request string         JSON payload for additional optional fields not exposed as flags
      --state string           state field
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.config/equinix/equinix.yaml)
      --debug           Enable debug logging for HTTP requests
  -f, --format string   Format to use for output (json or yaml) (default "json")
```

### SEE ALSO

* [equinix metalv1 hardware-reservations](equinix_metalv1_hardware-reservations.md)	 - Manage hardware-reservations resources

