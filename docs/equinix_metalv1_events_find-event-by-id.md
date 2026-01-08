## equinix metalv1 events find-event-by-id

Retrieve an event

### Synopsis

Returns a single event if the user has access

Use --request flag to provide optional JSON payload fields.

```
equinix metalv1 events find-event-by-id [flags]
```

### Options

```
      --exclude string   exclude field (JSON or string)
  -h, --help             help for find-event-by-id
      --id string        Event UUID (required)
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

* [equinix metalv1 events](equinix_metalv1_events.md)	 - Manage events resources

