## equinix metalv1 events find-events

Retrieve current user's events

### Synopsis

Returns a list of the current user’s events

Use --request flag to provide optional JSON payload fields.

```
equinix metalv1 events find-events [flags]
```

### Options

```
      --exclude string   exclude field (JSON or string)
  -h, --help             help for find-events
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

* [equinix metalv1 events](equinix_metalv1_events.md)	 - Manage events resources

