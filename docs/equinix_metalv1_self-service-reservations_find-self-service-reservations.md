## equinix metalv1 self-service-reservations find-self-service-reservations

Retrieve all reservations

### Synopsis

Returns all reservations.

Use --request flag to provide optional JSON payload fields.

```
equinix metalv1 self-service-reservations find-self-service-reservations [flags]
```

### Options

```
      --categories string   categories field (JSON or string)
  -h, --help                help for find-self-service-reservations
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

* [equinix metalv1 self-service-reservations](equinix_metalv1_self-service-reservations.md)	 - Manage self-service-reservations resources

